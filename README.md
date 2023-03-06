## Description  
CLI to interact with OpenAI API.

## Getting Started
### Dependencies
* [OpenAI API Key](https://beta.openai.com/)  
You need to set the API key as an environment variable:  
`export OPENAI_API_KEY=<your_api_key>`
* [Go](https://golang.org/) 

### Installing  
Clone the repo by using the following command:     
`git clone https://github.com/serhhatsari/askgpt`

### Building the program
Build the CLI by using the following command:  
`go build -o askgpt`

### Usage
```shell
askgpt <command>

Available Commands:
  cmp             Given a prompt, the model will return one or more predicted completions.  
  chat            Chat with GPT-3  

Examples:
  askgpt cmp "How do I make a HTTP request in Go?"
  askgpt chat // This will start a chat session with GPT-3

```