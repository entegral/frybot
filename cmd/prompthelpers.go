package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	frybot "github.com/entegral/frybot/api"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func FindMatchingFiles(pattern string) ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("dude, can't get the current working directory: %v", err)
	}

	files, err := filepath.Glob(filepath.Join(cwd, pattern))
	if err != nil {
		return nil, fmt.Errorf("man, there's something wrong with the pattern: %v", err)
	}

	var matchingFiles []string
	for _, file := range files {
		if !strings.Contains(file, "*") {
			matchingFiles = append(matchingFiles, file)
		}
	}
	color.Yellow("using targetFiles as context:")
	for _, file := range matchingFiles {
		color.Yellow(strings.Replace(file, cwd, " - .", 1))
	}
	return matchingFiles, nil
}

// PromptAboutFile prompts the bot about a file
func PromptAboutFile(prompt string, model frybot.Models, filePaths []string) {
	context := ""
	for _, filePath := range filePaths {
		data, _ := os.ReadFile(filePath)

		// input := "a cyborg version of fry from futurama"
		if context != "" {
			context += "\n"
		}
		context += "$!FILENAME=" + filePath + "\n" + string(data) + "\nEOF\n"
	}
	response, err := frybot.AnalyzeCode(model, prompt, context, 0.3)
	if err != nil {
		logrus.Println("Error:", err)
		return
	}
	fmt.Printf("\n")
	respString := response.String()
	if respString == "No choices returned" {
		response.LogAPIResponse()
	}
	color.Green(respString)
}

func IsIgnorableFile(file fs.DirEntry) bool {
	if file.IsDir() {
		return true
	}
	if file.Name() == "frybot_conversation.md" || file.Name() == "frybot_dump.md" {
		return true
	}
	if file.Name() == ".git" || file.Name() == ".gitignore" {
		return true
	}
	if file.Name() == "go.mod" || file.Name() == "go.sum" {
		return true
	}
	if strings.Contains(file.Name(), ".png") || strings.Contains(file.Name(), ".jpg") {
		return true
	}
	return false
}

// PromptAboutWorkingDirectory prompts the bot about the working directory
func PromptAboutWorkingDirectory(prompt string, model frybot.Models) {
	workingDir, err := os.Getwd()
	if err != nil {
		color.Red("could not get working directory")
		return
	}
	data, _ := os.ReadDir(workingDir)
	// use directory contents to parse and combine all files into a single prompt
	var combinedData string
	for _, file := range data {
		if IsIgnorableFile(file) {
			continue
		}
		fileData, err := os.ReadFile(file.Name())
		if err != nil {
			color.Red("could not read file %s", file.Name())
			return
		}
		contents := "$!FILENAME=" + file.Name() + "\n" + string(fileData) + "\nEOF\n"
		combinedData = combinedData + contents
	}
	response, err := frybot.AnalyzeCode(model, prompt, combinedData, 0.3)
	if err != nil {
		logrus.Println("Error:", err)
		return
	}
	color.Green(response.String())
	if response.String() == "No choices returned" {
		response.LogAPIResponse()
	}
	outputWithHeader := fmt.Sprintf("Prompt: %s\n\nResponse:\n%s", prompt, response.String())
	if saveOutput {
		if filename == "" {
			filename = "frybot_output.md"
		}
		err := os.WriteFile(filename, []byte(outputWithHeader), 0644)
		if err != nil {
			logrus.Println("Error:", err)
		}
	}
}
