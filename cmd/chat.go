/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/entegral/frybot/api"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "this command will initialize a chat session with the bot. use ctrl + c to exit.",
	Long: `the chat command will initialize a chat session with the bot. Your message context 
	is saved between messages, so you can ask a question and then ask a follow-up question, but
	it does not retain context between sessions.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StartConversation()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
