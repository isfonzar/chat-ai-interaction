package summarizer

import (
	"context"
	"errors"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestSummarizer(t *testing.T) {
	tests := []struct {
		desc         string
		expectedResp string
		expectedErr  error
	}{
		{
			"failed to call API",
			"",
			errors.New("openai error"),
		},
		{
			"successful call",
			"bar",
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			clientMock := &AIClientMock{
				resp: test.expectedResp,
				err:  test.expectedErr,
			}

			summarizer := New(clientMock, "model")

			got, err := summarizer.Summarize(context.Background(), "foo")
			if got != test.expectedResp {
				t.Errorf("got: %v, want: %v", got, test.expectedResp)
			}
			if err != nil && test.expectedErr == nil {
				t.Errorf("did not expect err, got: %v", test.expectedErr)
			}
			if err == nil && test.expectedErr != nil {
				t.Errorf("expected err: %v, got none", test.expectedErr)
			}
		})
	}
}

type AIClientMock struct {
	resp string
	err  error
}

func (m *AIClientMock) CreateChatCompletion(_ context.Context, _ openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	return openai.ChatCompletionResponse{
		Choices: []openai.ChatCompletionChoice{
			{
				Message: openai.ChatCompletionMessage{Content: m.resp},
			},
		},
	}, m.err
}
