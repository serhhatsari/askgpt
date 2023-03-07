package cmd

import (
	"github.com/serhhatsari/askgpt/internal/image"
	"github.com/spf13/cobra"
)

var cmdImage = &cobra.Command{
	Use:     "image",
	Short:   "Create an image from a prompt using the Dall-E model.",
	Long:    "Create an image from a prompt using the Dall-E model.",
	Example: "askgpt image \"A drawing of a cat.\"",
	Run:     image.GenerateImage,
}
