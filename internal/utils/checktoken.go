package utils

import (
	"os"

	"github.com/pterm/pterm"
)

func GetToken() string {
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	if openAiApiKey == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
	return openAiApiKey
}
