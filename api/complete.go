package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CompletionRequest is the request body for the OpenAI Completion API
type CompletionRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
	TopP        int    `json:"top_p"`
	N           int    `json:"n"`
	Stream      bool   `json:"stream"`
	Logprobs    any    `json:"logprobs"`
	Stop        string `json:"stop"`
}

// CompletionChoice is the response body for the OpenAI Completion API
type CompletionChoice struct {
	Text         string          `json:"text"`
	Index        int             `json:"index"`
	Logprobs     any             `json:"logprobs"`
	FinishReason string          `json:"finish_reason"`
	Usage        CompletionUsage `json:"usage"`
}

// CompletionUsage is part of the response body for the OpenAI Completion API
type CompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CompletionResponse is the response body for the OpenAI Completion API
type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int                `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
}

// Complete is a function that completes a CompletionRequest
func Complete(cr CompletionRequest) (CompletionResponse, error) {

	apiURL := "https://api.openai.com/v1/completions"

	requestBody, err := json.Marshal(cr)
	if err != nil {
		return CompletionResponse{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return CompletionResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ev.OpenAIAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return CompletionResponse{}, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CompletionResponse{}, err
	}
	logrus.Println("responseBody:", string(responseBody))

	var completionResponse CompletionResponse
	err = json.Unmarshal(responseBody, &completionResponse)
	if err != nil {
		return CompletionResponse{}, err
	}
	if len(completionResponse.Choices) == 0 {
		logrus.Println("No choices returned:", completionResponse)
		return completionResponse, nil
	}
	return completionResponse, nil
}
