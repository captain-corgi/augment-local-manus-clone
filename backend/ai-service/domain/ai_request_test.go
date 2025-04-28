package domain_test

import (
	"testing"

	"github.com/augment-local-manus-clone/backend/ai-service/domain"
)

func TestAIRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request domain.AIRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			request: domain.AIRequest{
				Prompt:      "Hello, world!",
				MaxTokens:   100,
				Temperature: 0.7,
			},
			wantErr: false,
		},
		{
			name: "Empty prompt",
			request: domain.AIRequest{
				Prompt:      "",
				MaxTokens:   100,
				Temperature: 0.7,
			},
			wantErr: true,
		},
		{
			name: "Negative max tokens",
			request: domain.AIRequest{
				Prompt:      "Hello, world!",
				MaxTokens:   -1,
				Temperature: 0.7,
			},
			wantErr: true,
		},
		{
			name: "Temperature too low",
			request: domain.AIRequest{
				Prompt:      "Hello, world!",
				MaxTokens:   100,
				Temperature: -0.1,
			},
			wantErr: true,
		},
		{
			name: "Temperature too high",
			request: domain.AIRequest{
				Prompt:      "Hello, world!",
				MaxTokens:   100,
				Temperature: 2.1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("Validate() error = nil, wantErr %v", tt.wantErr)
				}
				return
			}
			
			if err != nil {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
