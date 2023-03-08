package chat

import (
	"bufio"
	"github.com/serhhatsari/askgpt/pkg/openai"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var Messages []Message

var CmdChat = &cobra.Command{
	Use:     "chat",
	Short:   "Chat with ChatGPT",
	Long:    "Chat with ChatGPT model and get answers to your questions.",
	Example: "askgpt chat",
	Run:     AskGPT,
}

func AskGPT(cmd *cobra.Command, args []string) {

	checkToken()

	printDescription()

	for {
		getMessage()

		body := createBody()

		jsonBody := convertBodyToJSON(body)

		res := openai.SendRequestToChatGPT(jsonBody)

		parsedResponse := parseResponse(res)

		printResponse(parsedResponse)

	}

}

func checkToken() {
	OpenaiApiKey := os.Getenv("OPENAI_API_KEY")
	if OpenaiApiKey == "" {
		pterm.Error.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}
}

func printDescription() {
	pterm.DefaultHeader.Println("Welcome to AskGPT!")
	pterm.Println(pterm.Blue("AskGPT is a CLI to interact with ChatGPT"))
	pterm.Println(pterm.Blue("You can ask any question and ChatGPT will answer it."))
	pterm.Println(pterm.Red("To exit, type \"exit\" or \"quit\" or press Ctrl C.\n"))
}

func checkExit(message string) {
	message = strings.ToLower(message)
	if message == "exit" || message == "quit" {
		os.Exit(0)
	}

}

func getMessage() Message {
	var message string
	pterm.Print(pterm.Yellow("You: "))

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}

	checkExit(message)

	var UserMessage Message
	UserMessage.Role = "user"
	UserMessage.Content = message
	Messages = append(Messages, UserMessage)

	return UserMessage
}

func createBody() ChatRequest {
	// Create the request body
	request := ChatRequest{
		Messages:    Messages,
		Model:       "gpt-3.5-turbo",
		MaxTokens:   2040,
		Temperature: 0,
	}
	return request
}

func convertBodyToJSON(request ChatRequest) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		pterm.Error.Println(err)
		return nil
	}
	return jsonBody
}

func parseResponse(res []byte) Response {
	// Parse the response body
	var response Response

	err := jsoniter.Unmarshal(res, &response)

	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response Response) {
	// Print the response
	result := response.Choices[0].Message.Content
	result = strings.TrimSpace(result)

	Messages = append(Messages, response.Choices[0].Message)

	pterm.Print(pterm.LightGreen("\nGPT-3: "))
	pterm.Println(result + "\n")
}
