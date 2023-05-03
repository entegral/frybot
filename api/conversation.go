package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func PromptForInput() (input string, breakLoop bool) {
	color.Blue("Enter text: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		breakLoop = true
	}
	return input, breakLoop
}

var reader = bufio.NewReader(os.Stdin)

// StartConversation starts a conversation with the bot in a for loop
// reading from stdin and sending the input to the OpenAI Chat API.
//
// The bot responds to the user's input with a response from the OpenAI Chat API
// and an ongoing conversation is saved to a ChatThread so it can be provided
// as context to the ongoing discussion.
//
// Commands:
// - "save this conversation" saves the conversation to frybot_conversation.md
// - "exit" exits the conversation
// and saves the conversation to frybot_conversation.md when the user
// types "save this conversation"
func StartConversation() {
	var chat = ChatThread{}

	color.Yellow("Welcome to your frybot conversation!\n")
	color.Yellow("Type 'save this conversation' to save the conversation to frybot_conversation.md\n")
	color.Yellow("Type 'exit' to exit the conversation\n")

	for {
		input, breakLoop := PromptForInput()
		if breakLoop {
			break
		}

		// frybot commands
		if strings.Contains(strings.ToLower(input), "save this conversation") {
			if input == "save this conversation\n" {
				color.Blue("Saving conversation")
				SaveConversation(input, &chat)
				continue
			}
			defer SaveConversation(input, &chat)
		} else if input == "exit\n" {
			SaveConversation(input, &chat)
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
