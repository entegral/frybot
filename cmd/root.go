/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "frybot",
	Short: "A chatGPT bot that answers helpful questions about local files",
	Long: `frybot has two primary functions:
	1. prompt: a one-off question about local files
	2. chat: a chatbot that can answer multiple questions interactively and give you
		a conversation history to review and save.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fry.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// AddFrybotCommand binds the frybot root command to the parent command. Use this to integrate this package into your own CLI.
func AddFrybotCommand(parent *cobra.Command) {
	parent.AddCommand(RootCmd)
}
