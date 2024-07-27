package copywriter

import (
	"bytes"
	"context"
	"text/template"

	"fmt"

	"github.com/ayush6624/go-chatgpt"
)

func CallOpenAI(apiKey, prompt string) (string, error) {
	client, err := chatgpt.NewClient(apiKey)
	if err != nil {
		fmt.Print(err)
	}
	res, err := client.SimpleSend(context.Background(), prompt)
	if err != nil {
		return "", fmt.Errorf("fail to call openAI: ", err)
	}

	return res.Choices[0].Message.Content, nil
}

type DishToRecommend struct {
	Name string
}

type Writer struct {
	ApiKey string
}

func (w Writer) AsBestFriend(dish DishToRecommend) (string, error) {
	const promptTmpl = `Buat copywriting untuk jualan sebagai berikut: Anggap kamu adalah seorang sahabat. Sebagai sahabat kamu ingin meyakinkan sahabatmu untuk mencoba {{.Name}}`

	prompt, err := stringTmplRenderer(promptTmpl, dish)
	if err != nil {
		return "", fmt.Errorf("fail to render prompt: %w", err)
	}
	return CallOpenAI(w.ApiKey, prompt)
}

func stringTmplRenderer(promptTmpl string, data interface{}) (string, error) {
	t, err := template.New("").
		Parse(promptTmpl)
	if err != nil {
		return "", err
	}

	out := bytes.Buffer{}
	err = t.Execute(&out, data)
	if err != nil {
		return "", fmt.Errorf("fail to render prompt: %w", err)
	}
	return out.String(), nil
}
