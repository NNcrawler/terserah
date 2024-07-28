package copywriter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ayush6624/go-chatgpt"
)

const systemInstruction = `Summarize the following review into a single review. Provide the summary as if you are the one who gives it. Return the review in HTML. Bold the strong point of the review. Have the summary in Bahasa Indonesia

Have 2 kinds of summary. One is a food review summary. The second is a place review summary. Have one review for each category. Each summary has a maximum of 280 characters.

Return the summaries in an array of JSON. The JSON has 2 fields called "category" and "review"`

type ReviewSummarizer struct {
	ApiKey string
}

func (w ReviewSummarizer) AsReviewer(reviews []string) (string, error) {
	bytesReviews, _ := json.Marshal(reviews)

	client, err := chatgpt.NewClient(w.ApiKey)
	if err != nil {
		fmt.Print(err)
	}
	res, err := client.Send(context.Background(), &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT35Turbo,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: systemInstruction,
			},
			{
				Role:    chatgpt.ChatGPTModelRoleUser,
				Content: string(bytesReviews),
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("fail to call openAI: ", err)
	}

	return res.Choices[0].Message.Content, nil
}
