package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)

var (
	local_url = "http://localhost:11434"
	model     = "gemma3:latest"
)

func main() {
	u, err := url.Parse(local_url)
	if err != nil {
		log.Fatal(err)
	}
	ollama := api.NewClient(u, http.DefaultClient)

	messages := []api.Message{
		{Role: "user", Content: "Hello, Ollama! Please tell me a joke about programmers."},
	}
	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	err = ollama.Chat(context.Background(), &api.ChatRequest{
		Model:    model,
		Messages: messages,
	}, respFunc)
	if err != nil {
		log.Fatal(err)
	}

}
