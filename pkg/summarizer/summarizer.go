package summarizer

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const (
	promptFmt = "Você vai receber um histórico de conversa entre algumas pessoas em um dado dia. Seu trabalho é resumir o que foi conversado neste dia. Comece falando da data atual e o que foi falado. \n\nA conversa é a seguinte:\n%s"
)

type AIClient interface {
	CreateChatCompletion(ctx context.Context, request openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)
}

type Summarizer struct {
	client AIClient
	model  string
}

func New(client AIClient, model string) *Summarizer {
	return &Summarizer{
		client: client,
		model:  model,
	}
}

func (s *Summarizer) Summarize(ctx context.Context, input string) (string, error) {
	// Prepare the prompt
	prompt := fmt.Sprintf(promptFmt, input)

	// Create a request to the GPT API
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: s.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to call OpenAI API: %w", err)
	}

	// Extract and return the response content
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no response received from OpenAI API")
}
