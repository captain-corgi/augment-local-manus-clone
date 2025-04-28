#!/bin/bash

# Deploy script for Local AI Agent System
# This script deploys the system to the target environment

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

# Parse command line arguments
ENVIRONMENT="development"
SKIP_TESTS=false

while [[ $# -gt 0 ]]; do
  case $1 in
    --environment|-e)
      ENVIRONMENT="$2"
      shift 2
      ;;
    --skip-tests)
      SKIP_TESTS=true
      shift
      ;;
    --help|-h)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  --environment, -e <env>  Target environment (development, production)"
      echo "  --skip-tests             Skip running tests before deployment"
      echo "  --help, -h               Show this help message"
      exit 0
      ;;
    *)
      print_error "Unknown option: $1"
      exit 1
      ;;
  esac
done

print_info "Deploying Local AI Agent System to $ENVIRONMENT environment..."

# Run tests if not skipped
if [ "$SKIP_TESTS" = false ]; then
  print_info "Running tests before deployment..."
  "$ROOT_DIR/cicd/scripts/test.sh"
  
  if [ $? -ne 0 ]; then
    print_error "Tests failed. Aborting deployment."
    exit 1
  fi
else
  print_info "Skipping tests as requested."
fi

# Build the system
print_info "Building the system..."
"$ROOT_DIR/cicd/scripts/build.sh"

if [ $? -ne 0 ]; then
  print_error "Build failed. Aborting deployment."
  exit 1
fi

# Deploy based on environment
if [ "$ENVIRONMENT" = "development" ]; then
  print_info "Deploying to development environment..."
  
  # Start services using docker-compose
  docker-compose -f docker-compose.dev.yml up -d
  
  if [ $? -eq 0 ]; then
    print_success "Deployed to development environment successfully"
  else
    print_error "Failed to deploy to development environment"
    exit 1
  fi
  
elif [ "$ENVIRONMENT" = "production" ]; then
  print_info "Deploying to production environment..."
  
  # Check if SSH key is available
  if [ ! -f "$HOME/.ssh/id_rsa" ]; then
    print_error "SSH key not found. Please set up SSH key for production deployment."
    exit 1
  fi
  
  # Deploy to production server
  print_info "Connecting to production server..."
  
  # Replace with actual production server details
  PROD_SERVER="user@production-server"
  PROD_DIR="/opt/local-ai-agent-system"
  
  # Copy files to production server
  rsync -avz --exclude 'node_modules' --exclude '.git' "$ROOT_DIR/" "$PROD_SERVER:$PROD_DIR/"
  
  # Start services on production server
  ssh "$PROD_SERVER" "cd $PROD_DIR && docker-compose -f docker-compose.prod.yml up -d"
  
  if [ $? -eq 0 ]; then
    print_success "Deployed to production environment successfully"
  else
    print_error "Failed to deploy to production environment"
    exit 1
  fi
  
else
  print_error "Unknown environment: $ENVIRONMENT"
  exit 1
fi

print_success "Deployment completed successfully"
