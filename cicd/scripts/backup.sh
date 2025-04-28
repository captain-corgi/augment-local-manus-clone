#!/bin/bash

# Backup script for Local AI Agent System
# This script creates a backup of the system data

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
BACKUP_DIR="$ROOT_DIR/backups"
INCLUDE_WORKSPACE=true
INCLUDE_DATABASE=true

while [[ $# -gt 0 ]]; do
  case $1 in
    --backup-dir|-b)
      BACKUP_DIR="$2"
      shift 2
      ;;
    --no-workspace)
      INCLUDE_WORKSPACE=false
      shift
      ;;
    --no-database)
      INCLUDE_DATABASE=false
      shift
      ;;
    --help|-h)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  --backup-dir, -b <dir>  Backup directory (default: ./backups)"
      echo "  --no-workspace          Exclude workspace directory from backup"
      echo "  --no-database           Exclude database files from backup"
      echo "  --help, -h              Show this help message"
      exit 0
      ;;
    *)
      print_error "Unknown option: $1"
      exit 1
      ;;
  esac
done

# Create backup directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

# Generate timestamp for backup filename
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="$BACKUP_DIR/backup_$TIMESTAMP.tar.gz"

print_info "Creating backup at $BACKUP_FILE..."

# Create temporary directory for backup
TEMP_DIR=$(mktemp -d)
mkdir -p "$TEMP_DIR/data"
mkdir -p "$TEMP_DIR/workspace"

# Backup database files
if [ "$INCLUDE_DATABASE" = true ]; then
  print_info "Backing up database files..."
  
  # Check if data directory exists
  if [ -d "$ROOT_DIR/data" ]; then
    # Stop services to ensure data consistency
    print_info "Stopping services to ensure data consistency..."
    if command -v docker-compose &> /dev/null && [ -f "$ROOT_DIR/docker-compose.yml" ]; then
      docker-compose stop task-service
    fi
    
    # Copy database files
    cp -r "$ROOT_DIR/data" "$TEMP_DIR/"
    
    # Restart services
    print_info "Restarting services..."
    if command -v docker-compose &> /dev/null && [ -f "$ROOT_DIR/docker-compose.yml" ]; then
      docker-compose start task-service
    fi
    
    print_success "Database files backed up successfully"
  else
    print_warning "Data directory not found, skipping database backup"
  fi
fi

# Backup workspace directory
if [ "$INCLUDE_WORKSPACE" = true ]; then
  print_info "Backing up workspace directory..."
  
  # Check if workspace directory exists
  if [ -d "$ROOT_DIR/workspace" ]; then
    cp -r "$ROOT_DIR/workspace" "$TEMP_DIR/"
    print_success "Workspace directory backed up successfully"
  else
    print_warning "Workspace directory not found, skipping workspace backup"
  fi
fi

# Create backup archive
print_info "Creating backup archive..."
tar -czf "$BACKUP_FILE" -C "$TEMP_DIR" .

# Clean up temporary directory
rm -rf "$TEMP_DIR"

print_success "Backup created successfully: $BACKUP_FILE"
print_info "Backup size: $(du -h "$BACKUP_FILE" | cut -f1)"

# List all backups
print_info "Available backups:"
ls -lh "$BACKUP_DIR" | grep -v "^total" | awk '{print $9 " (" $5 ")"}'
