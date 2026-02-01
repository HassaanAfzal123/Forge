package ai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaStreamLine struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func GenerateStream(ctx context.Context, prompt string) (string, error) {
	reqBody := ollamaRequest{
		Model:  "mistral",
		Prompt: prompt,
		Stream: true,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"http://localhost:11434/api/generate",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama error: %s", resp.Status)
	}

	var fullResponse strings.Builder
	scanner := bufio.NewScanner(resp.Body)

	fmt.Println("AI response:\n")

	for scanner.Scan() {
		var sLine ollamaStreamLine
		if err := json.Unmarshal(scanner.Bytes(), &sLine); err != nil {
			continue
		}

		fmt.Print(sLine.Response)
		fullResponse.WriteString(sLine.Response)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Println()
	return fullResponse.String(), nil
}



