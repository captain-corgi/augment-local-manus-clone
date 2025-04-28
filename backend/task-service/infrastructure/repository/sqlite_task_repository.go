package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/augment-local-manus-clone/backend/task-service/domain"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteTaskRepository implements the TaskRepository interface using SQLite
type SQLiteTaskRepository struct {
	db *sql.DB
}

// NewSQLiteTaskRepository creates a new SQLiteTaskRepository
func NewSQLiteTaskRepository(dbPath string) (*SQLiteTaskRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create tasks table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL,
			input TEXT,
			result TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create tasks table: %w", err)
	}

	return &SQLiteTaskRepository{
		db: db,
	}, nil
}

// Create stores a new task
func (r *SQLiteTaskRepository) Create(task *domain.Task) error {
	// Generate a unique ID if not provided
	if task.ID == "" {
		task.ID = fmt.Sprintf("task_%d", time.Now().UnixNano())
	}

	_, err := r.db.Exec(
		`INSERT INTO tasks (id, title, description, status, input, result, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.Input,
		task.Result,
		task.CreatedAt,
		task.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}

	return nil
}

// GetByID retrieves a task by its ID
func (r *SQLiteTaskRepository) GetByID(id string) (*domain.Task, error) {
	row := r.db.QueryRow(
		`SELECT id, title, description, status, input, result, created_at, updated_at
		FROM tasks WHERE id = ?`,
		id,
	)

	var task domain.Task
	var createdAt, updatedAt string

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Input,
		&task.Result,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task not found: %s", id)
		}
		return nil, fmt.Errorf("failed to scan task: %w", err)
	}

	// Parse timestamps
	task.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	task.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &task, nil
}

// List retrieves all tasks
func (r *SQLiteTaskRepository) List() ([]*domain.Task, error) {
	rows, err := r.db.Query(
		`SELECT id, title, description, status, input, result, created_at, updated_at
		FROM tasks ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*domain.Task

	for rows.Next() {
		var task domain.Task
		var createdAt, updatedAt string

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Input,
			&task.Result,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		// Parse timestamps
		task.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		task.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tasks: %w", err)
	}

	return tasks, nil
}

// Update updates an existing task
func (r *SQLiteTaskRepository) Update(task *domain.Task) error {
	task.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		`UPDATE tasks SET title = ?, description = ?, status = ?, input = ?, result = ?, updated_at = ?
		WHERE id = ?`,
		task.Title,
		task.Description,
		task.Status,
		task.Input,
		task.Result,
		task.UpdatedAt,
		task.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

// Delete removes a task by its ID
func (r *SQLiteTaskRepository) Delete(id string) error {
	result, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task not found: %s", id)
	}

	return nil
}
