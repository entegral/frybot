package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChatMessage is a struct that represents a message in the OpenAI Chat API
type ChatMessage struct {
	Role    ChatRole `json:"role"`
	Content string   `json:"content"`
}

// ChatThread is a slice of ChatMessages with special helper functions
type ChatThread []ChatMessage

func (ct ChatThread) String() string {
	var str string
	for _, message := range ct {
		str += fmt.Sprintf("%s: %s\n", message.Role, message.Content)
	}
	return str
}

// Add will add a new ChatMessage to the ChatThread
func (ct ChatThread) Add(role ChatRole, message string) ChatThread {
	return append(ct, NewChatMessage(role, message))
}

// RequestBody is the request body for the OpenAI Chat API
type RequestBody struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float32       `json:"temperature,omitempty"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
}

// NewChatMessage is a function that returns a new ChatMessage
func NewChatMessage(role ChatRole, content string) ChatMessage {
	return ChatMessage{Role: role, Content: content}
}

// NewRequestBody is a function that returns the request body for the OpenAI Chat API
func NewRequestBody(model Models, messages []ChatMessage, temperature float32) RequestBody {
	return RequestBody{
		Model:       string(model),
		Messages:    messages,
		Temperature: temperature,
	}
}

// SendChatRequest is a function that sends a request to the OpenAI Chat API
func SendChatRequest(requestBody RequestBody) (*ChatResponse, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", ev.ChatAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ev.OpenAIAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response ChatResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type ChatResponse struct {
	Error   *APIError    `json:"error,omitempty"`
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []ChatChoice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// String is a function that returns the text of the first ChatChoice
// if an error occurs, it returns the error
func (cr ChatResponse) String() string {
	if len(cr.Choices) == 0 {
		return "No choices returned"
	}
	if cr.Error != nil {
		return cr.Error.Message
	}
	return cr.Choices[0].Message.Content
}

type ChatChoice struct {
	Index        int                 `json:"index"`
	Message      ChatMessageResponse `json:"message"`
	FinishReason string              `json:"finish_reason"`
}

type ChatMessageResponse struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
