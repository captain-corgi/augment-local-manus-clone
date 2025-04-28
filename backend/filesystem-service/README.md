# Filesystem Service

This microservice performs secure file operations within a restricted workspace directory. It provides a safe interface for the AI agent to interact with the filesystem.

## Architecture

The Filesystem Service follows Clean Architecture principles and Domain-Driven Design (DDD):

- **Domain Layer**: Contains the core business entities and interfaces
- **Use Case Layer**: Implements the business logic
- **Infrastructure Layer**: Provides implementations for filesystem operations
- **Delivery Layer**: Handles HTTP requests and responses

## Setup

1. Install Go: `go version >= 1.21`
2. Initialize: `go mod tidy`
3. Run: `go run main.go`

## API Endpoints

- `GET /files`: List files in a directory
- `GET /files/{path}`: Read a file
- `POST /files/{path}`: Write to a file
- `DELETE /files/{path}`: Delete a file
- `POST /files/mkdir/{path}`: Create a directory

## Security

The service restricts all filesystem operations to the `/app/workspace` directory to prevent unauthorized access to the host system. Path traversal attacks are prevented by validating and sanitizing all file paths.

## Testing

All tests follow the Table-Driven Testing approach. Run tests with:

```bash
go test ./...
```

## Dependencies

- Gin: HTTP web framework
