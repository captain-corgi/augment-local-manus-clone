package usecase

import (
	"github.com/augment-local-manus-clone/backend/ai-service/domain"
)

// ProcessAIRequestUseCase handles processing AI requests
type ProcessAIRequestUseCase struct {
	llmClient domain.LLMClient
}

// NewProcessAIRequestUseCase creates a new instance of ProcessAIRequestUseCase
func NewProcessAIRequestUseCase(llmClient domain.LLMClient) *ProcessAIRequestUseCase {
	return &ProcessAIRequestUseCase{
		llmClient: llmClient,
	}
}

// Execute processes an AI request
func (uc *ProcessAIRequestUseCase) Execute(request *domain.AIRequest) (*domain.AIResponse, error) {
	// Validate the request
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Set default values if not provided
	if request.MaxTokens == 0 {
		request.MaxTokens = 2048
	}

	if request.Temperature == 0 {
		request.Temperature = 0.7
	}

	// Process the request using the LLM client
	return uc.llmClient.Process(request)
}
