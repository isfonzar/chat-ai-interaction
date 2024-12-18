# chat-ai-interaction
From an exported chat, interact with the chat using AI.

- Get a summary of the conversation in a date ("On January 1st, members of the project discussed the implementation of ...")
- Interact with the conversation by asking questions ("Who delivered the project at the city hall on march 1st?")

## Supported Chat Platforms

### Whatsapp

The input directory must contain the file "_chat.txt" exported from WhatsApp.

## Supported AI engines

### OpenAI (ChatGPT)

Set the OpenAI key as an environment variable named `OPENAI_API_KEY`

## Build

```shell
$ go build -o chat-ai-interaction
```

## Usage

```shell
./chat-ai-interaction -h
```
