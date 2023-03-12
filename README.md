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
## askgpt

Simple CLI to interact with OpenAI API

### Synopsis

Simple CLI to interact with OpenAI API and get answers to your questions or generate images.

```
askgpt [flags]
```

### Options

```
  -h, --help   help for askgpt
```

## askgpt chat

Chat with ChatGPT

```
askgpt chat [flags]
```

### Examples

```
askgpt chat
```

### Options

```
  -h, --help                  help for chat
  -t, --temperature float32   Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0
```
## askgpt cmp

Given a prompt, the model will return one or more predicted completions.

```
askgpt cmp [flags]
```

### Examples

```
askgpt cmp "How do I make an HTTP request in Go?
```

### Options

```
  -h, --help                  help for cmp
  -t, --temperature float32   Temperature of the model. Higher values will result in more creative completions, but also more likelihood of nonsensical text. Try 0, 0.5, 1.0, 1.5, 2.0
```

## askgpt image

Create an image from a prompt using the Dall-E model.

```
askgpt image [flags]
```

### Examples

```
askgpt image "A drawing of a cat."
```

### Options

```
  -h, --help       help for image
  -s, --size int   Size of the image to generate. Try 256, 512, 1024 (default 256)
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

