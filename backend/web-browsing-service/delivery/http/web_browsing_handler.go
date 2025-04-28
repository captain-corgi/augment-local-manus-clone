package http

import (
	"net/http"

	"github.com/augment-local-manus-clone/backend/web-browsing-service/domain"
	"github.com/augment-local-manus-clone/backend/web-browsing-service/usecase"
	"github.com/gin-gonic/gin"
)

// WebBrowsingHandler handles HTTP requests for web browsing
type WebBrowsingHandler struct {
	browseWebUseCase       *usecase.BrowseWebUseCase
	searchWebUseCase       *usecase.SearchWebUseCase
	interactWithWebUseCase *usecase.InteractWithWebUseCase
}

// NewWebBrowsingHandler creates a new WebBrowsingHandler
func NewWebBrowsingHandler(
	router *gin.Engine,
	browseWebUseCase *usecase.BrowseWebUseCase,
	searchWebUseCase *usecase.SearchWebUseCase,
	interactWithWebUseCase *usecase.InteractWithWebUseCase,
) *WebBrowsingHandler {
	handler := &WebBrowsingHandler{
		browseWebUseCase:       browseWebUseCase,
		searchWebUseCase:       searchWebUseCase,
		interactWithWebUseCase: interactWithWebUseCase,
	}

	// Register routes
	router.POST("/web/browse", handler.BrowseWeb)
	router.POST("/web/search", handler.SearchWeb)
	router.POST("/web/interact", handler.InteractWithWeb)

	return handler
}

// BrowseWeb handles browsing a web page
func (h *WebBrowsingHandler) BrowseWeb(c *gin.Context) {
	var request domain.WebBrowsingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.browseWebUseCase.Execute(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// SearchWeb handles searching the web
func (h *WebBrowsingHandler) SearchWeb(c *gin.Context) {
	var request domain.WebSearchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.searchWebUseCase.Execute(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// InteractWithWeb handles interacting with a web page
func (h *WebBrowsingHandler) InteractWithWeb(c *gin.Context) {
	var request domain.WebInteractionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.interactWithWebUseCase.Execute(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
