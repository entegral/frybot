package api

import "fmt"

// AnalyzeCode is a function that completes a CompletionRequest
func AnalyzeCode(model Models, prompt, context string, temp float32) (*ChatResponse, error) {
	userMessage := NewChatMessage(RoleUser, prompt)
	systemMessages := NewChatMessage(RoleSystem, "You are an AI language model and your task is to assist the user with answering questions about the provided code. the provided code may be a combination of more than one file. The start of each file begins with '$!FILENAME=' + the filename, and the file ends with 'EOF'")
	codeContext := NewChatMessage(RoleSystem, fmt.Sprintf("the provided code:\n%s", context))
	requestBody := NewRequestBody(model, []ChatMessage{systemMessages, codeContext, userMessage}, temp)
	return SendChatRequest(requestBody)
}
