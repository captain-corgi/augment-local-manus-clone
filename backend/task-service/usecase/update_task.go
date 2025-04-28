package usecase

import (
	"errors"

	"github.com/augment-local-manus-clone/backend/task-service/domain"
)

// UpdateTaskInput represents the input for updating a task
type UpdateTaskInput struct {
	ID     string           `json:"id"`
	Status domain.TaskStatus `json:"status"`
	Result string           `json:"result"`
}

// UpdateTaskUseCase handles updating a task
type UpdateTaskUseCase struct {
	taskRepo domain.TaskRepository
}

// NewUpdateTaskUseCase creates a new instance of UpdateTaskUseCase
func NewUpdateTaskUseCase(taskRepo domain.TaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskRepo: taskRepo,
	}
}

// Execute updates a task
func (uc *UpdateTaskUseCase) Execute(input UpdateTaskInput) (*domain.Task, error) {
	if input.ID == "" {
		return nil, errors.New("task ID cannot be empty")
	}

	task, err := uc.taskRepo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	if input.Status != "" {
		if err := task.UpdateStatus(input.Status); err != nil {
			return nil, err
		}
	}

	if input.Result != "" {
		task.SetResult(input.Result)
	}

	if err := uc.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}
