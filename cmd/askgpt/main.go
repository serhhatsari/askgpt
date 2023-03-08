package main

import (
	"fmt"
	"github.com/serhhatsari/askgpt/internal/chat"
	"github.com/serhhatsari/askgpt/internal/completions"
	"github.com/serhhatsari/askgpt/internal/image"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdRoot = &cobra.Command{
	Use:   "askgpt",
	Short: "Simple CLI to interact with ChatGPT",
	Long:  "Simple CLI to interact with ChatGPT by wrapping the API provided by it.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		pterm.DefaultHeader.Println("Welcome to AskGPT!")
		pterm.Println(pterm.Blue("AskGPT is a CLI to interact with OpenAI API\n"))
	},
}

func main() {

	cmdRoot.MarkPersistentFlagRequired("port")
	cmdRoot.AddCommand(chat.CmdChat)
	cmdRoot.AddCommand(completions.CmdCompletion)
	cmdRoot.AddCommand(image.CmdImage)
	cmdRoot.Execute()
}

func init() {

	viper.Set("Verbose", true)

	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("version", "0.0.1")
}
