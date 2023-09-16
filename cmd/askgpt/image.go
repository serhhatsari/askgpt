package askgpt

import (
	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/internal/image"
	"github.com/serhhatsari/askgpt/pkg/openai"
	"github.com/spf13/cobra"
	"os"
)

func checkArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		pterm.Error.Println("Please provide a prompt, example: askgpt image \"A drawing of a cat.\"")
		os.Exit(1)
	}
	return nil
}

func GenerateImage(cmd *cobra.Command, args []string) {
	prompt := image.GetPrompt(args)
	body := image.CreateBody(prompt)
	jsonBody := image.ConvertBodyToJSON(body)
	res := openai.SendRequestToDallE(jsonBody)
	parsedResponse := image.ParseResponse(res)
	image.PrintResponse(parsedResponse)
}

var cmdImage = &cobra.Command{
	Use:     "image",
	Short:   "Create an image from a prompt using the Dall-E model.",
	Example: "askgpt image \"A drawing of a cat.\"",
	Run:     GenerateImage,
	Args:    checkArgs,
}
