package logic

import (
	"ai-gozero-agent/mcp/internal/svc"
	"ai-gozero-agent/mcp/internal/utils"
	"ai-gozero-agent/mcp/types/mcp"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExtractTextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExtractTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExtractTextLogic {
	return &ExtractTextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流式上传PDF并返回解析文本
func (l *ExtractTextLogic) ExtractText(stream mcp.PdfProcessor_ExtractTextServer) error {
	// todo: add your logic here and delete this line
	firstChunk, err := stream.Recv()
	if err != nil {
		logx.Errorf("stream.Recv() failed: %v", err)
		return err
	}

	metadata := firstChunk.GetMetadata()
	if metadata == nil {
		return stream.SendAndClose(&mcp.PdfResponse{Error: "缺少元数据"})
	}
	if metadata.MimeType != "application/pdf" {
		return stream.SendAndClose(&mcp.PdfResponse{Error: "只支持PDF文件"})
	}

	temFile, err := os.CreateTemp("", "pdf-*.pdf")
	if err != nil {
		logx.Errorf("os.CreateTemp() failed: %v", err)
		return err
	}

	defer os.Remove(temFile.Name())
	defer temFile.Close()
	if chunk := firstChunk.GetChunk(); chunk != nil {
		if _, err := temFile.Write(chunk); err != nil {
			logx.Errorf("temFile.Write() failed: %v", err)
			return err
		}
	}

	if chunk := firstChunk.GetChunk(); chunk != nil {
		if _, err := temFile.Write(chunk); err != nil {
			logx.Errorf("temFile.Write() failed: %v", err)
			return err
		}
	}

	for {
		req, err := stream.Recv()
		if err != nil {
			logx.Errorf("stream.Recv() failed: %v", err)
			return err
		}
		if err == io.EOF {
			break
		}
		if chunk := req.GetChunk(); chunk != nil {
			if _, err := temFile.Write(chunk); err != nil {
				logx.Errorf("temFile.Write() failed: %v", err)
				return err
			}
		}
	}
	content, err := extractPdfText(temFile.Name())
	if err != nil {
		logx.Errorf("extractPdfText() failed: %v", err)
		return stream.SendAndClose(&mcp.PdfResponse{Error: err.Error()})
	}
	fmt.Println("消息解析完成， 打包发给api", content)
	return stream.SendAndClose(&mcp.PdfResponse{Content: content})
}

func extractPdfText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return utils.ExtractPDFText(file)
}
