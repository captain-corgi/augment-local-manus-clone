#!/bin/bash

# Setup Development Environment script for Local AI Agent System
# This script sets up the development environment for the system

set -e

# Print colored output
print_info() {
  echo -e "\033[0;34m[INFO]\033[0m $1"
}

print_success() {
  echo -e "\033[0;32m[SUCCESS]\033[0m $1"
}

print_error() {
  echo -e "\033[0;31m[ERROR]\033[0m $1"
}

print_warning() {
  echo -e "\033[0;33m[WARNING]\033[0m $1"
}

# Get the project root directory
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT_DIR"

print_info "Setting up development environment for Local AI Agent System..."

# Check prerequisites
print_info "Checking prerequisites..."

# Check if Go is installed
if command -v go &> /dev/null; then
  GO_VERSION=$(go version | awk '{print $3}')
  print_info "Go is installed: $GO_VERSION"
else
  print_error "Go is not installed. Please install Go 1.21 or later."
  exit 1
fi

# Check if Node.js is installed
if command -v node &> /dev/null; then
  NODE_VERSION=$(node --version)
  print_info "Node.js is installed: $NODE_VERSION"
else
  print_error "Node.js is not installed. Please install Node.js 16 or later."
  exit 1
fi

# Check if Docker is installed
if command -v docker &> /dev/null; then
  DOCKER_VERSION=$(docker --version)
  print_info "Docker is installed: $DOCKER_VERSION"
else
  print_error "Docker is not installed. Please install Docker."
  exit 1
fi

# Check if Ollama is installed
if command -v ollama &> /dev/null; then
  OLLAMA_VERSION=$(ollama --version)
  print_info "Ollama is installed: $OLLAMA_VERSION"
else
  print_warning "Ollama is not installed. The AI Service requires Ollama."
  print_info "You can install Ollama from https://ollama.ai/"
fi

# Create workspace directory
print_info "Creating workspace directory..."
mkdir -p "$ROOT_DIR/workspace"

# Set up backend services
print_info "Setting up backend services..."

backend_services=("task-service" "ai-service" "code-execution-service" "web-browsing-service" "filesystem-service")

for service in "${backend_services[@]}"; do
  print_info "Setting up $service..."
  cd "$ROOT_DIR/backend/$service"
  
  # Initialize Go module
  go mod tidy
  
  if [ $? -eq 0 ]; then
    print_success "$service set up successfully"
  else
    print_error "Failed to set up $service"
    exit 1
  fi
done

# Set up frontend applications
print_info "Setting up frontend applications..."

frontend_apps=("shell" "task-management-app" "result-viewer-app")

for app in "${frontend_apps[@]}"; do
  print_info "Setting up $app..."
  cd "$ROOT_DIR/frontend/$app"
  
  # Install dependencies
  npm install
  
  if [ $? -eq 0 ]; then
    print_success "$app set up successfully"
  else
    print_error "Failed to set up $app"
    exit 1
  fi
done

# Pull Ollama model if Ollama is installed
if command -v ollama &> /dev/null; then
  print_info "Pulling DeepSeek-R1 model for Ollama..."
  ollama pull deepseek-r1
  
  if [ $? -eq 0 ]; then
    print_success "DeepSeek-R1 model pulled successfully"
  else
    print_warning "Failed to pull DeepSeek-R1 model. The AI Service may not work correctly."
  fi
fi

print_success "Development environment set up successfully"
print_info "You can now start the services:"
print_info "1. Backend services: cd backend/<service> && go run main.go"
print_info "2. Frontend applications: cd frontend/<app> && npm run dev"
print_info "3. Or use docker-compose: docker-compose -f docker-compose.dev.yml up"
