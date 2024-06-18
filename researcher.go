package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/tmc/langchaingo/llms"
)

type Topics []string

func llmRequest(systemPrompt, prompt string, llm llms.Model) (string, error) {
	const MaxTokens = 4096
	ctx := context.Background()
	options := []llms.CallOption{
		llms.WithMaxTokens(MaxTokens),
		llms.WithTemperature(0.0),
	}
	// if seed != NoSeed {
	// 	options = append(options, llms.WithSeed(seed))
	// }

	start := time.Now()
	response, err := llms.GenerateFromSinglePrompt(ctx, llm, systemPrompt+"\n"+prompt, options...)
	elapsed := time.Since(start)
	fmt.Printf("- Query generation execution time: %s\n", elapsed)

	if err != nil {
		return "", err
	}

	return response, nil

}

func generateSubTopics(topic string, numTopics int, model llms.Model) (Topics, error) {
	const JsonExample = `["Global Warming", "Greenhouse Gases", "Climate Crisis"]`
	const TopicGeneratorSystemPrompt = `
	You are a subtopic generation API that returns a list of exactly %d strings — no less, no more —
	that are subtopics of the message so the result can be read by a JSON parser,
	with no additional narrative, escape quotes or backticks before or after, 
	for example: %s. JSON output only.`

	systemPrompt := fmt.Sprintf(TopicGeneratorSystemPrompt, numTopics, JsonExample)
	rawText, err := llmRequest(systemPrompt, topic, model)
	if err != nil {
		log.Printf("Error generating subtopics: %v", err)
		return nil, err
	}

	var topics Topics
	err = json.Unmarshal([]byte(rawText), &topics)
	if err != nil {
		return nil, fmt.Errorf("JSON decoding error: %v", err)
	}

	fmt.Printf("Extracted topics: %s\n", topics)

	return topics, nil
}
