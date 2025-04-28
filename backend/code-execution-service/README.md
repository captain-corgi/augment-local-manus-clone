# Code Execution Service

This microservice executes code in sandboxed Docker containers. It provides a secure environment for running code snippets in various programming languages.

## Architecture

The Code Execution Service follows Clean Architecture principles and Domain-Driven Design (DDD):

- **Domain Layer**: Contains the core business entities and interfaces
- **Use Case Layer**: Implements the business logic
- **Infrastructure Layer**: Provides implementations for external services (Docker client)
- **Delivery Layer**: Handles HTTP requests and responses

## Setup

1. Install Go: `go version >= 1.21`
2. Install Docker: Follow instructions at [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
3. Initialize: `go mod tidy`
4. Run: `go run main.go`

## API Endpoints

- `POST /code/execute`: Execute code in a sandboxed environment

## Supported Languages

- Python
- JavaScript (Node.js)
- Go
- Ruby
- Java

## Security

The service uses Docker containers to sandbox code execution with:

- Limited resources (CPU, memory)
- No network access
- Read-only filesystem (except for the code file)
- Limited execution time

## Testing

All tests follow the Table-Driven Testing approach. Run tests with:

```bash
go test ./...
```

## Dependencies

- Gin: HTTP web framework
- Docker Go client: Interface with Docker for container management
