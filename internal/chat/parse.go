package chat

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
)

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
