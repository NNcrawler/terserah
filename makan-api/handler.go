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
	Name string `json:"name"`
}

type Response struct {
	Content string `json:"message"`
}

func CopyWriteFood(w http.ResponseWriter, r *http.Request) {
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

	cpw := copywriter.Writer{
		ApiKey: openAIKey,
	}

	response, err := cpw.AsBestFriend(copywriter.DishToRecommend{
		Name: req.Name,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintf(w, response)
}
