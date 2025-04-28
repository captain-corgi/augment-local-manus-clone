#!/bin/bash

# Monitor Health script for Local AI Agent System
# This script monitors the health of all services

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

print_info "Monitoring health of Local AI Agent System services..."

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

# Service names
service_names=(
  "Task Service"
  "AI Service"
  "Code Execution Service"
  "Web Browsing Service"
  "Filesystem Service"
  "Shell"
  "Task Management App"
  "Result Viewer App"
)

# Check each service
all_services_healthy=true
unhealthy_services=()

for i in "${!services[@]}"; do
  service_url="${services[$i]}"
  service_name="${service_names[$i]}"
  
  print_info "Checking $service_name at $service_url..."
  
  if curl --output /dev/null --silent --head --fail --max-time 5 "$service_url"; then
    print_success "$service_name is healthy"
  else
    print_error "$service_name is not healthy"
    all_services_healthy=false
    unhealthy_services+=("$service_name")
  fi
done

if [ "$all_services_healthy" = true ]; then
  print_success "All services are healthy"
  exit 0
else
  print_error "The following services are not healthy:"
  for service in "${unhealthy_services[@]}"; do
    print_error "- $service"
  done
  
  # Provide troubleshooting suggestions
  print_info "Troubleshooting suggestions:"
  print_info "1. Check if all services are running: docker-compose ps"
  print_info "2. Check service logs: docker-compose logs <service-name>"
  print_info "3. Restart unhealthy services: docker-compose restart <service-name>"
  print_info "4. Restart all services: docker-compose down && docker-compose up -d"
  
  exit 1
fi
