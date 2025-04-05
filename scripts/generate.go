// scripts/generate.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message Message `json:"message"`
}

type ResponseBody struct {
	Choices []Choice `json:"choices"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	baseURL := "https://api.openai.com/v1/chat/completions"
	model := "gpt-3.5-turbo"

	if apiKey == "" {
		apiKey = os.Getenv("OPENROUTER_API_KEY")
		if apiKey == "" {
			fmt.Println("❌ OPENAI_API_KEY atau OPENROUTER_API_KEY belum diset!")
			return
		}
		baseURL = "https://openrouter.ai/api/v1/chat/completions"
		model = "deepseek-coder:free"
	}

	prompt := "Berikan 1 fakta unik dan menarik tentang bahasa pemrograman Golang."
	reqBody := RequestBody{
		Model: model,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", baseURL, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	if os.Getenv("OPENAI_API_KEY") != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	} else {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Gagal request:", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("❌ Gagal dari API:", string(respBody))
		return
	}

	var res ResponseBody
	json.Unmarshal(respBody, &res)

	content := res.Choices[0].Message.Content
	date := time.Now().Format("2006-01-02")

	f, _ := os.OpenFile("thoughts.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(fmt.Sprintf("\n## %s\n%s\n", date, content))

	fmt.Println("✅ Fakta berhasil ditulis ke thoughts.md")
}
