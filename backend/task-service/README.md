# Task Service

This microservice manages the lifecycle of tasks, including submission, status tracking, and result storage. It uses SQLite for persistence and communicates with other services via HTTP/JSON APIs.

## Architecture

The Task Service follows Clean Architecture principles and Domain-Driven Design (DDD):

- **Domain Layer**: Contains the core business entities and interfaces
- **Use Case Layer**: Implements the business logic
- **Infrastructure Layer**: Provides implementations for repositories and external services
- **Delivery Layer**: Handles HTTP requests and responses

## Setup

1. Install Go: `go version >= 1.21`
2. Initialize: `go mod tidy`
3. Run: `go run main.go`

## API Endpoints

- `POST /tasks`: Submit a new task
- `GET /tasks/{id}`: Get task details
- `GET /tasks`: List all tasks
- `PUT /tasks/{id}`: Update task status
- `DELETE /tasks/{id}`: Delete a task

## Testing

All tests follow the Table-Driven Testing approach. Run tests with:

```bash
go test ./...
```

## Dependencies

- Gin: HTTP web framework
- SQLite: Local database storage
- go-sqlite3: SQLite driver for Go
