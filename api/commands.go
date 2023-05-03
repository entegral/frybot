package api

import (
	"os"

	"github.com/fatih/color"
)

// SaveConversation is a function that saves the conversation to a file
func SaveConversation(prompt string, context *ChatThread) {
	color.Blue("Saving conversation to frybot_conversation.md")
	err := os.WriteFile("frybot_conversation.md", []byte(context.String()), 0644)
	if err != nil {
		color.Red("Error saving conversation: %s", err)
	}
}

// DumpChatThread is a function that archives the chat thread by marshalling it into json before writing it to a file
func DumpChatThread(chatThread *ChatThread) {
	chatThreadJSON, err := chatThread.MarshalJSON()
	if err != nil {
		color.Red("Error archiving conversation: %s", err)
	}
	err = os.WriteFile("frybot_dump.md", []byte(chatThreadJSON), 0644)
	if err != nil {
		color.Red("Error archiving conversation: %s", err)
	}
}

// RecoverChatThread is a function that recovers the chat thread from a file
func RecoverChatThread(chat *ChatThread) *ChatThread {
	chatThread, err := os.ReadFile("frybot_dump.md")
	if err != nil {
		color.Red("Error recovering conversation: %s", err)
	}
	chat.UnmarshalJSON(chatThread)
	return chat
}
