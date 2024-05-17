package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Response struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

var conversationHistory []struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func ChatWithBot(userInput, role string) (string, error) {

	// 检查是否存在 Role 为 "user" 的条目
	foundUser := false

	for _, item := range conversationHistory {
		if item.Role == "user" {
			foundUser = true
			break // 找到后就不需要继续循环了
		}
	}

	if !foundUser && role == "assistant" {
		conversationHistory = append(conversationHistory, struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{Role: "assistant", Content: userInput})
		return "", nil
	} else {
		// 添加用户输入到对话历史记录
		conversationHistory = append(conversationHistory, struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{Role: "user", Content: userInput})

		// 构建请求数据
		url := "http://localhost:11434/api/chat"
		jsonData, err := json.Marshal(map[string]interface{}{
			"model":    globalSettings["model"],
			"messages": conversationHistory,
			"prompt":   globalSettings["prompt"],
		})
		if err != nil {
			return "", err
		}

		// 发送请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		// 读取响应
		var reply string
		dec := json.NewDecoder(resp.Body)
		for dec.More() { // 使用More()来检查是否还有更多数据可读
			var r Response
			err := dec.Decode(&r)
			if err != nil {
				return "", err
			}

			reply += r.Message.Content
		}

		conversationHistory = append(conversationHistory, struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{Role: "assistant", Content: reply})

		// 返回机器人的回复
		return reply, nil
	}

}
