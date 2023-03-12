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


## Usage

### `askgpt chat`
Start a chat session with ChatGPT.
```
askgpt chat [flags]
```
### Examples
```
askgpt chat # This will start a chat session with ChatGPT  
askgpt chat -t 0.5 # This will start a chat session with ChatGPT with a temperature of 0.5 
```
### Options
```
  -h, --help                  help for chat
  -t, --temperature float32   Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0
```
### `askgpt cmp`
Ask one thing to OpenAI's GPT-3 model and get a completion.
```
askgpt cmp <your_prompt> [flags]
```
### Examples
```
askgpt cmp "How do I make an HTTP request in Go?    
askgpt cmp "Who is Pedro Pascal?" -t 1.5  
```
### Options
```
  -h, --help                  help for cmp
  -t, --temperature float32   Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0
```
### `askgpt image`
Create an image from a prompt using the Dall-E model.
```
askgpt image <your_prompt> [flags]
```
### Examples
```
askgpt image "A drawing of a cat."  
askgpt image "Dog is driving a car." -s 1024  
```
### Options
```
  -h, --help       help for image
  -s, --size int   Size of the image to generate. Try 256, 512, 1024 (default 512)
```

## Demo

`$ askgpt chat`  
<img src="https://raw.githubusercontent.com/serhhatsari/askgpt/master/assets/chatusage.gif" width="70%" height="50%"/>

`$ askgpt cmp "<your_prompt>"`  
<img src="https://raw.githubusercontent.com/serhhatsari/askgpt/master/assets/cmpusage.gif" width="70%" height="50%"/>

`$ askgpt image "<your_prompt>"`  
<img src="https://raw.githubusercontent.com/serhhatsari/askgpt/master/assets/imageusage.gif" width="70%" height="50%"/>


## Contributions
Contributions to AskGPT are always welcome! If you find a bug or have an idea for a new feature, feel free to submit a pull request or open an issue on the GitHub repository.

## License
AskGPT is open-source software licensed under the MIT License.

