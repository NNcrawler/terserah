package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nncrawler/makan-api/copywriter"
)

// HelloWorld is an HTTP Cloud Function.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func CopyWriteFood(w http.ResponseWriter, r *http.Request) {
	openAIKey := os.Getenv("OPEN_AI_KEY")
	text := "hello open ai"
	response, err := copywriter.CallOpenAI(openAIKey, text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintf(w, response)
}
