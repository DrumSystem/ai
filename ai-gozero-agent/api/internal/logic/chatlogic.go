package logic

import (
	"ai-gozero-agent/api/internal/utils"
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"strings"

	"ai-gozero-agent/api/internal/svc"
	"ai-gozero-agent/api/internal/types"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.InterviewAPPChatReq) (<-chan *types.ChatResponse, error) {
	ch := make(chan *types.ChatResponse)

	go func() {
		defer close(ch)

		// 1. 保存用户消息到向量数据库
		if err := l.svcCtx.VectorStore.SaveMessage(req.ChatId, openai.ChatMessageRoleUser, req.Message); err != nil {
			l.Logger.Errorf("保存用户消息失败: %v", err)
			// 不返回，继续处理对话
		}

		//2. 知识检索
		knowledge, err := l.svcCtx.VectorStore.RetrieveKnowledge(req.Message, 3)
		if err != nil {
			l.Logger.Errorf("知识检索失败: %v", err)
			knowledge = []types.KnowledgeChunk{}
		}

		// 2. 获取会话历史
		messages, err := l.getSessionHistory(req.ChatId, knowledge)
		if err != nil {
			l.Logger.Errorf("获取会话历史失败: %v", err)
			ch <- &types.ChatResponse{Content: "系统错误：无法获取对话历史", IsLast: true}
			return
		}

		// 3. 创建OpenAI请求
		request := openai.ChatCompletionRequest{
			Model:       l.svcCtx.Config.OpenAI.Model,
			Messages:    messages,
			Stream:      true,
			MaxTokens:   l.svcCtx.Config.OpenAI.MaxTokens,
			Temperature: l.svcCtx.Config.OpenAI.Temperature,
		}

		// 4. 创建流式响应
		stream, err := l.svcCtx.OpenAIClient.CreateChatCompletionStream(l.ctx, request)
		if err != nil {
			l.Logger.Error("创建聊天完成流失败: ", err)
			ch <- &types.ChatResponse{Content: "系统错误：无法连接AI服务", IsLast: true}
			return
		}
		defer stream.Close()

		// 5. 处理流式响应
		var fullResponse strings.Builder
		for {
			select {
			case <-l.ctx.Done(): // 上下文取消
				return
			default:
				response, err := stream.Recv()
				if errors.Is(err, io.EOF) { // 流结束
					// 保存助手回复
					if fullResponse.String() != "" {
						if saveErr := l.svcCtx.VectorStore.SaveMessage(
							req.ChatId,
							openai.ChatMessageRoleAssistant,
							fullResponse.String(),
						); saveErr != nil {
							l.Logger.Errorf("保存助手消息失败: %v", saveErr)
						}
					}
					ch <- &types.ChatResponse{IsLast: true}
					return
				}
				if err != nil {
					l.Logger.Error("接收流数据失败: ", err)
					return
				}

				// 处理有效响应
				if len(response.Choices) > 0 && response.Choices[0].Delta.Content != "" {
					content := response.Choices[0].Delta.Content
					fullResponse.WriteString(content)
					ch <- &types.ChatResponse{
						Content: content,
						IsLast:  false,
					}
				}
			}
		}
	}()

	return ch, nil
}

// 获取会话历史
func (l *ChatLogic) getSessionHistory(chatId string, knowledge []types.KnowledgeChunk) ([]openai.ChatCompletionMessage, error) {
	// 获取最近的10条消息（约5轮对话）
	vectorMessages, err := l.svcCtx.VectorStore.GetMessages(chatId, 10)
	if err != nil {
		return nil, err
	}

	// 构建系统消息 - 注入知识
	systemMessage := "你是一个专业的Go语言面试官，负责评估候选人的Go语言能力。请提出有深度的问题并评估回答。"
	if len(knowledge) > 0 {
		systemMessage += "\n\n相关背景知识："
		for id, chunk := range knowledge {
			truncatedContent := utils.TruncateText(chunk.Content, 500)
			systemMessage += fmt.Sprintf("\n[知识片段%d] %s: %s", id+1, chunk.Title, truncatedContent)
		}
	}

	fmt.Println("检索的数据", systemMessage)
	// 转换为OpenAI消息格式
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemMessage,
		},
	}

	// 添加历史消息
	for _, msg := range vectorMessages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return messages, nil
}
