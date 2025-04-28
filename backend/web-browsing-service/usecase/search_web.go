package usecase

import (
	"github.com/augment-local-manus-clone/backend/web-browsing-service/domain"
)

// SearchWebUseCase handles web search operations
type SearchWebUseCase struct {
	browserClient domain.BrowserClient
}

// NewSearchWebUseCase creates a new instance of SearchWebUseCase
func NewSearchWebUseCase(browserClient domain.BrowserClient) *SearchWebUseCase {
	return &SearchWebUseCase{
		browserClient: browserClient,
	}
}

// Execute performs a search query and returns the results
func (uc *SearchWebUseCase) Execute(request *domain.WebSearchRequest) (*domain.WebBrowsingResult, error) {
	// Validate the request
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Set default values if not provided
	if request.SearchEngine == "" {
		request.SearchEngine = "google" // Default search engine
	}

	if request.NumResults == 0 {
		request.NumResults = 10 // Default number of results
	}

	// Search the web
	return uc.browserClient.Search(request)
}
