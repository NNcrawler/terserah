package copywriter

import (
	"context"

	"fmt"

	"github.com/ayush6624/go-chatgpt"
)

func CallOpenAI(apiKey, prompt string) (string, error) {
	client, err := chatgpt.NewClient(apiKey)
	if err != nil {
		fmt.Print(err)
	}
	res, err := client.SimpleSend(context.Background(), "Hello, how are you?")
	if err != nil {
		return "", fmt.Errorf("fail to call openAI: ", err)
	}

	return res.Choices[0].Message.Content, nil
}
