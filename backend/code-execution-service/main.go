package main

import (
	"log"

	"github.com/augment-local-manus-clone/backend/code-execution-service/delivery/http"
	"github.com/augment-local-manus-clone/backend/code-execution-service/infrastructure/docker"
	"github.com/augment-local-manus-clone/backend/code-execution-service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Docker client
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		log.Fatalf("Failed to initialize Docker client: %v", err)
	}

	// Initialize use cases
	executeCodeUseCase := usecase.NewExecuteCodeUseCase(dockerClient)

	// Initialize Gin router
	router := gin.Default()

	// Register HTTP handlers
	http.NewCodeExecutionHandler(router, executeCodeUseCase)

	// Start server
	log.Println("Starting Code Execution Service on :8083")
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
