package openai

import (
	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/internal/utils"
)

const (
	gptUrl   = "https://api.openai.com/v1/chat/completions"
	imageUrl = "https://api.openai.com/v1/images/generations"
)

var openaiApiKey string

func SendRequestToChatGPT(jsonBody []byte) []byte {
	openaiApiKey = utils.GetToken()
	req := createRequest(jsonBody, gptUrl)
	return sendRequest(req)
}

func SendRequestToDallE(jsonBody []byte) []byte {
	openaiApiKey = utils.GetToken()
	req := createRequest(jsonBody, imageUrl)
	pterm.Info.Println("Request sent to OpenAI. Waiting for response...")
	res := sendRequest(req)
	pterm.Info.Println("Response received from OpenAI:")
	return res
}
