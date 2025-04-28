package http

import (
	"net/http"

	"github.com/augment-local-manus-clone/backend/code-execution-service/usecase"
	"github.com/gin-gonic/gin"
)

// CodeExecutionHandler handles HTTP requests for code execution
type CodeExecutionHandler struct {
	executeCodeUseCase *usecase.ExecuteCodeUseCase
}

// NewCodeExecutionHandler creates a new CodeExecutionHandler
func NewCodeExecutionHandler(router *gin.Engine, executeCodeUseCase *usecase.ExecuteCodeUseCase) *CodeExecutionHandler {
	handler := &CodeExecutionHandler{
		executeCodeUseCase: executeCodeUseCase,
	}

	// Register routes
	router.POST("/code/execute", handler.ExecuteCode)

	return handler
}

// ExecuteCode handles executing code
func (h *CodeExecutionHandler) ExecuteCode(c *gin.Context) {
	var input usecase.ExecuteCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	execution, err := h.executeCodeUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, execution)
}
