#!/bin/bash

# Update script for Local AI Agent System
# This script updates the system to the latest version

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

# Parse command line arguments
SKIP_BACKUP=false
SKIP_TESTS=false
FORCE_UPDATE=false

while [[ $# -gt 0 ]]; do
  case $1 in
    --skip-backup)
      SKIP_BACKUP=true
      shift
      ;;
    --skip-tests)
      SKIP_TESTS=true
      shift
      ;;
    --force)
      FORCE_UPDATE=true
      shift
      ;;
    --help|-h)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  --skip-backup    Skip creating a backup before updating"
      echo "  --skip-tests     Skip running tests after updating"
      echo "  --force          Force update even if there are uncommitted changes"
      echo "  --help, -h       Show this help message"
      exit 0
      ;;
    *)
      print_error "Unknown option: $1"
      exit 1
      ;;
  esac
done

print_info "Updating Local AI Agent System..."

# Check if Git is installed
if ! command -v git &> /dev/null; then
  print_error "Git is not installed. Please install Git and try again."
  exit 1
fi

# Check if the directory is a Git repository
if [ ! -d "$ROOT_DIR/.git" ]; then
  print_error "Not a Git repository. Please clone the repository and try again."
  exit 1
fi

# Check for uncommitted changes
if [ "$(git -C "$ROOT_DIR" status --porcelain)" != "" ]; then
  if [ "$FORCE_UPDATE" = true ]; then
    print_warning "Uncommitted changes detected, but --force flag is set. Proceeding with update."
  else
    print_error "Uncommitted changes detected. Please commit or stash your changes before updating."
    print_info "You can use --force to update anyway, but you may lose your changes."
    exit 1
  fi
fi

# Create a backup before updating
if [ "$SKIP_BACKUP" = false ]; then
  print_info "Creating backup before updating..."
  "$ROOT_DIR/cicd/scripts/backup.sh"
  
  if [ $? -ne 0 ]; then
    print_error "Backup failed. Aborting update."
    exit 1
  fi
fi

# Get current version
CURRENT_VERSION=$(git -C "$ROOT_DIR" rev-parse HEAD)
print_info "Current version: $CURRENT_VERSION"

# Fetch latest changes
print_info "Fetching latest changes..."
git -C "$ROOT_DIR" fetch origin

# Get latest version
LATEST_VERSION=$(git -C "$ROOT_DIR" rev-parse origin/main)
print_info "Latest version: $LATEST_VERSION"

# Check if already up to date
if [ "$CURRENT_VERSION" = "$LATEST_VERSION" ]; then
  print_success "Already up to date"
  exit 0
fi

# Update to latest version
print_info "Updating to latest version..."
git -C "$ROOT_DIR" pull origin main

# Run setup script
print_info "Running setup script..."
"$ROOT_DIR/cicd/scripts/setup-dev.sh"

# Build the system
print_info "Building the system..."
"$ROOT_DIR/cicd/scripts/build.sh"

# Run tests
if [ "$SKIP_TESTS" = false ]; then
  print_info "Running tests..."
  "$ROOT_DIR/cicd/scripts/test.sh"
  
  if [ $? -ne 0 ]; then
    print_error "Tests failed. Please check the logs for details."
    print_info "You may need to restore from the backup created before updating."
    exit 1
  fi
fi

# Restart services
print_info "Restarting services..."
if command -v docker-compose &> /dev/null && [ -f "$ROOT_DIR/docker-compose.yml" ]; then
  docker-compose down
  docker-compose up -d
  
  if [ $? -ne 0 ]; then
    print_error "Failed to restart services. Please check the logs for details."
    exit 1
  fi
else
  print_warning "Docker Compose not found or docker-compose.yml not found. Please restart services manually."
fi

print_success "Update completed successfully"
print_info "Updated from $CURRENT_VERSION to $LATEST_VERSION"
