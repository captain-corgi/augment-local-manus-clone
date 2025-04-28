package domain_test

import (
	"testing"

	"github.com/augment-local-manus-clone/backend/web-browsing-service/domain"
)

func TestWebBrowsingRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request domain.WebBrowsingRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			request: domain.WebBrowsingRequest{
				URL:     "https://example.com",
				Timeout: 30,
			},
			wantErr: false,
		},
		{
			name: "Empty URL",
			request: domain.WebBrowsingRequest{
				URL:     "",
				Timeout: 30,
			},
			wantErr: true,
		},
		{
			name: "Invalid URL",
			request: domain.WebBrowsingRequest{
				URL:     "not-a-url",
				Timeout: 30,
			},
			wantErr: true,
		},
		{
			name: "Negative timeout",
			request: domain.WebBrowsingRequest{
				URL:     "https://example.com",
				Timeout: -1,
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

func TestWebSearchRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request domain.WebSearchRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			request: domain.WebSearchRequest{
				Query:        "test query",
				SearchEngine: "google",
				NumResults:   10,
			},
			wantErr: false,
		},
		{
			name: "Empty query",
			request: domain.WebSearchRequest{
				Query:        "",
				SearchEngine: "google",
				NumResults:   10,
			},
			wantErr: true,
		},
		{
			name: "Negative num results",
			request: domain.WebSearchRequest{
				Query:        "test query",
				SearchEngine: "google",
				NumResults:   -1,
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

func TestWebInteractionRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request domain.WebInteractionRequest
		wantErr bool
	}{
		{
			name: "Valid click request",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "#button",
				Action:   "click",
			},
			wantErr: false,
		},
		{
			name: "Valid type request",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "#input",
				Action:   "type",
				Value:    "test value",
			},
			wantErr: false,
		},
		{
			name: "Empty URL",
			request: domain.WebInteractionRequest{
				URL:      "",
				Selector: "#button",
				Action:   "click",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL",
			request: domain.WebInteractionRequest{
				URL:      "not-a-url",
				Selector: "#button",
				Action:   "click",
			},
			wantErr: true,
		},
		{
			name: "Empty selector",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "",
				Action:   "click",
			},
			wantErr: true,
		},
		{
			name: "Empty action",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "#button",
				Action:   "",
			},
			wantErr: true,
		},
		{
			name: "Invalid action",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "#button",
				Action:   "invalid",
			},
			wantErr: true,
		},
		{
			name: "Type action without value",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "#input",
				Action:   "type",
				Value:    "",
			},
			wantErr: true,
		},
		{
			name: "Select action without value",
			request: domain.WebInteractionRequest{
				URL:      "https://example.com",
				Selector: "#select",
				Action:   "select",
				Value:    "",
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
