package cmd

import (
	"fmt"
	"github.com/serhhatsari/askgpt/internal/completions"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdRoot = &cobra.Command{
	Use:        "askgpt",
	Aliases:    []string{"chatgpt"},
	SuggestFor: []string{"gpt", "askchatgpt"},
	Short:      "Simple CLI to interact with ChatGPT",
	Long:       "Simple CLI to interact with ChatGPT by wrapping the API provided by it.",
	Example:    "askgpt \"How do I make an HTTP request in Go?",
	Run:        completions.GetCompletion,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
		}

		// Check if the prompt is too long
		prompt := args[0]
		if len(prompt) > 2048 {
			return fmt.Errorf("Prompt is too long, max length is 2048")
		}
		return nil
	},
}

func Execute() error {

	cmdRoot.MarkPersistentFlagRequired("port")
	cmdRoot.AddCommand(cmdChat)

	return cmdRoot.Execute()
}

func init() {

	viper.Set("Verbose", true)

	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("version", "0.0.1")
}
