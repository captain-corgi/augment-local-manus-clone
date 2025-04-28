package domain

import (
	"errors"
	"time"
)

// Language represents a programming language
type Language string

const (
	LanguagePython    Language = "python"
	LanguageJavaScript Language = "javascript"
	LanguageGo        Language = "go"
	LanguageRuby      Language = "ruby"
	LanguageJava      Language = "java"
)

// CodeExecution represents a code execution request
type CodeExecution struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Language  Language  `json:"language"`
	Input     string    `json:"input,omitempty"`
	Output    string    `json:"output,omitempty"`
	Error     string    `json:"error,omitempty"`
	ExitCode  int       `json:"exit_code"`
	CreatedAt time.Time `json:"created_at"`
	Duration  float64   `json:"duration,omitempty"`
}

// NewCodeExecution creates a new code execution request
func NewCodeExecution(code string, language Language, input string) (*CodeExecution, error) {
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}

	if !isValidLanguage(language) {
		return nil, errors.New("unsupported language")
	}

	return &CodeExecution{
		Code:      code,
		Language:  language,
		Input:     input,
		CreatedAt: time.Now(),
	}, nil
}

// Validate validates the code execution request
func (ce *CodeExecution) Validate() error {
	if ce.Code == "" {
		return errors.New("code cannot be empty")
	}

	if !isValidLanguage(ce.Language) {
		return errors.New("unsupported language")
	}

	return nil
}

// isValidLanguage checks if the language is supported
func isValidLanguage(language Language) bool {
	switch language {
	case LanguagePython, LanguageJavaScript, LanguageGo, LanguageRuby, LanguageJava:
		return true
	default:
		return false
	}
}

// DockerClient defines the interface for interacting with Docker
type DockerClient interface {
	// ExecuteCode executes code in a Docker container
	ExecuteCode(execution *CodeExecution) error
}
