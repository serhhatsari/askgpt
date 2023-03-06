package completions

import (
	"bytes"
	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/internal"
	"github.com/serhhatsari/askgpt/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

func GetCompletion(cmd *cobra.Command, args []string) {

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
		panic("Please set the OPENAI_API_KEY environment variable")
	}
	internal.OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
}

func getPrompt(args []string) string {
	// Check if the user provided a prompt
	if len(args) != 1 {
		panic("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
	}

	// Check if the prompt is too long
	prompt := args[0]
	if len(prompt) > 2048 {
		panic("Prompt is too long, max length is 2048")
	}
	return prompt
}

func createBody(prompt string) models.CompletionRequest {
	body := models.CompletionRequest{
		Prompt:    prompt,
		Model:     "text-davinci-003",
		MaxTokens: 2040,
	}
	return body
}

func convertBodyToJSON(request models.CompletionRequest) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		panic(err)
	}
	return jsonBody
}

func createRequest(jsonBody []byte) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", internal.COMPLETIONS_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+internal.OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendRequest(req *http.Request) models.CompletionResponse {
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
	var response models.CompletionResponse
	err = jsoniter.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response models.CompletionResponse) {
	// Print the response
	result := response.Choices[0].Text
	result = strings.TrimSpace(result)

	pterm.Print(pterm.LightGreen("\nGPT-3: "))
	pterm.Println(result + "\n")
}
