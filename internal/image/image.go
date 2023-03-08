package image

import (
	"fmt"
	"github.com/serhhatsari/askgpt/pkg/openai"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"
)

var CmdImage = &cobra.Command{
	Use:     "image",
	Short:   "Create an image from a prompt using the Dall-E model.",
	Long:    "Create an image from a prompt using the Dall-E model.",
	Example: "askgpt image \"A drawing of a cat.\"",
	Run:     GenerateImage,
}

func GenerateImage(cmd *cobra.Command, args []string) {

	prompt := getPrompt(args)

	body := createBody(prompt)

	jsonBody := convertBodyToJSON(body)

	res := openai.SendRequestToDallE(jsonBody)

	parsedResponse := parseResponse(res)

	printResponse(parsedResponse)

}

func getPrompt(args []string) string {
	// Check if the user provided a prompt
	if len(args) != 1 {
		pterm.Error.Println("Please provide a prompt, example: askgpt image \"A drawing of a cat.\"")
		os.Exit(1)
	}

	// Check if the prompt is too long
	prompt := args[0]
	if len(prompt) > 2048 {
		pterm.Error.Println("Prompt is too long, max length is 2048")
		os.Exit(1)
	}
	return prompt
}

func createBody(prompt string) Request {
	body := Request{
		Prompt:         prompt,
		Size:           "1024x1024",
		N:              1,
		ResponseFormat: "url",
	}
	return body
}

func convertBodyToJSON(request Request) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		panic(err)
	}
	return jsonBody
}

func parseResponse(res []byte) Response {
	// Parse the response body
	var response Response
	err := jsoniter.Unmarshal(res, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response Response) {
	ImageUrl := response.Data[0].Url

	filename := strconv.Itoa(int(response.Created)) + ".png"

	// Create the images directory if it doesn't exist
	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		os.Mkdir("./images", 0755)
	}

	filepath := filepath.Join("./images/", filename)

	// Create the file
	out, err := os.Create(filepath)

	// Get the image
	resp, err := http.Get(ImageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Download the image
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	printImage(filepath)
}

func printImage(filepath string) {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(filepath, &convertOptions))

	pterm.DefaultSection.Println("Image saved to " + filepath)
}
