package main

import (
	"ai-gozero-agent/api/internal/types"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	req := types.InterviewAPPChatReq{
		Message: "go语言的channel是怎么实现的",
		ChatId:  "1",
	}

	params := url.Values{}
	params.Add("message", req.Message)
	params.Add("chatId", req.ChatId)

	baseUrl := "http://localhost:8123/interview_app/chat/sse"
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	request, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return
	}

	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(body))
}
