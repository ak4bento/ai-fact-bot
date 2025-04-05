package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func getFact(apiKey, baseURL string) (string, error) {
  prompt := os.Getenv("FACT_PROMPT")
	if prompt == "" {
		prompt = "Give me one interesting fact about AI in 1 sentence."
	}

	reqBody := map[string]interface{}{
		"model": "gpt-3.5-turbo", // atau "deepseek-chat" (dari OpenRouter)
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	bodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", baseURL+"/chat/completions", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OpenAIResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Choices[0].Message.Content, nil
}

func updateReadme(fact string) error {
	readme, err := os.ReadFile("README.md")
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(?s)(<!-- AI-FACT-START -->)(.*?)(<!-- AI-FACT-END -->)`)
	newContent := re.ReplaceAllString(string(readme), fmt.Sprintf("$1\n%s\n$3", fact))
	return os.WriteFile("README.md", []byte(newContent), 0644)
}

func main() {
	apiKey := os.Getenv("AI_API_KEY")
	baseURL := os.Getenv("AI_API_BASE_URL")

	fact, err := getFact(apiKey, baseURL)
	if err != nil {
		panic(err)
	}

	if err := updateReadme(fact); err != nil {
		panic(err)
	}

	fmt.Println("✅ README updated with new AI fact")
}
