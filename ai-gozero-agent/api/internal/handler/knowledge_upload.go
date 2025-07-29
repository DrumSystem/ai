package handler

import (
	"ai-gozero-agent/api/internal/logic"
	"ai-gozero-agent/api/internal/svc"
	"ai-gozero-agent/api/internal/types"
	"ai-gozero-agent/api/internal/utils"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func KnowledgeUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: add your logic here and delete this line
		setSSEHeader(w)
		fmt.Println("进入上传知识库！！！")

		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		defer file.Close()

		// 验证PDF
		if header.Header.Get("Content-Type") != "application/pdf" {
			httpx.Error(w, errors.New("仅支持PDF文件"))
			return
		}

		// 提取文本
		content, err := utils.ExtractPDFText(file)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		title := header.Filename
		fmt.Println("标题：", title)
		l := logic.NewKnowledgeUploadLogic(r.Context(), svcCtx)
		resp, err := l.KnowledgeUpload(&types.KnowledgeUploadReq{
			Title:   title,
			Content: content,
		})
		if err != nil {
			httpx.Error(w, err)
			return
		} else {
			httpx.OkJson(w, resp)
		}

	}
}
