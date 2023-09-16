package chat

import (
	"bufio"
	"github.com/serhhatsari/askgpt/utils"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/serhhatsari/askgpt/pkg/openai"
	"github.com/spf13/cobra"
)

var (
	Messages    []Message
	Temperature float32 = 0
)

var CmdChat = &cobra.Command{
	Use:     "chat",
	Short:   "Start a chat session with ChatGPT.",
	Long:    "Start a chat session with ChatGPT. Talk however you want, ChatGPT will respond.",
	Example: "askgpt chat",
	Run:     chatWithGPT,
}

func chatWithGPT(cmd *cobra.Command, args []string) {

	utils.CheckToken()

	utils.PrintDescription()

	for {
		getMessage()

		body := createBody()

		jsonBody := convertBodyToJSON(body)

		res := openai.SendRequestToChatGPT(jsonBody)

		parsedResponse := parseResponse(res)

		printResponse(parsedResponse)

	}

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

func createBody() Request {
	if Temperature < 0 || Temperature > 2 {
		pterm.Error.Println("Temperature must be between 0 and 2")
		pterm.Info.Println("Setting temperature to 0 automatically")
		Temperature = 0
	}

	request := Request{
		Messages:    Messages,
		Model:       "gpt-3.5-turbo",
		MaxTokens:   2000,
		Temperature: Temperature,
	}
	return request
}

func convertBodyToJSON(request Request) []byte {
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
		pterm.Error.Println(err)
		return response
	}
	return response
}

func printResponse(response Response) {
	if len(response.Choices) == 0 {
		pterm.Error.Println("Content length exceed. Start another conversation.")
		os.Exit(1)
	}
	result := response.Choices[0].Message.Content
	result = strings.TrimSpace(result)

	Messages = append(Messages, response.Choices[0].Message)

	pterm.Print(pterm.LightGreen("\nGPT-3: "))
	pterm.Println(result + "\n")
}
