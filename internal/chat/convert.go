package chat

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
)

func convertBodyToJSON(request Request) []byte {
	// Convert the request body to Byte Array
	jsonBody, err := jsoniter.Marshal(&request)
	if err != nil {
		pterm.Error.Println(err)
		return nil
	}
	return jsonBody
}
