package main

import (
	"fmt"
	"github.com/serhhatsari/askgpt/internal/chat"
	"github.com/serhhatsari/askgpt/internal/completions"
	"github.com/serhhatsari/askgpt/internal/image"
	"github.com/spf13/cobra/doc"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdRoot = &cobra.Command{
	Use:   "askgpt",
	Short: "Simple CLI to interact with OpenAI API",
	Long:  "Simple CLI to interact with OpenAI API and get answers to your questions or generate images.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		pterm.DefaultHeader.Println(pterm.Red("Welcome to AskGPT!"))
		pterm.Println(pterm.Blue("AskGPT is a CLI to interact with OpenAI API\n"))
	},
}

func main() {

	cmdRoot.MarkPersistentFlagRequired("port")

	cmdRoot.AddCommand(chat.CmdChat)
	chat.CmdChat.Flags().Float32VarP(&chat.Temperature, "temperature", "t", 0, "Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0")

	cmdRoot.AddCommand(completions.CmdCompletion)
	completions.CmdCompletion.Flags().Float32VarP(&completions.Temperature, "temperature", "t", 0, "Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0")

	cmdRoot.AddCommand(image.CmdImage)
	image.CmdImage.Flags().IntVarP(&image.Size, "size", "s", 256, "Size of the image to generate. Try 256, 512, 1024")

	err := doc.GenMarkdownTree(cmdRoot, "./docs")
	if err != nil {
		pterm.Error.Println(err)
	}

	cmdRoot.Execute()
}

func init() {

	viper.Set("Verbose", true)

	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("version", "0.0.1")
}
