package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchQueryClimateChange(t *testing.T) {

	// Perform the search
	result := search("climate change")

	// Assert that there is at least one organic result
	require.NotEmpty(t, result.OrganicResults, "Expected at least one organic result")

	// Check the first organic result for non-null values
	firstResult := result.OrganicResults[0]
	assert.NotEmpty(t, firstResult.Link, "Expected non-null link in the first organic result")
	assert.NotEmpty(t, firstResult.Snippet, "Expected non-null snippet in the first organic result")
	assert.NotEmpty(t, firstResult.Title, "Expected non-null title in the first organic result")
}
