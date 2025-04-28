package domain_test

import (
	"errors"
	"testing"
	"time"

	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

func TestNewFileOperation(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		operation string
		content   string
	}{
		{
			name:      "Read operation",
			path:      "test.txt",
			operation: "read",
			content:   "",
		},
		{
			name:      "Write operation",
			path:      "test.txt",
			operation: "write",
			content:   "test content",
		},
		{
			name:      "List operation",
			path:      ".",
			operation: "list",
			content:   "",
		},
		{
			name:      "Delete operation",
			path:      "test.txt",
			operation: "delete",
			content:   "",
		},
		{
			name:      "Mkdir operation",
			path:      "test",
			operation: "mkdir",
			content:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation := domain.NewFileOperation(tt.path, tt.operation, tt.content)
			
			if operation.Path != tt.path {
				t.Errorf("FileOperation.Path = %v, want %v", operation.Path, tt.path)
			}
			
			if operation.Operation != tt.operation {
				t.Errorf("FileOperation.Operation = %v, want %v", operation.Operation, tt.operation)
			}
			
			if operation.Content != tt.content {
				t.Errorf("FileOperation.Content = %v, want %v", operation.Content, tt.content)
			}
			
			// Check that Timestamp is set
			now := time.Now()
			if operation.Timestamp.After(now) || operation.Timestamp.Before(now.Add(-time.Second)) {
				t.Errorf("FileOperation.Timestamp not set correctly: %v", operation.Timestamp)
			}
			
			// Check that Success is false by default
			if operation.Success {
				t.Errorf("FileOperation.Success = %v, want %v", operation.Success, false)
			}
			
			// Check that Error is empty by default
			if operation.Error != "" {
				t.Errorf("FileOperation.Error = %v, want %v", operation.Error, "")
			}
		})
	}
}

func TestSetResult(t *testing.T) {
	tests := []struct {
		name    string
		success bool
		err     error
	}{
		{
			name:    "Success",
			success: true,
			err:     nil,
		},
		{
			name:    "Failure with error",
			success: false,
			err:     errors.New("test error"),
		},
		{
			name:    "Failure without error",
			success: false,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation := domain.NewFileOperation("test.txt", "read", "")
			operation.SetResult(tt.success, tt.err)
			
			if operation.Success != tt.success {
				t.Errorf("FileOperation.Success = %v, want %v", operation.Success, tt.success)
			}
			
			if tt.err != nil {
				if operation.Error != tt.err.Error() {
					t.Errorf("FileOperation.Error = %v, want %v", operation.Error, tt.err.Error())
				}
			} else {
				if operation.Error != "" {
					t.Errorf("FileOperation.Error = %v, want %v", operation.Error, "")
				}
			}
		})
	}
}
