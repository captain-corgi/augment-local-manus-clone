# AI Service

This microservice interfaces with the local LLM (Ollama) for reasoning and planning. It processes AI requests and returns responses from the local model.

## Architecture

The AI Service follows Clean Architecture principles and Domain-Driven Design (DDD):

- **Domain Layer**: Contains the core business entities and interfaces
- **Use Case Layer**: Implements the business logic
- **Infrastructure Layer**: Provides implementations for external services (Ollama client)
- **Delivery Layer**: Handles HTTP requests and responses

## Setup

1. Install Go: `go version >= 1.21`
2. Install Ollama: Follow instructions at [https://ollama.ai/](https://ollama.ai/)
3. Pull the required model: `ollama pull deepseek-r1`
4. Initialize: `go mod tidy`
5. Run: `go run main.go`

## API Endpoints

- `POST /ai/process`: Process an AI request

## Testing

All tests follow the Table-Driven Testing approach. Run tests with:

```bash
go test ./...
```

## Dependencies

- Gin: HTTP web framework
- Ollama Go client: Interface with the local LLM
