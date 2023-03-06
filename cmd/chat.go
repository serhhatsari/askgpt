package cmd

import (
	"github.com/serhhatsari/askgpt/internal/chat"
	"github.com/spf13/cobra"
)

var cmdChat = &cobra.Command{
	Use:        "chat",
	Short:      "Chat with ChatGPT",
	Long:       "Chat with ChatGPT model and get answers to your questions.",
	Example:    "askgpt chat",
	Run:        chat.AskGPT,
}
