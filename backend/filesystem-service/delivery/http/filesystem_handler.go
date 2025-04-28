package http

import (
	"net/http"

	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
	"github.com/augment-local-manus-clone/backend/filesystem-service/usecase"
	"github.com/gin-gonic/gin"
)

// FilesystemHandler handles HTTP requests for filesystem operations
type FilesystemHandler struct {
	readFileUseCase      *usecase.ReadFileUseCase
	writeFileUseCase     *usecase.WriteFileUseCase
	listFilesUseCase     *usecase.ListFilesUseCase
	deleteFileUseCase    *usecase.DeleteFileUseCase
	makeDirectoryUseCase *usecase.MakeDirectoryUseCase
}

// FileWriteRequest represents a request to write to a file
type FileWriteRequest struct {
	Content string `json:"content" binding:"required"`
}

// NewFilesystemHandler creates a new FilesystemHandler
func NewFilesystemHandler(
	router *gin.Engine,
	readFileUseCase *usecase.ReadFileUseCase,
	writeFileUseCase *usecase.WriteFileUseCase,
	listFilesUseCase *usecase.ListFilesUseCase,
	deleteFileUseCase *usecase.DeleteFileUseCase,
	makeDirectoryUseCase *usecase.MakeDirectoryUseCase,
) *FilesystemHandler {
	handler := &FilesystemHandler{
		readFileUseCase:      readFileUseCase,
		writeFileUseCase:     writeFileUseCase,
		listFilesUseCase:     listFilesUseCase,
		deleteFileUseCase:    deleteFileUseCase,
		makeDirectoryUseCase: makeDirectoryUseCase,
	}

	// Register routes
	router.GET("/files", handler.ListFiles)
	router.GET("/files/*path", handler.ReadFile)
	router.POST("/files/*path", handler.WriteFile)
	router.DELETE("/files/*path", handler.DeleteFile)
	router.POST("/files/mkdir/*path", handler.MakeDirectory)

	return handler
}

// ReadFile handles reading a file
func (h *FilesystemHandler) ReadFile(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
		return
	}

	// Remove leading slash
	if path[0] == '/' {
		path = path[1:]
	}

	content, err := h.readFileUseCase.Execute(path)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case domain.ErrPathTraversal:
			statusCode = http.StatusForbidden
		case domain.ErrPathNotFound:
			statusCode = http.StatusNotFound
		case domain.ErrNotAFile:
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, content)
}

// WriteFile handles writing to a file
func (h *FilesystemHandler) WriteFile(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
		return
	}

	// Remove leading slash
	if path[0] == '/' {
		path = path[1:]
	}

	var request FileWriteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operation, err := h.writeFileUseCase.Execute(path, request.Content)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case domain.ErrPathTraversal:
			statusCode = http.StatusForbidden
		case domain.ErrPermissionDenied:
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, operation)
}

// ListFiles handles listing files in a directory
func (h *FilesystemHandler) ListFiles(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		path = "."
	}

	files, err := h.listFilesUseCase.Execute(path)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case domain.ErrPathTraversal:
			statusCode = http.StatusForbidden
		case domain.ErrPathNotFound:
			statusCode = http.StatusNotFound
		case domain.ErrNotADirectory:
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}

// DeleteFile handles deleting a file
func (h *FilesystemHandler) DeleteFile(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
		return
	}

	// Remove leading slash
	if path[0] == '/' {
		path = path[1:]
	}

	operation, err := h.deleteFileUseCase.Execute(path)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case domain.ErrPathTraversal:
			statusCode = http.StatusForbidden
		case domain.ErrPathNotFound:
			statusCode = http.StatusNotFound
		case domain.ErrPermissionDenied:
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, operation)
}

// MakeDirectory handles creating a directory
func (h *FilesystemHandler) MakeDirectory(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
		return
	}

	// Remove leading slash
	if path[0] == '/' {
		path = path[1:]
	}

	operation, err := h.makeDirectoryUseCase.Execute(path)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case domain.ErrPathTraversal:
			statusCode = http.StatusForbidden
		case domain.ErrPermissionDenied:
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, operation)
}
