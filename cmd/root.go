package cmd

import (
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
	Run:        AskGPT,
}

func Execute() error {

	cmdRoot.MarkPersistentFlagRequired("port")

	return cmdRoot.Execute()
}

func init() {

	viper.Set("Verbose", true)

	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("version", "0.0.1")
}
