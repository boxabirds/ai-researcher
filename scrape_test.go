package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapeBBCNewsArticle(t *testing.T) {
	url := "https://www.bbc.com/news/science-environment-24021772"
	result, err := scrape(url)
	assert.Nil(t, err, "Error should be nil")

	assert.True(t, result.Success, "Expected Success to be true")

	assert.NotEmpty(t, result.Data.Content, "Expected non-empty content")
	assert.NotEmpty(t, result.Data.Markdown, "Expected non-empty markdown")

	assert.NotEmpty(t, result.Data.Metadata.Title, "Expected non-empty title")
	assert.NotEmpty(t, result.Data.Metadata.Description, "Expected non-empty description")
	assert.NotEmpty(t, result.Data.Metadata.SourceURL, "Expected non-empty sourceURL")
}
