package domain

import (
	"errors"
	"time"
)

// TaskStatus represents the current status of a task
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

// Task represents a task in the system
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	Input       string     `json:"input"`
	Result      string     `json:"result"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// NewTask creates a new task with the given title and description
func NewTask(title, description, input string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	return &Task{
		Title:       title,
		Description: description,
		Status:      TaskStatusPending,
		Input:       input,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// UpdateStatus updates the task status
func (t *Task) UpdateStatus(status TaskStatus) error {
	if status != TaskStatusPending && status != TaskStatusRunning &&
		status != TaskStatusCompleted && status != TaskStatusFailed {
		return errors.New("invalid task status")
	}

	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}

// SetResult sets the task result
func (t *Task) SetResult(result string) {
	t.Result = result
	t.UpdatedAt = time.Now()
}

// Validate validates the task
func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}
