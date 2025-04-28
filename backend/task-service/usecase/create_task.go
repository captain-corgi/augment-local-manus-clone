package usecase

import (
	"github.com/augment-local-manus-clone/backend/task-service/domain"
)

// CreateTaskInput represents the input for creating a task
type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Input       string `json:"input"`
}

// CreateTaskUseCase handles the creation of tasks
type CreateTaskUseCase struct {
	taskRepo domain.TaskRepository
}

// NewCreateTaskUseCase creates a new instance of CreateTaskUseCase
func NewCreateTaskUseCase(taskRepo domain.TaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		taskRepo: taskRepo,
	}
}

// Execute creates a new task
func (uc *CreateTaskUseCase) Execute(input CreateTaskInput) (*domain.Task, error) {
	task, err := domain.NewTask(input.Title, input.Description, input.Input)
	if err != nil {
		return nil, err
	}

	if err := uc.taskRepo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}
