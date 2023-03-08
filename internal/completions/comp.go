package completions

import (
	"github.com/serhhatsari/askgpt/pkg/openai"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var CmdCompletion = &cobra.Command{
	Use:     "cmp",
	Short:   "Given a prompt, the model will return one or more predicted completions.  ",
	Example: "askgpt cmp \"How do I make an HTTP request in Go?",
	Run:     GetCompletion,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			pterm.Error.Println("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
			os.Exit(1)
		}
		return nil
	},
}

func GetCompletion(cmd *cobra.Command, args []string) {

	prompt := getPrompt(args)

	body := createBody(prompt)

	jsonBody := convertBodyToJSON(body)

	res := openai.SendRequestToCompletions(jsonBody)

	parsedResponse := parseResponse(res)

	printResponse(parsedResponse)
}

func getPrompt(args []string) string {
	// Check if the user provided a prompt
	if len(args) != 1 {
		panic("Please provide a prompt, example: askgpt \"How do I make an HTTP request in Go?")
	}

	// Check if the prompt is too long
	prompt := args[0]
	if len(prompt) > 2048 {
		panic("Prompt is too long, max length is 2048")
	}
	return prompt
}

func createBody(prompt string) Request {
	body := Request{
		Prompt:    prompt,
		Model:     "text-davinci-003",
		MaxTokens: 2040,
	}
	return body
}

func convertBodyToJSON(request Request) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		panic(err)
	}
	return jsonBody
}

func parseResponse(respBody []byte) Response {
	// Parse the response body
	var response Response
	err := jsoniter.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func printResponse(response Response) {
	// Print the response
	result := response.Choices[0].Text
	result = strings.TrimSpace(result)

	pterm.Print(pterm.LightGreen("\nGPT-3: "))
	pterm.Println(result + "\n")
}
