package http

import (
	"net/http"

	"github.com/augment-local-manus-clone/backend/ai-service/domain"
	"github.com/augment-local-manus-clone/backend/ai-service/usecase"
	"github.com/gin-gonic/gin"
)

// AIHandler handles HTTP requests for AI processing
type AIHandler struct {
	processAIRequestUseCase *usecase.ProcessAIRequestUseCase
}

// NewAIHandler creates a new AIHandler
func NewAIHandler(router *gin.Engine, processAIRequestUseCase *usecase.ProcessAIRequestUseCase) *AIHandler {
	handler := &AIHandler{
		processAIRequestUseCase: processAIRequestUseCase,
	}

	// Register routes
	router.POST("/ai/process", handler.ProcessAIRequest)

	return handler
}

// ProcessAIRequest handles processing an AI request
func (h *AIHandler) ProcessAIRequest(c *gin.Context) {
	var request domain.AIRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.processAIRequestUseCase.Execute(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
