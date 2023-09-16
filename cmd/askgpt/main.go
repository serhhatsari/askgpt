package main

import (
	"github.com/joho/godotenv"
	"github.com/serhhatsari/askgpt/internal/chat"
	"github.com/serhhatsari/askgpt/internal/completions"
	"github.com/serhhatsari/askgpt/internal/image"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func welcomeMessage(cmd *cobra.Command, args []string) {
	pterm.DefaultHeader.Println(pterm.Red("Welcome to AskGPT!"))
	pterm.Println(pterm.Blue("AskGPT is a CLI to interact with OpenAI API\n"))
}

var cmdRoot = &cobra.Command{
	Use:   "askgpt",
	Short: "Simple CLI to interact with OpenAI API",
	Long:  "Simple CLI to interact with OpenAI API and get answers to your questions or generate images.",
	Run:   welcomeMessage,
}

func main() {

	cmdRoot.AddCommand(chat.CmdChat)
	chat.CmdChat.Flags().Float32VarP(&chat.Temperature, "temperature", "t", 0, "Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0")

	cmdRoot.AddCommand(completions.CmdCompletion)
	completions.CmdCompletion.Flags().Float32VarP(&completions.Temperature, "temperature", "t", 0, "Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0")

	cmdRoot.AddCommand(image.CmdImage)
	image.CmdImage.Flags().IntVarP(&image.Size, "size", "s", 256, "Size of the image to generate. Try 256, 512, 1024")

	err := cmdRoot.Execute()
	if err != nil {
		pterm.Error.Println(err)
		return
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		// Handle errors if the .env file cannot be loaded
		panic("Error loading .env file")
	}
}
