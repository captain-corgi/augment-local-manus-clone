package usecase

import (
	"github.com/augment-local-manus-clone/backend/task-service/domain"
)

// ListTasksUseCase handles retrieving all tasks
type ListTasksUseCase struct {
	taskRepo domain.TaskRepository
}

// NewListTasksUseCase creates a new instance of ListTasksUseCase
func NewListTasksUseCase(taskRepo domain.TaskRepository) *ListTasksUseCase {
	return &ListTasksUseCase{
		taskRepo: taskRepo,
	}
}

// Execute retrieves all tasks
func (uc *ListTasksUseCase) Execute() ([]*domain.Task, error) {
	return uc.taskRepo.List()
}
