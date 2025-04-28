package main

import (
	"log"

	"github.com/augment-local-manus-clone/backend/web-browsing-service/delivery/http"
	"github.com/augment-local-manus-clone/backend/web-browsing-service/infrastructure/browser"
	"github.com/augment-local-manus-clone/backend/web-browsing-service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize browser client
	browserClient, err := browser.NewChromeDPClient()
	if err != nil {
		log.Fatalf("Failed to initialize browser client: %v", err)
	}
	defer browserClient.Close()

	// Initialize use cases
	browseWebUseCase := usecase.NewBrowseWebUseCase(browserClient)
	searchWebUseCase := usecase.NewSearchWebUseCase(browserClient)
	interactWithWebUseCase := usecase.NewInteractWithWebUseCase(browserClient)

	// Initialize Gin router
	router := gin.Default()

	// Register HTTP handlers
	http.NewWebBrowsingHandler(
		router,
		browseWebUseCase,
		searchWebUseCase,
		interactWithWebUseCase,
	)

	// Start server
	log.Println("Starting Web Browsing Service on :8084")
	if err := router.Run(":8084"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
