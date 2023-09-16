package askgpt

import (
	"log/slog"

	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/internal/chat"
	"github.com/serhhatsari/askgpt/internal/image"
	"github.com/spf13/cobra"
)

func welcomeMessage(cmd *cobra.Command, args []string) {
	pterm.DefaultHeader.Println(pterm.Red("Welcome to AskGPT!"))
	pterm.Println(pterm.Blue("AskGPT is a CLI to interact with OpenAI API\n"))
}

var CmdRoot = &cobra.Command{
	Use:   "askgpt",
	Short: "Simple CLI to interact with OpenAI API",
	Long:  "Simple CLI to interact with OpenAI API and get answers to your questions or generate images.",
	Run:   welcomeMessage,
}

func Execute() {

	CmdRoot.AddCommand(CmdChat)
	CmdChat.Flags().Float32VarP(&chat.Temperature, "temperature", "t", 0, "Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0")

	CmdRoot.AddCommand(cmdImage)
	cmdImage.Flags().IntVarP(&image.Size, "size", "s", 256, "Size of the image to generate. Try 256, 512, 1024")

	err := CmdRoot.Execute()
	if err != nil {
		slog.Error("Error executing command: %v", err)
	}
}
