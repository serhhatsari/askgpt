package image

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"
)

const IMAGE_URL = "https://api.openai.com/v1/images/generations"

var OPENAI_API_KEY string

var CmdImage = &cobra.Command{
	Use:     "image",
	Short:   "Create an image from a prompt using the Dall-E model.",
	Long:    "Create an image from a prompt using the Dall-E model.",
	Example: "askgpt image \"A drawing of a cat.\"",
	Run:     GenerateImage,
}

func GenerateImage(cmd *cobra.Command, args []string) {

	setToken()

	prompt := getPrompt(args)

	body := createBody(prompt)

	jsonBody := convertBodyToJSON(body)

	req := createRequest(jsonBody)

	response := sendRequest(req)

	printResponse(response)

}

func setToken() {
	// Check if the OPENAI_API_KEY environment variable is set
	if os.Getenv("OPENAI_API_KEY") == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
}

func getPrompt(args []string) string {
	// Check if the user provided a prompt
	if len(args) != 1 {
		panic("Please provide a prompt, example: askgpt image \"A drawing of a cat.\"")
	}

	// Check if the prompt is too long
	prompt := args[0]
	if len(prompt) > 2048 {
		panic("Prompt is too long, max length is 2048")
	}
	return prompt
}

func createBody(prompt string) ImageRequest {
	body := ImageRequest{
		Prompt:         prompt,
		Size:           "1024x1024",
		N:              1,
		ResponseFormat: "url",
	}
	return body
}

func convertBodyToJSON(request ImageRequest) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		panic(err)
	}
	return jsonBody
}

func createRequest(jsonBody []byte) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", IMAGE_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendRequest(req *http.Request) ImageResponse {
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the response body
	var response ImageResponse
	err = jsoniter.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response ImageResponse) {
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
