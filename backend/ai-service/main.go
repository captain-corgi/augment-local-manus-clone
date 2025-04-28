package main

import (
	"log"

	"github.com/augment-local-manus-clone/backend/ai-service/delivery/http"
	"github.com/augment-local-manus-clone/backend/ai-service/infrastructure/llm"
	"github.com/augment-local-manus-clone/backend/ai-service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Ollama client
	ollamaClient, err := llm.NewOllamaClient("http://localhost:11434", "deepseek-r1")
	if err != nil {
		log.Fatalf("Failed to initialize Ollama client: %v", err)
	}

	// Initialize use cases
	processAIRequestUseCase := usecase.NewProcessAIRequestUseCase(ollamaClient)

	// Initialize Gin router
	router := gin.Default()

	// Register HTTP handlers
	http.NewAIHandler(router, processAIRequestUseCase)

	// Start server
	log.Println("Starting AI Service on :8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
