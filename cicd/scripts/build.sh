#!/bin/bash

# Build script for Local AI Agent System
# This script builds all components of the system

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

# Get the project root directory
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT_DIR"

print_info "Building Local AI Agent System..."

# Build backend services
print_info "Building backend services..."

backend_services=("task-service" "ai-service" "code-execution-service" "web-browsing-service" "filesystem-service")

for service in "${backend_services[@]}"; do
  print_info "Building $service..."
  cd "$ROOT_DIR/backend/$service"
  
  # Check if Go is installed
  if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go and try again."
    exit 1
  fi
  
  # Build the service
  go mod tidy
  go build -o bin/$service
  
  if [ $? -eq 0 ]; then
    print_success "$service built successfully"
  else
    print_error "Failed to build $service"
    exit 1
  fi
done

# Build frontend applications
print_info "Building frontend applications..."

frontend_apps=("shell" "task-management-app" "result-viewer-app")

for app in "${frontend_apps[@]}"; do
  print_info "Building $app..."
  cd "$ROOT_DIR/frontend/$app"
  
  # Check if Node.js is installed
  if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed. Please install Node.js and try again."
    exit 1
  fi
  
  # Install dependencies and build the app
  npm ci
  npm run build
  
  if [ $? -eq 0 ]; then
    print_success "$app built successfully"
  else
    print_error "Failed to build $app"
    exit 1
  fi
done

print_success "All components built successfully"
