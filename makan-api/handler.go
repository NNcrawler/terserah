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

	cpw := copywriter.Writer{
		ApiKey: openAIKey,
	}

	response, err := cpw.AsLocalGuide(copywriter.DishToRecommend{
		Name: req.Name,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintf(w, response)
}

func handleCors(w http.ResponseWriter, r *http.Request) bool {
	// Set CORS headers
	// Allow all origins
	// Allow specific methods
	// Allow specific headers
	// Handle preflight requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}
