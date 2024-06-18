package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type SearchResult struct {
	SearchMetadata    SearchMetadata    `json:"search_metadata"`
	SearchParameters  SearchParameters  `json:"search_parameters"`
	SearchInformation SearchInformation `json:"search_information"`
	Ads               []Ad              `json:"ads"`
	KnowledgeGraph    KnowledgeGraph    `json:"knowledge_graph"`
	OrganicResults    []OrganicResult   `json:"organic_results"`
}

type SearchMetadata struct {
	ID               string  `json:"id"`
	Status           string  `json:"status"`
	CreatedAt        string  `json:"created_at"`
	RequestTimeTaken float64 `json:"request_time_taken"`
	ParsingTimeTaken float64 `json:"parsing_time_taken"`
	TotalTimeTaken   float64 `json:"total_time_taken"`
	RequestURL       string  `json:"request_url"`
	HTMLURL          string  `json:"html_url"`
	JSONURL          string  `json:"json_url"`
}

type SearchParameters struct {
	Engine       string `json:"engine"`
	Query        string `json:"q"`
	Device       string `json:"device"`
	GoogleDomain string `json:"google_domain"`
	HL           string `json:"hl"`
	GL           string `json:"gl"`
}

type SearchInformation struct {
	QueryDisplayed     string  `json:"query_displayed"`
	TotalResults       int64   `json:"total_results"`
	TimeTakenDisplayed float64 `json:"time_taken_displayed"`
}

type Ad struct {
	Position                int      `json:"position"`
	BlockPosition           string   `json:"block_position"`
	Title                   string   `json:"title"`
	Link                    string   `json:"link"`
	Source                  string   `json:"source"`
	Domain                  string   `json:"domain"`
	DisplayedLink           string   `json:"displayed_link"`
	Snippet                 string   `json:"snippet"`
	SnippetHighlightedWords []string `json:"snippet_highlighted_words"`
}

type KnowledgeGraph struct {
	KGMID                    string                `json:"kgmid"`
	KnowledgeGraphType       string                `json:"knowledge_graph_type"`
	Title                    string                `json:"title"`
	Type                     string                `json:"type"`
	Description              string                `json:"description"`
	Source                   Source                `json:"source"`
	Developer                string                `json:"developer"`
	DeveloperLinks           []Link                `json:"developer_links"`
	InitialReleaseDate       string                `json:"initial_release_date"`
	ProgrammingLanguage      string                `json:"programming_language"`
	ProgrammingLanguageLinks []Link                `json:"programming_language_links"`
	Engine                   string                `json:"engine"`
	EngineLinks              []Link                `json:"engine_links"`
	License                  string                `json:"license"`
	Platform                 string                `json:"platform"`
	PlatformLinks            []Link                `json:"platform_links"`
	StableRelease            string                `json:"stable_release"`
	PeopleAlsoSearchFor      []PeopleAlsoSearchFor `json:"people_also_search_for"`
	Image                    string                `json:"image"`
}

type Source struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Link struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type PeopleAlsoSearchFor struct {
	Name  string `json:"name"`
	Link  string `json:"link"`
	Image string `json:"image"`
}

type OrganicResult struct {
	Position                int       `json:"position"`
	Title                   string    `json:"title"`
	Link                    string    `json:"link"`
	Source                  string    `json:"source"`
	Domain                  string    `json:"domain"`
	DisplayedLink           string    `json:"displayed_link"`
	Snippet                 string    `json:"snippet"`
	SnippetHighlightedWords []string  `json:"snippet_highlighted_words"`
	Date                    string    `json:"date"`
	Sitelinks               Sitelinks `json:"sitelinks"`
	Favicon                 string    `json:"favicon"`
	Thumbnail               string    `json:"thumbnail"`
}

type Sitelinks struct {
	Inline []InlineLink `json:"inline"`
	List   []ListLink   `json:"list"`
}

type InlineLink struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type ListLink struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Date  string `json:"date"`
}

func search(query string) SearchResult {
	const BaseURL = "https://www.searchapi.io/api/v1/search"
	apiKey := os.Getenv("SEARCHAPI_API_KEY")
	queryParams := url.Values{
		"engine":  {"google"},
		"q":       {query},
		"api_key": []string{apiKey},
	}
	fullURL := fmt.Sprintf("%s?%s", BaseURL, queryParams.Encode())
	//fmt.Printf("search query: '%s'\n", fullURL)

	req, _ := http.NewRequest("GET", fullURL, nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status code: %d", res.StatusCode))
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Println(string(body))

	// Declare a variable of type SearchResult to hold the decoded JSON
	var searchResult SearchResult

	// Unmarshal the JSON response into the searchResult variable
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Now you can access fields within searchResult, for example:
	//fmt.Printf("First link: %s\n", searchResult.OrganicResults[0].Link)
	//fmt.Printf("Number of Organic Results: %d\n", len(searchResult.OrganicResults))
	return searchResult
}
