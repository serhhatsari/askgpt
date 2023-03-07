package cmd

import (
	"fmt"

	"github.com/serhhatsari/askgpt/internal/completions"
	"github.com/spf13/cobra"
)

var cmdCompletion = &cobra.Command{
	Use:     "cmp",
	Short:   "Given a prompt, the model will return one or more predicted completions.  ",
	Example: "askgpt cmp \"How do I make an HTTP request in Go?",
	Run:     completions.GetCompletion,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
		}
		return nil
	},
}
