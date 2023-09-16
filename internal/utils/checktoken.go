package utils

import (
	"os"

	"github.com/pterm/pterm"
)

func CheckToken() {
	OpenaiApiKey := os.Getenv("OPENAI_API_KEY")
	if OpenaiApiKey == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
}
