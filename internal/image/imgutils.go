package image

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/qeesung/image2ascii/convert"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var Size int = 512

func GetPrompt(args []string) string {
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

func CreateBody(prompt string) Request {
	if Size != 256 && Size != 512 && Size != 1024 {
		pterm.Error.Println("Size can be 256, 512 or 1024")
		pterm.Info.Println("Defaulting to 512")
		Size = 512
	}
	body := Request{
		Prompt:         prompt,
		Size:           strconv.Itoa(Size) + "x" + strconv.Itoa(Size),
		N:              1,
		ResponseFormat: "url",
	}
	return body
}

func ConvertBodyToJSON(request Request) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		pterm.Error.Println("Error converting request body to JSON")
		os.Exit(1)
	}
	return jsonBody
}

func ParseResponse(res []byte) Response {
	// Parse the response body
	var response Response
	err := jsoniter.Unmarshal(res, &response)
	if err != nil {
		pterm.Error.Println("Error parsing response")
		os.Exit(1)
	}

	return response
}

func PrintResponse(response Response) {

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

	PrintImage(filepath)
}

func PrintImage(filepath string) {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(filepath, &convertOptions))

	pterm.DefaultSection.Println("Image saved to " + filepath)
}
