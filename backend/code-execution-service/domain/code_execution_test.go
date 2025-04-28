package domain_test

import (
	"testing"
	"time"

	"github.com/augment-local-manus-clone/backend/code-execution-service/domain"
)

func TestNewCodeExecution(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		language domain.Language
		input    string
		wantErr  bool
	}{
		{
			name:     "Valid Python code",
			code:     "print('Hello, world!')",
			language: domain.LanguagePython,
			input:    "",
			wantErr:  false,
		},
		{
			name:     "Valid JavaScript code",
			code:     "console.log('Hello, world!');",
			language: domain.LanguageJavaScript,
			input:    "",
			wantErr:  false,
		},
		{
			name:     "Valid Go code",
			code:     "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world!\")\n}",
			language: domain.LanguageGo,
			input:    "",
			wantErr:  false,
		},
		{
			name:     "Valid Ruby code",
			code:     "puts 'Hello, world!'",
			language: domain.LanguageRuby,
			input:    "",
			wantErr:  false,
		},
		{
			name:     "Valid Java code",
			code:     "public class Main {\n\tpublic static void main(String[] args) {\n\t\tSystem.out.println(\"Hello, world!\");\n\t}\n}",
			language: domain.LanguageJava,
			input:    "",
			wantErr:  false,
		},
		{
			name:     "Empty code",
			code:     "",
			language: domain.LanguagePython,
			input:    "",
			wantErr:  true,
		},
		{
			name:     "Unsupported language",
			code:     "print('Hello, world!')",
			language: "unsupported",
			input:    "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execution, err := domain.NewCodeExecution(tt.code, tt.language, tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("NewCodeExecution() error = nil, wantErr %v", tt.wantErr)
				}
				return
			}
			
			if err != nil {
				t.Errorf("NewCodeExecution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if execution.Code != tt.code {
				t.Errorf("CodeExecution.Code = %v, want %v", execution.Code, tt.code)
			}
			
			if execution.Language != tt.language {
				t.Errorf("CodeExecution.Language = %v, want %v", execution.Language, tt.language)
			}
			
			if execution.Input != tt.input {
				t.Errorf("CodeExecution.Input = %v, want %v", execution.Input, tt.input)
			}
			
			// Check that CreatedAt is set
			now := time.Now()
			if execution.CreatedAt.After(now) || execution.CreatedAt.Before(now.Add(-time.Second)) {
				t.Errorf("CodeExecution.CreatedAt not set correctly: %v", execution.CreatedAt)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		execution *domain.CodeExecution
		wantErr   bool
	}{
		{
			name: "Valid execution",
			execution: &domain.CodeExecution{
				Code:     "print('Hello, world!')",
				Language: domain.LanguagePython,
			},
			wantErr: false,
		},
		{
			name: "Empty code",
			execution: &domain.CodeExecution{
				Code:     "",
				Language: domain.LanguagePython,
			},
			wantErr: true,
		},
		{
			name: "Unsupported language",
			execution: &domain.CodeExecution{
				Code:     "print('Hello, world!')",
				Language: "unsupported",
			},
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.execution.Validate()
			
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
