package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		pterm.DefaultHeader.Println("Welcome to AskGPT!")
		pterm.Println(pterm.Blue("AskGPT is a CLI to interact with OpenAI API\n"))

	},
}

func Execute() error {

	cmdRoot.MarkPersistentFlagRequired("port")
	cmdRoot.AddCommand(cmdChat)
	cmdRoot.AddCommand(cmdCompletion)
	return cmdRoot.Execute()
}

func init() {

	viper.Set("Verbose", true)

	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("version", "0.0.1")
}
