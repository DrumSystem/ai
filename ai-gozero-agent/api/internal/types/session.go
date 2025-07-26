package types

import openai "github.com/sashabaranov/go-openai"

// 新增向量存储消息结构
type VectorMessage struct {
	Role    string `json:"role"`    // 消息角色
	Content string `json:"content"` // 消息内容
}

type ChatSession struct {
	Messages []openai.ChatCompletionMessage
}

// 会话存储接口更新
type SessionStore interface {
	GetSession(chatId string) ([]openai.ChatCompletionMessage, error) // 获取消息历史
	SaveMessage(chatId, role, content string) error                   // 保存单条消息
}
