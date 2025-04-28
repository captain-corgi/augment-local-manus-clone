package usecase

import (
	"github.com/augment-local-manus-clone/backend/web-browsing-service/domain"
)

// InteractWithWebUseCase handles web interaction operations
type InteractWithWebUseCase struct {
	browserClient domain.BrowserClient
}

// NewInteractWithWebUseCase creates a new instance of InteractWithWebUseCase
func NewInteractWithWebUseCase(browserClient domain.BrowserClient) *InteractWithWebUseCase {
	return &InteractWithWebUseCase{
		browserClient: browserClient,
	}
}

// Execute interacts with elements on a web page
func (uc *InteractWithWebUseCase) Execute(request *domain.WebInteractionRequest) (*domain.WebBrowsingResult, error) {
	// Validate the request
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Interact with the web page
	return uc.browserClient.Interact(request)
}
