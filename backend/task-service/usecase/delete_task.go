package usecase

import (
	"errors"

	"github.com/augment-local-manus-clone/backend/task-service/domain"
)

// DeleteTaskUseCase handles deleting a task
type DeleteTaskUseCase struct {
	taskRepo domain.TaskRepository
}

// NewDeleteTaskUseCase creates a new instance of DeleteTaskUseCase
func NewDeleteTaskUseCase(taskRepo domain.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskRepo: taskRepo,
	}
}

// Execute deletes a task by its ID
func (uc *DeleteTaskUseCase) Execute(id string) error {
	if id == "" {
		return errors.New("task ID cannot be empty")
	}

	// Check if the task exists
	_, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	return uc.taskRepo.Delete(id)
}
