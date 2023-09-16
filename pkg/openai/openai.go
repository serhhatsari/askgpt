package openai

import (
	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/internal/utils"
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
