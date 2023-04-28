package api

// ChatRole is a string that represents the role of a message in the OpenAI Chat API
type ChatRole string

const (
	// RoleUser is the role of the user in the OpenAI Chat API
	RoleUser ChatRole = "user"
	// RoleSystem is the role of the system in the OpenAI Chat API
	RoleSystem ChatRole = "system"
	// RoleAssistant is the role of the assistant in the OpenAI Chat API
	RoleAssistant ChatRole = "assistant"
)
