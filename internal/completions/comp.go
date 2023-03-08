package completions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const COMPLETIONS_URL = "https://api.openai.com/v1/completions"

var OPENAI_API_KEY string

var CmdCompletion = &cobra.Command{
	Use:     "cmp",
	Short:   "Given a prompt, the model will return one or more predicted completions.  ",
	Example: "askgpt cmp \"How do I make an HTTP request in Go?",
	Run:     GetCompletion,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
		}
		return nil
	},
}

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
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
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

func createBody(prompt string) CompletionRequest {
	body := CompletionRequest{
		Prompt:    prompt,
		Model:     "text-davinci-003",
		MaxTokens: 2040,
	}
	return body
}

func convertBodyToJSON(request CompletionRequest) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		panic(err)
	}
	return jsonBody
}

func createRequest(jsonBody []byte) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", COMPLETIONS_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendRequest(req *http.Request) CompletionResponse {
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
	var response CompletionResponse
	err = jsoniter.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response CompletionResponse) {
	// Print the response
	result := response.Choices[0].Text
	result = strings.TrimSpace(result)

	pterm.Print(pterm.LightGreen("\nGPT-3: "))
	pterm.Println(result + "\n")
}
