package main

import (
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func newLLMClient(baseUrl string) llms.Model {
	model, err := openai.New(openai.WithModel("llama3:instruct"), openai.WithBaseURL(baseUrl))
	if err != nil {
		log.Printf("Error initializing model %s: %v", "llama3", err)
		return nil
	}
	return model
}

func main() {
	// use the llm.go newLLMClientWithCustomBaseUrlModelName to create a new client
	// with a custom base url and model name
	client := newLLMClient("http://gruntus:11434/v1")

	result := search("climate change")
	fmt.Printf("result: %v\n", result)
}
