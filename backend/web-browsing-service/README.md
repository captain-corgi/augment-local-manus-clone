# Web Browsing Service

This microservice provides web browsing capabilities using headless Chrome via chromedp. It allows the AI agent to navigate web pages, extract content, and interact with web elements.

## Architecture

The Web Browsing Service follows Clean Architecture principles and Domain-Driven Design (DDD):

- **Domain Layer**: Contains the core business entities and interfaces
- **Use Case Layer**: Implements the business logic
- **Infrastructure Layer**: Provides implementations for external services (chromedp)
- **Delivery Layer**: Handles HTTP requests and responses

## Setup

1. Install Go: `go version >= 1.21`
2. Install Chrome or Chromium: Required for headless browsing
3. Initialize: `go mod tidy`
4. Run: `go run main.go`

## API Endpoints

- `POST /web/browse`: Navigate to a URL and extract content
- `POST /web/search`: Perform a search query and extract results
- `POST /web/interact`: Interact with elements on a web page

## Features

- Page navigation
- Content extraction
- Element interaction (click, type, etc.)
- Screenshot capture
- Search engine queries

## Testing

All tests follow the Table-Driven Testing approach. Run tests with:

```bash
go test ./...
```

## Dependencies

- Gin: HTTP web framework
- chromedp: Go library for driving browsers
