package domain

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	// Create stores a new task
	Create(task *Task) error

	// GetByID retrieves a task by its ID
	GetByID(id string) (*Task, error)

	// List retrieves all tasks
	List() ([]*Task, error)

	// Update updates an existing task
	Update(task *Task) error

	// Delete removes a task by its ID
	Delete(id string) error
}
