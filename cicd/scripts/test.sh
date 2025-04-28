#!/bin/bash

# Test script for Local AI Agent System
# This script runs all tests for the system

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

print_info "Running tests for Local AI Agent System..."

# Test backend services
print_info "Testing backend services..."

backend_services=("task-service" "ai-service" "code-execution-service" "web-browsing-service" "filesystem-service")

for service in "${backend_services[@]}"; do
  print_info "Testing $service..."
  cd "$ROOT_DIR/backend/$service"
  
  # Check if Go is installed
  if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go and try again."
    exit 1
  fi
  
  # Run tests
  go test -v ./...
  
  if [ $? -eq 0 ]; then
    print_success "$service tests passed"
  else
    print_error "$service tests failed"
    exit 1
  fi
done

# Test frontend applications
print_info "Testing frontend applications..."

frontend_apps=("shell" "task-management-app" "result-viewer-app")

for app in "${frontend_apps[@]}"; do
  print_info "Testing $app..."
  cd "$ROOT_DIR/frontend/$app"
  
  # Check if Node.js is installed
  if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed. Please install Node.js and try again."
    exit 1
  fi
  
  # Run tests
  npm test
  
  if [ $? -eq 0 ]; then
    print_success "$app tests passed"
  else
    print_error "$app tests failed"
    exit 1
  fi
done

# Run integration tests
print_info "Running integration tests..."

# Check if all services are running
print_info "Checking if all services are running..."

# Array of service URLs to check
services=(
  "http://localhost:8081/health"  # Task Service
  "http://localhost:8082/health"  # AI Service
  "http://localhost:8083/health"  # Code Execution Service
  "http://localhost:8084/health"  # Web Browsing Service
  "http://localhost:8085/health"  # Filesystem Service
  "http://localhost:5000"         # Shell
  "http://localhost:5001"         # Task Management App
  "http://localhost:5002"         # Result Viewer App
)

# Check each service
all_services_running=true
for service_url in "${services[@]}"; do
  if curl --output /dev/null --silent --head --fail "$service_url"; then
    print_info "$service_url is running"
  else
    print_error "$service_url is not running"
    all_services_running=false
  fi
done

if [ "$all_services_running" = false ]; then
  print_error "Not all services are running. Please start all services before running integration tests."
  exit 1
fi

# Run API integration tests
print_info "Running API integration tests..."
cd "$ROOT_DIR/tests/integration/api"
go test -v ./...

if [ $? -eq 0 ]; then
  print_success "API integration tests passed"
else
  print_error "API integration tests failed"
  exit 1
fi

# Run UI integration tests
print_info "Running UI integration tests..."
cd "$ROOT_DIR/tests/integration/ui"
npx cypress run

if [ $? -eq 0 ]; then
  print_success "UI integration tests passed"
else
  print_error "UI integration tests failed"
  exit 1
fi

print_success "All tests passed successfully"
