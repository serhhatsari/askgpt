package chat

import (
	"bufio"
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/internal"
	"github.com/serhhatsari/askgpt/models"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Messages []models.ChatMessage

func AskGPT(cmd *cobra.Command, args []string) {

	setEnv()

	// Ask the user for a prompt until they Ctrl C
	for {
		getMessage()

		body := createBody()

		jsonBody := convertBodyToJSON(body)

		request := createRequest(jsonBody)

		response := sendRequest(request)

		printResponse(response)
	}

}

func setEnv() {
	// Check if the OPENAI_API_KEY environment variable is set
	if os.Getenv("OPENAI_API_KEY") == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		return
	}
	internal.OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

}

func getMessage() models.ChatMessage {
	var message string
	pterm.Print(pterm.Yellow("You: "))

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}

	var UserMessage models.ChatMessage
	UserMessage.Role = "user"
	UserMessage.Content = message
	Messages = append(Messages, UserMessage)

	return UserMessage
}

func createBody() models.ChatRequest {
	// Create the request body
	request := models.ChatRequest{
		Messages:    Messages,
		Model:       "gpt-3.5-turbo",
		MaxTokens:   2040,
		Temperature: 0,
	}
	return request
}

func convertBodyToJSON(request models.ChatRequest) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		pterm.Error.Println(err)
		return nil
	}
	return jsonBody
}

func createRequest(jsonBody []byte) *http.Request {
	// Create the HTTP request
	req, err := http.NewRequest("POST", internal.GPT_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+internal.OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendRequest(req *http.Request) models.ChatResponse {
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
	var response models.ChatResponse
	err = jsoniter.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response models.ChatResponse) {
	// Print the response
	result := response.Choices[0].Message.Content
	result = strings.TrimSpace(result)

	Messages = append(Messages, response.Choices[0].Message)

	pterm.Print(pterm.LightGreen("\nGPT-3: "))
	pterm.Println(result + "\n")
}
