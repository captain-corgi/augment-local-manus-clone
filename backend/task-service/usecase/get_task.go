package usecase

import (
	"errors"

	"github.com/augment-local-manus-clone/backend/task-service/domain"
)

// GetTaskUseCase handles retrieving a task by ID
type GetTaskUseCase struct {
	taskRepo domain.TaskRepository
}

// NewGetTaskUseCase creates a new instance of GetTaskUseCase
func NewGetTaskUseCase(taskRepo domain.TaskRepository) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskRepo: taskRepo,
	}
}

// Execute retrieves a task by its ID
func (uc *GetTaskUseCase) Execute(id string) (*domain.Task, error) {
	if id == "" {
		return nil, errors.New("task ID cannot be empty")
	}

	return uc.taskRepo.GetByID(id)
}
