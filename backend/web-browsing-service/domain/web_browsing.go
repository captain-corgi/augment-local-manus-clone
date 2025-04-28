package domain

import (
	"errors"
	"net/url"
	"time"
)

// WebBrowsingRequest represents a request to browse a web page
type WebBrowsingRequest struct {
	URL           string            `json:"url"`
	Timeout       int               `json:"timeout,omitempty"`
	WaitForSelector string          `json:"wait_for_selector,omitempty"`
	ExtractSelectors []string       `json:"extract_selectors,omitempty"`
	Headers       map[string]string `json:"headers,omitempty"`
	TakeScreenshot bool             `json:"take_screenshot,omitempty"`
}

// WebSearchRequest represents a request to search the web
type WebSearchRequest struct {
	Query         string `json:"query"`
	SearchEngine  string `json:"search_engine,omitempty"`
	NumResults    int    `json:"num_results,omitempty"`
	TakeScreenshot bool  `json:"take_screenshot,omitempty"`
}

// WebInteractionRequest represents a request to interact with a web page
type WebInteractionRequest struct {
	URL           string `json:"url"`
	Selector      string `json:"selector"`
	Action        string `json:"action"`
	Value         string `json:"value,omitempty"`
	WaitForSelector string `json:"wait_for_selector,omitempty"`
	TakeScreenshot bool   `json:"take_screenshot,omitempty"`
}

// WebBrowsingResult represents the result of a web browsing operation
type WebBrowsingResult struct {
	URL           string            `json:"url"`
	Title         string            `json:"title"`
	Content       string            `json:"content"`
	ExtractedData map[string]string `json:"extracted_data,omitempty"`
	Screenshot    string            `json:"screenshot,omitempty"`
	StatusCode    int               `json:"status_code"`
	Error         string            `json:"error,omitempty"`
	Duration      float64           `json:"duration"`
	Timestamp     time.Time         `json:"timestamp"`
}

// Validate validates the web browsing request
func (r *WebBrowsingRequest) Validate() error {
	if r.URL == "" {
		return errors.New("URL cannot be empty")
	}

	// Check if URL is valid
	_, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return errors.New("invalid URL format")
	}

	if r.Timeout < 0 {
		return errors.New("timeout cannot be negative")
	}

	return nil
}

// Validate validates the web search request
func (r *WebSearchRequest) Validate() error {
	if r.Query == "" {
		return errors.New("query cannot be empty")
	}

	if r.NumResults < 0 {
		return errors.New("num_results cannot be negative")
	}

	return nil
}

// Validate validates the web interaction request
func (r *WebInteractionRequest) Validate() error {
	if r.URL == "" {
		return errors.New("URL cannot be empty")
	}

	// Check if URL is valid
	_, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return errors.New("invalid URL format")
	}

	if r.Selector == "" {
		return errors.New("selector cannot be empty")
	}

	if r.Action == "" {
		return errors.New("action cannot be empty")
	}

	// Validate action
	validActions := map[string]bool{
		"click":     true,
		"type":      true,
		"select":    true,
		"focus":     true,
		"hover":     true,
		"scroll":    true,
		"screenshot": true,
	}

	if !validActions[r.Action] {
		return errors.New("invalid action")
	}

	// If action is "type" or "select", value is required
	if (r.Action == "type" || r.Action == "select") && r.Value == "" {
		return errors.New("value is required for type and select actions")
	}

	return nil
}

// BrowserClient defines the interface for interacting with a browser
type BrowserClient interface {
	// Browse navigates to a URL and returns the page content
	Browse(request *WebBrowsingRequest) (*WebBrowsingResult, error)

	// Search performs a search query and returns the results
	Search(request *WebSearchRequest) (*WebBrowsingResult, error)

	// Interact interacts with elements on a web page
	Interact(request *WebInteractionRequest) (*WebBrowsingResult, error)

	// Close closes the browser
	Close() error
}
