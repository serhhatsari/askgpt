package askgpt

import (
	"github.com/serhhatsari/askgpt/internal/chat"
	utils2 "github.com/serhhatsari/askgpt/internal/utils"
	"github.com/serhhatsari/askgpt/pkg/openai"
	"github.com/spf13/cobra"
)

var CmdChat = &cobra.Command{
	Use:     "chat",
	Short:   "Start a chat session with ChatGPT.",
	Long:    "Start a chat session with ChatGPT. Talk however you want, ChatGPT will respond.",
	Example: "askgpt chat",
	Run:     chatWithGPT,
}

func chatWithGPT(cmd *cobra.Command, args []string) {

	utils2.CheckToken()

	utils2.PrintDescription()

	for {
		chat.GetMessage()

		body := chat.CreateBody()

		jsonBody := chat.ConvertBodyToJSON(body)

		res := openai.SendRequestToChatGPT(jsonBody)

		parsedResponse := chat.ParseResponse(res)

		chat.PrintResponse(parsedResponse)

	}

}