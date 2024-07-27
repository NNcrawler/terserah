package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ayush6624/go-chatgpt"
)

type ReviewSummarizer struct {
	ApiKey string
}

func (w ReviewSummarizer) AsReviewer(reviews []string) (ReviewSummary, error) {
	const prompt = `Summarize the following review into a single review. Provide the summary as if you are the one who gives it. Return the review in HTML. Bold the strong point of the review. Have the summary in Bahasa Indonesia

Have 2 kinds of summary. One is a food review summary. The second is a place review summary. Have one review for each category. Each summary has a maximum of 280 characters.

Return the summaries in a JSON. Where category becomes the key and the review summary becomes the value.`

	bytesReviews, _ := json.Marshal(reviews)

	client, err := chatgpt.NewClient(w.ApiKey)
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
		return ReviewSummary{}, fmt.Errorf("fail to call openAI: %v", err)
	}

	if len(res.Choices) == 0 {
		return ReviewSummary{}, fmt.Errorf("no response from openAI")
	}

	var reviewSummary ReviewSummary
	json.Unmarshal([]byte(res.Choices[0].Message.Content), &reviewSummary)

	return reviewSummary, nil
}

type ReviewSummary struct {
	Food  string
	Place string
}
