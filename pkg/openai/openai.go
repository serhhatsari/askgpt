package openai

import (
	"bytes"
	"github.com/pterm/pterm"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	GPT_URL         = "https://api.openai.com/v1/chat/completions"
	COMPLETIONS_URL = "https://api.openai.com/v1/completions"
	IMAGE_URL       = "https://api.openai.com/v1/images/generations"
)

var OPENAI_API_KEY string

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
	return sendRequest(req)
}

func SendRequestToDallE(jsonBody []byte) []byte {
	setEnv()
	req := createDallERequest(jsonBody)
	req = addHeaders(req)
	return sendRequest(req)
}

func setEnv() {
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
	if OPENAI_API_KEY == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
}

func createGPTRequest(jsonBody []byte) *http.Request {

	// Create the HTTP request
	req, err := http.NewRequest("POST", GPT_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	return req
}

func createCompletionsRequest(jsonBody []byte) *http.Request {

	// Create the HTTP request
	req, err := http.NewRequest("POST", COMPLETIONS_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	return req
}

func createDallERequest(jsonBody []byte) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", IMAGE_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	return req
}

func addHeaders(req *http.Request) *http.Request {
	// Set the headers
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)
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
