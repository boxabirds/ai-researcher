package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Define the response struct to match the JSON response structure
type ScrapeResult struct {
	Success bool `json:"success"`
	Data    struct {
		Content  string `json:"content"`
		Markdown string `json:"markdown"`
		Metadata struct {
			Title       string  `json:"title"`
			Description string  `json:"description"`
			Language    *string `json:"language"`
			SourceURL   string  `json:"sourceURL"`
		} `json:"metadata"`
	} `json:"data"`
}

// Function to make the API call
func scrape(targetURL string) (*ScrapeResult, error) {
	// Prepare the request body
	requestBody, err := json.Marshal(map[string]string{
		"url": targetURL,
	})
	if err != nil {
		return nil, err
	}

	// Create the request
	req, err := http.NewRequest("POST", "https://api.firecrawl.dev/v0/scrape", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer fc-ddc52f04a5664e4d801c1b617b0600c7")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse ScrapeResult
	err = json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
