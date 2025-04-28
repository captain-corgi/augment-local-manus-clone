package domain

import (
	"errors"
)

// AIRequest represents a request to the AI model
type AIRequest struct {
	Prompt     string                 `json:"prompt"`
	MaxTokens  int                    `json:"max_tokens,omitempty"`
	Temperature float64               `json:"temperature,omitempty"`
	Context     map[string]interface{} `json:"context,omitempty"`
}

// AIResponse represents a response from the AI model
type AIResponse struct {
	Text       string  `json:"text"`
	TokensUsed int     `json:"tokens_used,omitempty"`
	Model      string  `json:"model,omitempty"`
	Elapsed    float64 `json:"elapsed,omitempty"`
}

// Validate validates the AI request
func (r *AIRequest) Validate() error {
	if r.Prompt == "" {
		return errors.New("prompt cannot be empty")
	}
	
	if r.MaxTokens < 0 {
		return errors.New("max_tokens cannot be negative")
	}
	
	if r.Temperature < 0 || r.Temperature > 2 {
		return errors.New("temperature must be between 0 and 2")
	}
	
	return nil
}

// LLMClient defines the interface for interacting with the LLM
type LLMClient interface {
	// Process sends a request to the LLM and returns a response
	Process(request *AIRequest) (*AIResponse, error)
}
