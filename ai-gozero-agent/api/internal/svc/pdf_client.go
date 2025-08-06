package svc

import (
	"ai-gozero-agent/mcp/types/mcp"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"io"
	"mime/multipart"
)

type PdfClient struct {
	client mcp.PdfProcessorClient
}

func NewPdfClient(endpoint string) *PdfClient {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{endpoint},
		NonBlock:  true,
	})
	return &PdfClient{
		client: mcp.NewPdfProcessorClient(conn.Conn()),
	}

}

func (c *PdfClient) ExtractText(file multipart.File, filename string) (string, error) {
	stream, err := c.client.ExtractText(context.Background())
	if err != nil {
		logx.Error("grpc连接失败：%v", err)
		return "", err
	}
	defer func() {
		if err := stream.CloseSend(); err != nil {
			logx.Error("关闭grpc连接失败：%v", err)
		}
	}()

	if err := stream.Send(&mcp.PdfRequest{Data: &mcp.PdfRequest_Metadata{
		Metadata: &mcp.Metadata{
			Filename: filename,
			MimeType: "application/pdf",
		},
	}}); err != nil {
		logx.Error("发送元数据失败：%v", err)
		return "", err
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		logx.Error("读取文件失败：%v", err)
		return "", err
	}

	if err := stream.Send(&mcp.PdfRequest{Data: &mcp.PdfRequest_Chunk{Chunk: fileData}}); err != nil {
		logx.Error("发送数据块失败：%v", err)
		return "", err
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		logx.Error("接收响应失败：%v", err)
		return "", err
	}
	if response.Error != "" {
		logx.Errorf("处理PDF文件失败：%s", response.Error)
		return "", errors.New(response.Error)
	}
	return response.Content, nil
}
