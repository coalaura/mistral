# Go Mistral AI API Client

A lightweight Go client library for interacting with the Mistral AI API. This library provides easy access to Mistral's powerful language models and AI capabilities.

## Features

- Support for chat completions
- Function calling capabilities
- Vision model support
- Classification/moderation endpoints
- Agent completions

## Installation

```bash
go get github.com/coalaura/mistral
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/coalaura/mistral"
)

func main() {
    // Initialize client with your API key
    client := mistral.NewMistralClient("your-api-key")

    // Create a simple chat request
    request := mistral.ChatCompletionRequest{
        Model: mistral.ModelMistralSmall,
        Messages: []mistral.Message{
            {
                Role:    mistral.RoleUser,
                Content: "Hello, how are you?",
            },
        },
    }

    // Send the chat request
    response, err := client.Chat(request)
    if err != nil {
        panic(err)
    }

    fmt.Println(response.Choices[0].Message.Content)
}
```

## Available Models

See the [Mistral Docs](https://docs.mistral.ai/getting-started/models/models_overview/) for a list of available models.

## Features

### Chat Completions
Send chat messages and receive AI-generated responses.

### Classification/Moderation
Use Mistral's moderation capabilities to classify content.

### Function Calling
Implement function calling with custom tools and parameters.

### Agent Completions
Interact with Mistral's agent-based completions.

## Contributions

Contributions are welcome, however this is a side project and I may not be able to respond to issues or pull requests in a timely manner.

## License

This project is licensed under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.