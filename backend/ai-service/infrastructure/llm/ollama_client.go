package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/augment-local-manus-clone/backend/ai-service/domain"
)

// OllamaClient implements the LLMClient interface using Ollama
type OllamaClient struct {
	baseURL string
	model   string
	client  *http.Client
}

// ollamaRequest represents a request to the Ollama API
type ollamaRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Stream      bool    `json:"stream"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

// ollamaResponse represents a response from the Ollama API
type ollamaResponse struct {
	Model     string  `json:"model"`
	Response  string  `json:"response"`
	Done      bool    `json:"done"`
	TokensUsed int     `json:"tokens_used,omitempty"`
	Elapsed   float64 `json:"elapsed,omitempty"`
}

// NewOllamaClient creates a new OllamaClient
func NewOllamaClient(baseURL, model string) (*OllamaClient, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL cannot be empty")
	}

	if model == "" {
		return nil, fmt.Errorf("model cannot be empty")
	}

	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}, nil
}

// Process sends a request to the Ollama API and returns a response
func (c *OllamaClient) Process(request *domain.AIRequest) (*domain.AIResponse, error) {
	// Create Ollama request
	ollamaReq := ollamaRequest{
		Model:       c.model,
		Prompt:      request.Prompt,
		Stream:      false,
		MaxTokens:   request.MaxTokens,
		Temperature: request.Temperature,
	}

	// Convert request to JSON
	reqBody, err := json.Marshal(ollamaReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/generate", c.baseURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode response
	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Create AI response
	aiResp := &domain.AIResponse{
		Text:       ollamaResp.Response,
		TokensUsed: ollamaResp.TokensUsed,
		Model:      ollamaResp.Model,
		Elapsed:    ollamaResp.Elapsed,
	}

	return aiResp, nil
}
