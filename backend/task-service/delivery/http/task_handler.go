package http

import (
	"net/http"

	"github.com/augment-local-manus-clone/backend/task-service/usecase"
	"github.com/gin-gonic/gin"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	createTaskUseCase *usecase.CreateTaskUseCase
	getTaskUseCase    *usecase.GetTaskUseCase
	listTasksUseCase  *usecase.ListTasksUseCase
	updateTaskUseCase *usecase.UpdateTaskUseCase
	deleteTaskUseCase *usecase.DeleteTaskUseCase
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(
	router *gin.Engine,
	createTaskUseCase *usecase.CreateTaskUseCase,
	getTaskUseCase *usecase.GetTaskUseCase,
	listTasksUseCase *usecase.ListTasksUseCase,
	updateTaskUseCase *usecase.UpdateTaskUseCase,
	deleteTaskUseCase *usecase.DeleteTaskUseCase,
) *TaskHandler {
	handler := &TaskHandler{
		createTaskUseCase: createTaskUseCase,
		getTaskUseCase:    getTaskUseCase,
		listTasksUseCase:  listTasksUseCase,
		updateTaskUseCase: updateTaskUseCase,
		deleteTaskUseCase: deleteTaskUseCase,
	}

	// Register routes
	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks/:id", handler.GetTask)
	router.GET("/tasks", handler.ListTasks)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	return handler
}

// CreateTask handles the creation of a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input usecase.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.createTaskUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetTask handles retrieving a task by ID
func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.getTaskUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// ListTasks handles retrieving all tasks
func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.listTasksUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// UpdateTask handles updating a task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var input usecase.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = id

	task, err := h.updateTaskUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask handles deleting a task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.deleteTaskUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
