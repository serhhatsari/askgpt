package cmd

import (
	"fmt"
	"github.com/serhhatsari/askgpt/internal/completions"
	"github.com/spf13/cobra"
)

var cmdCompletion = &cobra.Command{
	Use:     "cmp",
	Short:   "Simple CLI to interact with ChatGPT",
	Long:    "Simple CLI to interact with ChatGPT by wrapping the API provided by it.",
	Example: "askgpt \"How do I make an HTTP request in Go?",
	Run:     completions.GetCompletion,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
		}
		return nil
	},
}
