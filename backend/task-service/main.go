package main

import (
	"log"

	"github.com/augment-local-manus-clone/backend/task-service/delivery/http"
	"github.com/augment-local-manus-clone/backend/task-service/infrastructure/repository"
	"github.com/augment-local-manus-clone/backend/task-service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize SQLite repository
	taskRepo, err := repository.NewSQLiteTaskRepository("./tasks.db")
	if err != nil {
		log.Fatalf("Failed to initialize task repository: %v", err)
	}

	// Initialize use cases
	createTaskUseCase := usecase.NewCreateTaskUseCase(taskRepo)
	getTaskUseCase := usecase.NewGetTaskUseCase(taskRepo)
	listTasksUseCase := usecase.NewListTasksUseCase(taskRepo)
	updateTaskUseCase := usecase.NewUpdateTaskUseCase(taskRepo)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCase(taskRepo)

	// Initialize Gin router
	router := gin.Default()

	// Register HTTP handlers
	http.NewTaskHandler(
		router,
		createTaskUseCase,
		getTaskUseCase,
		listTasksUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
	)

	// Start server
	log.Println("Starting Task Service on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
