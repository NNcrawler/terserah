package copywriter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ayush6624/go-chatgpt"
)

const systemInstruction = `Summarize the following review into a single review. Provide the summary as if you are the one who gives it. Return the review in HTML. Bold the strong point of the review. Have the summary in Bahasa Indonesia

Have 2 kinds of summary. One is food review summary. Second is place review summary

Return the summaries in array of json`

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
		Model: chatgpt.GPT4,
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
