package image

type Request struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type Data struct {
	Url     string `json:"url"`
	B64Json string `json:"b64_json"`
}

type Response struct {
	Created int64  `json:"created"`
	Data    []Data `json:"data"`
}
