package openai

import (
	"bytes"
	"github.com/serhhatsari/askgpt/internal/utils"
	"io"
	"net/http"
	"os"

	"github.com/pterm/pterm"
)

const (
	gptUrl   = "https://api.openai.com/v1/chat/completions"
	imageUrl = "https://api.openai.com/v1/images/generations"
)

var openaiApiKey string

func SendRequestToChatGPT(jsonBody []byte) []byte {
	utils.CheckToken()
	req := createRequest(jsonBody, gptUrl)
	return sendRequest(req)
}

func SendRequestToDallE(jsonBody []byte) []byte {
	utils.CheckToken()
	req := createRequest(jsonBody, imageUrl)
	pterm.Info.Println("Request sent to OpenAI. Waiting for response...")
	res := sendRequest(req)
	pterm.Info.Println("Response received from OpenAI:")
	return res
}

func createRequest(jsonBody []byte, url string) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		pterm.Error.Println("Error creating request")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+openaiApiKey)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendRequest(req *http.Request) []byte {
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		pterm.Error.Println("Error occurred while sending request to ChatGPT")
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		pterm.Error.Println("Error occurred while reading response body")
		os.Exit(1)
	}

	return respBody
}
