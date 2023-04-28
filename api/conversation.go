package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// StartConversation starts a conversation with the bot in the terminal
// and saves the conversation to frybot_conversation.md when the user
// types "save this conversation"
func StartConversation() {
	var chat = ChatThread{}

	reader := bufio.NewReader(os.Stdin)
	color.Yellow("Welcome to your frybot conversation!\n")
	color.Yellow("Type 'save this conversation' to save the conversation to frybot_conversation.md\n")
	color.Yellow("Type 'exit' to exit the conversation\n")

	for {
		color.Blue("Enter text: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if strings.Contains(input, "save this conversation") {
			color.Blue("Saving conversation to frybot_conversation.md")
			err := os.WriteFile("frybot_conversation.md", []byte(chat.String()), 0644)
			if err != nil {
				color.Red("Error saving conversation: %s", err)
			}
			continue
		} else if input == "exit\n" {
			color.Blue("Exiting conversation")
			break
		}

		userMessage := NewChatMessage(RoleUser, input)
		systemMessage := NewChatMessage(RoleSystem, "You are an AI language model and your task is to assist the user while speaking like fry from futurama.\n\n")
		color.Blue("Prompt: %s", input)
		chat = append(chat, systemMessage, userMessage)
		response, err := SendChatRequest(NewRequestBody(GPT4, chat, 0.3))
		if err != nil {
			color.Red("Error: %s", err)
			break
		}
		if response.Error != nil {
			color.Red(response.String())
		} else {
			color.Green(response.String())
			chat = chat.Add(RoleSystem, response.String())
		}
	}
}
