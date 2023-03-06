package cmd

import (
	"github.com/serhhatsari/askgpt/internal/chat"
	"github.com/spf13/cobra"
)

var cmdChat = &cobra.Command{
	Use:        "chat",
	Aliases:    []string{"chatgpt"},
	SuggestFor: []string{"gpt", "askchatgpt"},
	Short:      "Simple CLI to interact with ChatGPT",
	Long:       "Simple CLI to interact with ChatGPT by wrapping the API provided by it.",
	Example:    "askgpt \"How do I make an HTTP request in Go?",
	Run:        chat.AskGPT,
}
