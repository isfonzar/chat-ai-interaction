package parser

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"isfonzar/chat-ai-interaction/pkg/date"
	"isfonzar/chat-ai-interaction/pkg/summarizer"

	"github.com/sashabaranov/go-openai"
)

const (
	chatGPTModel = "gpt-4o"
)

func Parse(file *os.File, dateInput, openAIAPIKey string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("invalid file descriptor: file is nil")
	}

	regex, err := date.Filter(dateInput)
	if err != nil {
		return "", fmt.Errorf("invalid date input: %w", err)
	}

	// Use bufio to efficiently read the file
	reader := bufio.NewReader(file)
	var content string

	for {
		// Read line by line or until EOF
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// Append the last part of the file
				content += line
				break
			}
			return "", fmt.Errorf("error reading file: %v", err)
		}

		matches := regex.FindStringSubmatch(line)
		if len(matches) > 0 {
			content += line
		}
	}

	// now that we have a full day, send it to AI
	client := openai.NewClient(openAIAPIKey)
	s := summarizer.New(client, chatGPTModel)

	resp, err := s.Summarize(context.Background(), content)
	if err != nil {
		return "", fmt.Errorf("error summarizing: %w", err)
	}

	return resp, nil
}
