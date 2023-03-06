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
  cmp             Ask one thing to OpenAI
  chat            Make a conversation with GPT-3  
  help            Help about any command

Examples:
  askgpt help  
  askgpt cmp "How do I make a HTTP request in Go?"
  askgpt chat // This will start a chat session with GPT-3

```
## Demo 
![Usage](https://raw.githubusercontent.com/serhhatsari/askgpt/master/assets/usage.gif)
