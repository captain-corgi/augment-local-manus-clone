package main

import (
	"log"
	"os"

	"github.com/augment-local-manus-clone/backend/filesystem-service/delivery/http"
	"github.com/augment-local-manus-clone/backend/filesystem-service/infrastructure/fs"
	"github.com/augment-local-manus-clone/backend/filesystem-service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Define workspace directory
	workspaceDir := "/app/workspace"
	
	// For development, use a local directory if /app/workspace doesn't exist
	if _, err := os.Stat(workspaceDir); os.IsNotExist(err) {
		workspaceDir = "./workspace"
		// Create workspace directory if it doesn't exist
		if _, err := os.Stat(workspaceDir); os.IsNotExist(err) {
			if err := os.MkdirAll(workspaceDir, 0755); err != nil {
				log.Fatalf("Failed to create workspace directory: %v", err)
			}
		}
	}

	// Initialize filesystem client
	filesystemClient := fs.NewFilesystemClient(workspaceDir)

	// Initialize use cases
	readFileUseCase := usecase.NewReadFileUseCase(filesystemClient)
	writeFileUseCase := usecase.NewWriteFileUseCase(filesystemClient)
	listFilesUseCase := usecase.NewListFilesUseCase(filesystemClient)
	deleteFileUseCase := usecase.NewDeleteFileUseCase(filesystemClient)
	makeDirectoryUseCase := usecase.NewMakeDirectoryUseCase(filesystemClient)

	// Initialize Gin router
	router := gin.Default()

	// Register HTTP handlers
	http.NewFilesystemHandler(
		router,
		readFileUseCase,
		writeFileUseCase,
		listFilesUseCase,
		deleteFileUseCase,
		makeDirectoryUseCase,
	)

	// Start server
	log.Println("Starting Filesystem Service on :8085")
	if err := router.Run(":8085"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
