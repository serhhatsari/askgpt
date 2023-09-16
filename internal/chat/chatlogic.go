package chat

import (
	"bufio"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
)

var (
	Messages    []Message
	Temperature float32 = 0
)

func CheckExit(message string) {
	message = strings.ToLower(message)
	if message == "exit" || message == "quit" {
		os.Exit(0)
	}
}

func GetMessage() Message {
	var message string
	pterm.Print(pterm.Yellow("You: "))

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}

	CheckExit(message)

	var UserMessage Message
	UserMessage.Role = "user"
	UserMessage.Content = message
	Messages = append(Messages, UserMessage)

	return UserMessage
}

func CreateBody() Request {
	if Temperature < 0 || Temperature > 2 {
		pterm.Error.Println("Temperature must be between 0 and 2")
		pterm.Info.Println("Setting temperature to 0 automatically")
		Temperature = 0
	}

	request := Request{
		Messages:    Messages,
		Model:       "gpt-4",
		MaxTokens:   2000,
		Temperature: Temperature,
	}
	return request
}

func PrintResponse(response Response) {
	if len(response.Choices) == 0 {
		pterm.Error.Println("Content length exceed. Start another conversation.")
		os.Exit(1)
	}
	result := response.Choices[0].Message.Content
	result = strings.TrimSpace(result)

	Messages = append(Messages, response.Choices[0].Message)

	pterm.Print(pterm.LightGreen("\nGPT-4: "))
	pterm.Println(result + "\n")
}

func ConvertBodyToJSON(request Request) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		pterm.Error.Println(err)
		return nil
	}
	return jsonBody
}

func ParseResponse(res []byte) Response {
	// Parse the response body
	var response Response

	err := jsoniter.Unmarshal(res, &response)

	if err != nil {
		pterm.Error.Println(err)
		return response
	}
	return response
}
