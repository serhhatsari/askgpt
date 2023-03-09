package openai

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pterm/pterm"
)

const (
	gptUrl         = "https://api.openai.com/v1/chat/completions"
	completionsUrl = "https://api.openai.com/v1/completions"
	imageUrl       = "https://api.openai.com/v1/images/generations"
)

var openaiApiKey string

func SendRequestToChatGPT(jsonBody []byte) []byte {
	setEnv()
	req := createGPTRequest(jsonBody)
	req = addHeaders(req)
	return sendRequest(req)
}

func SendRequestToCompletions(jsonBody []byte) []byte {
	setEnv()
	req := createCompletionsRequest(jsonBody)
	req = addHeaders(req)
	pterm.Info.Println("Request sent to OpenAI. Waiting for response...")
	res := sendRequest(req)
	pterm.Info.Println("Response received from OpenAI:")
	return res
}

func SendRequestToDallE(jsonBody []byte) []byte {
	setEnv()
	req := createDallERequest(jsonBody)
	req = addHeaders(req)
	pterm.Info.Println("Request sent to OpenAI. Waiting for response...")
	res := sendRequest(req)
	pterm.Info.Println("Response received from OpenAI:")
	return res
}

func setEnv() {
	openaiApiKey = os.Getenv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
}

func createGPTRequest(jsonBody []byte) *http.Request {

	// Create the HTTP request
	req, err := http.NewRequest("POST", gptUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	return req
}

func createCompletionsRequest(jsonBody []byte) *http.Request {

	// Create the HTTP request
	req, err := http.NewRequest("POST", completionsUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	return req
}

func createDallERequest(jsonBody []byte) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", imageUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	return req
}

func addHeaders(req *http.Request) *http.Request {
	// Set the headers
	req.Header.Set("Authorization", "Bearer "+openaiApiKey)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendRequest(req *http.Request) []byte {
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

	return respBody
}
