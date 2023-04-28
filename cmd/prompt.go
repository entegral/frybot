/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	frybot "github.com/entegral/frybot/api"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// promptCmd represents the frybot command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "prompt chatGPT bot with a one-off question about local files",
	Long: `frybot's prompt command is a chatGPT bot that answers helpful questions about local files.
	It is a work in progress but will eventually be able to answer questions like:
	"How do I find a User record using platformservices?" or,
	"What is the composite key structure for a Group?"`,
	Run: func(cmd *cobra.Command, args []string) {
		switch modelInput {
		case "davinci":
			model = frybot.Davinci
		case "gpt3.5":
			model = frybot.GPT3Turbo
		case "gpt4":
			model = frybot.GPT4
		default:
			model = frybot.GPT3Turbo
		}
		if prompt == "" {
			if args[0] == "" {
				color.Red("you must provide a prompt to frybot")
				return
			}
			prompt = args[0]
		}
		color.Yellow("using modelInput: %s", modelInput)
		color.Yellow("using model: %s", string(model))
		color.Yellow("using targetFile: %s", targetFile)
		color.Blue("Prompt %s", prompt)

		if targetFile != "" {
			workingDir, err := os.Getwd()
			if err != nil {
				color.Red("could not get working directory")
				return
			}
			targetFile = workingDir + "/" + targetFile
			PromptAboutFile(prompt, model, targetFile)
			return
		}
		PromptAboutWorkingDirectory(prompt, model)
	},
}

var targetFile string
var prompt string
var modelInput string
var model frybot.Models
var saveOutput bool

// AddParent initializes the command and adds it as a parent to the root command of this directory
func AddParent(parent *cobra.Command) {
	promptCmd.PersistentFlags().StringVarP(&targetFile, "targetFile", "t", "", "(optional) add a file to provide context to the prompt")
	promptCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "(required) prompt for the bot to answer")
	promptCmd.PersistentFlags().StringVarP(&modelInput, "model", "m", "", "(optional) model used for processing the prompt, default is gpt3.5")
	promptCmd.PersistentFlags().BoolVarP(&saveOutput, "saveOutput", "s", false, "(optional) model used for processing the prompt, default is gpt3.5")
	parent.AddCommand(promptCmd)
}
