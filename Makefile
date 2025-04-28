.PHONY: help setup build test run run-frontend stop clean backup restore update lint

# Default target
help:
	@echo "Local AI Agent System Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make setup       - Set up the development environment"
	@echo "  make build       - Build all components"
	@echo "  make test        - Run all tests"
	@echo "  make run         - Run the system using Docker Compose"
	@echo "  make run-frontend - Run only the frontend applications in development mode"
	@echo "  make stop        - Stop the system"
	@echo "  make clean       - Clean up Docker resources"
	@echo "  make backup      - Create a backup of the system data"
	@echo "  make restore     - Restore from a backup"
	@echo "  make update      - Update the system to the latest version"
	@echo "  make lint        - Run linters on the codebase"
	@echo ""

# Set up the development environment
setup:
	@echo "Setting up development environment..."
	@./cicd/scripts/setup-dev.sh

# Build all components
build:
	@echo "Building all components..."
	@./cicd/scripts/build.sh

# Run all tests
test:
	@echo "Running all tests..."
	@./cicd/scripts/test.sh

# Run the system using Docker Compose
run:
	@echo "Starting the system..."
	@docker-compose up -d
	@echo "System started. Access the web interface at http://localhost:5000"

# Stop the system
stop:
	@echo "Stopping the system..."
	@docker-compose down

# Clean up Docker resources
clean:
	@echo "Cleaning up Docker resources..."
	@./cicd/scripts/cleanup.sh --all

# Create a backup of the system data
backup:
	@echo "Creating backup..."
	@./cicd/scripts/backup.sh

# Restore from a backup
restore:
	@echo "Restoring from backup..."
	@./cicd/scripts/restore.sh

# Run only the frontend applications in development mode
run-frontend:
	@echo "Starting frontend applications in development mode..."
	@echo "Setting up shell application..."
	@cd frontend/shell && npm install
	@echo "Starting shell application..."
	@cd frontend/shell && npm run dev &
	@echo "Setting up task management application..."
	@cd frontend/task-management-app && npm install
	@echo "Starting task management application..."
	@cd frontend/task-management-app && npm run dev &
	@echo "Setting up result viewer application..."
	@cd frontend/result-viewer-app && npm install
	@echo "Starting result viewer application..."
	@cd frontend/result-viewer-app && npm run dev &
	@echo "Frontend applications started."
	@echo "Access the shell application at http://localhost:5000"
	@echo "Access the task management application at http://localhost:5001"
	@echo "Access the result viewer application at http://localhost:5002"
	@echo "Press Ctrl+C to stop all applications."
	@wait

# Update the system to the latest version
update:
	@echo "Updating the system..."
	@./cicd/scripts/update.sh

# Run linters on the codebase
lint:
	@echo "Running linters..."
	@echo "Running Go linters..."
	@for service in backend/*; do \
		if [ -d "$$service" ]; then \
			echo "Linting $$service..."; \
			cd $$service && golangci-lint run ./... && cd -; \
		fi \
	done
	@echo "Running JavaScript/TypeScript linters..."
	@for app in frontend/*; do \
		if [ -d "$$app" ]; then \
			echo "Linting $$app..."; \
			cd $$app && npm run lint && cd -; \
		fi \
	done
