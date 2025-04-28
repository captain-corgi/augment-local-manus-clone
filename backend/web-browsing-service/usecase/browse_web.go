package usecase

import (
	"github.com/augment-local-manus-clone/backend/web-browsing-service/domain"
)

// BrowseWebUseCase handles web browsing operations
type BrowseWebUseCase struct {
	browserClient domain.BrowserClient
}

// NewBrowseWebUseCase creates a new instance of BrowseWebUseCase
func NewBrowseWebUseCase(browserClient domain.BrowserClient) *BrowseWebUseCase {
	return &BrowseWebUseCase{
		browserClient: browserClient,
	}
}

// Execute navigates to a URL and returns the page content
func (uc *BrowseWebUseCase) Execute(request *domain.WebBrowsingRequest) (*domain.WebBrowsingResult, error) {
	// Validate the request
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Set default values if not provided
	if request.Timeout == 0 {
		request.Timeout = 30 // Default timeout: 30 seconds
	}

	// Browse the web page
	return uc.browserClient.Browse(request)
}
