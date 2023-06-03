package llm_openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var MessageQueueOpenAI map[string][]openai.ChatCompletionMessage

func CompletionWithoutSession(ctx context.Context, client *openai.Client, prompt string, temperature float32) (string, error) {
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func CompletionWithoutSessionWithStream(ctx context.Context, client *openai.Client, prompt string) (*openai.ChatCompletionStream, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return nil, err
	}
	return stream, nil
}

func CompletionWithSession(ctx context.Context, client *openai.Client, conversationID string, prompt string) (string, error) {
	messages := AddMessage(conversationID, prompt)
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	AddMessage(conversationID, resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}

// CompletionWithSessionWithStream should call AddMessage after using stream
func CompletionWithSessionWithStream(ctx context.Context, client *openai.Client, conversationID string, prompt string) (*openai.ChatCompletionStream, error) {
	messages := AddMessage(conversationID, prompt)
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages:  messages,
		Stream:    true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return nil, err
	}
	return stream, nil
}

func AddMessage(conversationID string, prompt string) []openai.ChatCompletionMessage {
	if MessageQueueOpenAI == nil {
		MessageQueueOpenAI = make(map[string][]openai.ChatCompletionMessage)
	}
	if _, ok := MessageQueueOpenAI[conversationID]; !ok {
		MessageQueueOpenAI[conversationID] = []openai.ChatCompletionMessage{}
	}
	MessageQueueOpenAI[conversationID] = append(MessageQueueOpenAI[conversationID], openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	return MessageQueueOpenAI[conversationID]
}
