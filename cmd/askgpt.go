package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

type request struct {
	Prompt      string `json:"prompt"`
	Model       string `json:"model"`
	Max_tokens  int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
}

const URL = "https://api.openai.com/v1/completions"

func AskGPT(cmd *cobra.Command, args []string) {

	// Check if the OPENAI_API_KEY environment variable is set
	if os.Getenv("OPENAI_API_KEY") == "" {
		panic("Please set the OPENAI_API_KEY environment variable")
	}

	// Check if the user provided a prompt
	if len(args) != 1 {
		panic("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
	}

	// Check if the prompt is too long
	prompt := args[0]
	if len(prompt) > 2048 {
		panic("Prompt is too long, max length is 2048")
	}

	// Create the request body
	body := request{
		Prompt:     prompt,
		Model:      "text-davinci-003",
		Max_tokens: 2040,
	}

	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&body)
	if err != nil {
		panic(err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	// Set the headers
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")

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
	response := make(map[string]interface{})
	err = jsoniter.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	// Print the response
	result := response["choices"].([]interface{})[0].(map[string]interface{})["text"]
	if result == "" {
		fmt.Println("No response")
		return
	}
	fmt.Println(result)
}
