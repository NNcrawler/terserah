package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ahmadnaufal/recommender-worker/model"
)

type OpenAIClient struct {
	client *http.Client
	host   string
	apiKey string
}

func New(host, apiKey string) OpenAIClient {
	return OpenAIClient{
		client: http.DefaultClient,
		host:   host,
		apiKey: apiKey,
	}
}

func (c OpenAIClient) GetPossibleFoodsFromPlace(ctx context.Context, place model.Place) ([]string, error) {
	requestURL := fmt.Sprintf("%s/v1/chat/completions", c.host)
	requestCmp := CompletionRequest{
		Model: modelGPT35Turbo,
		Messages: []Message{
			{
				Role:    roleSystem,
				Content: systemPromptReview,
			},
			{
				Role:    roleUser,
				Content: strings.Join(place.Reviews, "; "),
			},
		},
	}

	requestBody, _ := json.Marshal(requestCmp)
	req, _ := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	// provide headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var decodedResp Response
	err = json.NewDecoder(resp.Body).Decode(&decodedResp)
	if err != nil {
		return nil, err
	}

	foods := strings.Split(decodedResp.Choices[0].Message.Content, ", ")

	return foods, nil
}
