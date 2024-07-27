package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"strings"

	"github.com/ayush6624/go-chatgpt"
)

type DishExtractor struct {
	apiKey string
}

func NewDishExtractor(apiKey string) DishExtractor {
	return DishExtractor{
		apiKey: apiKey,
	}
}

func (c DishExtractor) GetPossibleFoodsFromPlace(ctx context.Context, reviews []string) ([]string, error) {
	const prompt = "You are a extraction system, in which given a list of restaurant reviews, you provide the user with ONLY the list of foods served there. Output the list as a comma separated value & lowercase."

	bytesReviews, _ := json.Marshal(strings.Join(reviews, "; "))

	client, err := chatgpt.NewClient(c.apiKey)
	if err != nil {
		fmt.Print(err)
	}
	res, err := client.Send(context.Background(), &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: prompt,
			},
			{
				Role:    chatgpt.ChatGPTModelRoleUser,
				Content: string(bytesReviews),
			},
		},
	})
	if err != nil {
		return []string{}, fmt.Errorf("fail to call openAI: %v", err)
	}

	if len(res.Choices) == 0 {
		return []string{}, fmt.Errorf("no response from openAI")
	}

	foods := strings.Split(res.Choices[0].Message.Content, ", ")

	return foods, nil
}
