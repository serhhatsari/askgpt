package openai

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/pterm/pterm"
)

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
