package utils

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

// GenerateSciFiStory generates a short sci-fi story using OpenAI API
func GenerateSciFiStory() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}

	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a creative sci-fi story writer. Write a very short engaging story in 400 characters or less. The story should have a clear beginning, twist and end.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate a sci-fi story",
				},
			},
			MaxTokens: 150,
		},
	)

	if err != nil {
		// Check for quota exceeded error
		if strings.Contains(err.Error(), "exceeded your current quota") {
			return "", fmt.Errorf("OpenAI API quota exceeded. Please try again later or contact the bot administrator")
		}
		return "", fmt.Errorf("failed to generate story: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no story generated")
	}

	return resp.Choices[0].Message.Content, nil
} 