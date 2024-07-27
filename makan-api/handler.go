package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nncrawler/makan-api/copywriter"
)

// HelloWorld is an HTTP Cloud Function.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

type Request struct {
	Mode string          `json:"mode"`
	Data json.RawMessage `json:"data"`
}

type ReviewSummarizerRequest struct {
	Reviews []string `json:"reviews"`
}
type LocalGuideRecommendationRequest struct {
	Name string `json:"name"`
}

type Response struct {
	Content string `json:"message"`
}

func CopyWriteFood(w http.ResponseWriter, r *http.Request) {

	shouldReturn := handleCors(w, r)
	if shouldReturn {
		return
	}

	openAIKey := os.Getenv("OPEN_AI_KEY")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error while parsing request: %v", err)
		return
	}

	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Fprintf(w, "error while parsing request: %v", err)
		return
	}

	var taskMap = initiateCopyWriteTasks(openAIKey)

	taskFn, ok := taskMap[req.Mode]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid mode")
		return
	}

	response, err := taskFn(req.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	fmt.Fprintf(w, response)
}

func initiateCopyWriteTasks(openAIKey string) map[string]func([]byte) (string, error) {
	return map[string]func([]byte) (string, error){
		"reviewSummarizer": func(data []byte) (string, error) {
			var taskParam ReviewSummarizerRequest
			json.Unmarshal(data, &taskParam)

			reviewSummarizer := copywriter.ReviewSummarizer{
				ApiKey: openAIKey,
			}
			return reviewSummarizer.AsReviewer(taskParam.Reviews)
		},
		"localGuideRecommendation": func(data []byte) (string, error) {
			var taskParam LocalGuideRecommendationRequest
			json.Unmarshal(data, &taskParam)

			cpw := copywriter.Writer{
				ApiKey: openAIKey,
			}

			return cpw.AsLocalGuide(copywriter.DishToRecommend{
				Name: taskParam.Name,
			})
		},
	}
}

func handleCors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}
