package domain_test

import (
	"testing"
	"time"

	"github.com/augment-local-manus-clone/backend/task-service/domain"
)

func TestNewTask(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		input       string
		wantErr     bool
	}{
		{
			name:        "Valid task",
			title:       "Test Task",
			description: "This is a test task",
			input:       "test input",
			wantErr:     false,
		},
		{
			name:        "Empty title",
			title:       "",
			description: "This is a test task",
			input:       "test input",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := domain.NewTask(tt.title, tt.description, tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("NewTask() error = nil, wantErr %v", tt.wantErr)
				}
				return
			}
			
			if err != nil {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if task.Title != tt.title {
				t.Errorf("Task.Title = %v, want %v", task.Title, tt.title)
			}
			
			if task.Description != tt.description {
				t.Errorf("Task.Description = %v, want %v", task.Description, tt.description)
			}
			
			if task.Input != tt.input {
				t.Errorf("Task.Input = %v, want %v", task.Input, tt.input)
			}
			
			if task.Status != domain.TaskStatusPending {
				t.Errorf("Task.Status = %v, want %v", task.Status, domain.TaskStatusPending)
			}
			
			// Check that CreatedAt and UpdatedAt are set
			now := time.Now()
			if task.CreatedAt.After(now) || task.CreatedAt.Before(now.Add(-time.Second)) {
				t.Errorf("Task.CreatedAt not set correctly: %v", task.CreatedAt)
			}
			
			if task.UpdatedAt.After(now) || task.UpdatedAt.Before(now.Add(-time.Second)) {
				t.Errorf("Task.UpdatedAt not set correctly: %v", task.UpdatedAt)
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	task, _ := domain.NewTask("Test Task", "This is a test task", "test input")
	
	tests := []struct {
		name    string
		status  domain.TaskStatus
		wantErr bool
	}{
		{
			name:    "Update to running",
			status:  domain.TaskStatusRunning,
			wantErr: false,
		},
		{
			name:    "Update to completed",
			status:  domain.TaskStatusCompleted,
			wantErr: false,
		},
		{
			name:    "Update to failed",
			status:  domain.TaskStatusFailed,
			wantErr: false,
		},
		{
			name:    "Invalid status",
			status:  "invalid",
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalUpdatedAt := task.UpdatedAt
			time.Sleep(1 * time.Millisecond) // Ensure time difference
			
			err := task.UpdateStatus(tt.status)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("UpdateStatus() error = nil, wantErr %v", tt.wantErr)
				}
				return
			}
			
			if err != nil {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if task.Status != tt.status {
				t.Errorf("Task.Status = %v, want %v", task.Status, tt.status)
			}
			
			if !task.UpdatedAt.After(originalUpdatedAt) {
				t.Errorf("Task.UpdatedAt not updated: %v", task.UpdatedAt)
			}
		})
	}
}

func TestSetResult(t *testing.T) {
	task, _ := domain.NewTask("Test Task", "This is a test task", "test input")
	
	tests := []struct {
		name   string
		result string
	}{
		{
			name:   "Set result",
			result: "test result",
		},
		{
			name:   "Update result",
			result: "updated result",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalUpdatedAt := task.UpdatedAt
			time.Sleep(1 * time.Millisecond) // Ensure time difference
			
			task.SetResult(tt.result)
			
			if task.Result != tt.result {
				t.Errorf("Task.Result = %v, want %v", task.Result, tt.result)
			}
			
			if !task.UpdatedAt.After(originalUpdatedAt) {
				t.Errorf("Task.UpdatedAt not updated: %v", task.UpdatedAt)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		task    *domain.Task
		wantErr bool
	}{
		{
			name: "Valid task",
			task: &domain.Task{
				Title:       "Test Task",
				Description: "This is a test task",
			},
			wantErr: false,
		},
		{
			name: "Empty title",
			task: &domain.Task{
				Title:       "",
				Description: "This is a test task",
			},
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.task.Validate()
			
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
