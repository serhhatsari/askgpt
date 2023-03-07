## Description  
AskGPT is a CLI tool built in Go that allows you to interact with ChatGPT, Dall-E models trained by OpenAI.  
With this tool, you can easily ask ChatGPT for help with various tasks, from generating text to images.

## Getting Started
### Dependencies  
* [OpenAI API Key](https://platform.openai.com/account/api-keys)    
You need to generate an API key and export it as an env variable:    
`$ export OPENAI_API_KEY= <your_api_key>`

### Installation  
Via [HomeBrew](https://brew.sh/):   
`$ brew install serhhatsari/tools/askgpt`


### Usage
```shell
askgpt <command>

Available Commands:
  cmp             Ask one thing to OpenAI
  chat            Make a conversation with GPT-3  
  image           Create an image from a prompt using the Dall-E model.
  help            Help about any command

Examples:
  askgpt help  
  askgpt cmp "How do I make a HTTP request in Go?"
  askgpt chat // This will start a chat session with GPT-3
  askgpt image "Cats are playing football with dogs."

```
`$ askgpt chat`  
<img src="https://raw.githubusercontent.com/serhhatsari/askgpt/master/docs/assets/chatusage.gif" width="60%" height="50%"/>

`$ askgpt cmp "<your_prompt>"`  
<img src="https://raw.githubusercontent.com/serhhatsari/askgpt/master/docs/assets/cmpusage.gif" width="60%" height="50%"/>

`$ askgpt image "<your_prompt>"`  
<img src="https://raw.githubusercontent.com/serhhatsari/askgpt/master/docs/assets/imageusage.gif" width="60%" height="50%"/>


## Contributions
Contributions to AskGPT are always welcome! If you find a bug or have an idea for a new feature, feel free to submit a pull request or open an issue on the GitHub repository.

## License
AskGPT is open-source software licensed under the MIT License.

