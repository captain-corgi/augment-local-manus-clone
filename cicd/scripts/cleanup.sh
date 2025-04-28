#!/bin/bash

# Cleanup script for Local AI Agent System
# This script cleans up old Docker images, containers, and volumes

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
PRUNE_IMAGES=false
PRUNE_VOLUMES=false
REMOVE_BACKUPS=false
BACKUP_RETENTION_DAYS=30

while [[ $# -gt 0 ]]; do
  case $1 in
    --prune-images|-i)
      PRUNE_IMAGES=true
      shift
      ;;
    --prune-volumes|-v)
      PRUNE_VOLUMES=true
      shift
      ;;
    --remove-backups|-b)
      REMOVE_BACKUPS=true
      shift
      ;;
    --backup-retention|-r)
      BACKUP_RETENTION_DAYS="$2"
      shift 2
      ;;
    --all|-a)
      PRUNE_IMAGES=true
      PRUNE_VOLUMES=true
      REMOVE_BACKUPS=true
      shift
      ;;
    --help|-h)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  --prune-images, -i       Prune unused Docker images"
      echo "  --prune-volumes, -v      Prune unused Docker volumes"
      echo "  --remove-backups, -b     Remove old backups"
      echo "  --backup-retention, -r N Retain backups for N days (default: 30)"
      echo "  --all, -a                Perform all cleanup operations"
      echo "  --help, -h               Show this help message"
      exit 0
      ;;
    *)
      print_error "Unknown option: $1"
      exit 1
      ;;
  esac
done

print_info "Cleaning up Local AI Agent System..."

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
  print_error "Docker is not installed. Please install Docker and try again."
  exit 1
fi

# Remove stopped containers
print_info "Removing stopped containers..."
docker container prune -f
print_success "Stopped containers removed"

# Prune unused Docker images
if [ "$PRUNE_IMAGES" = true ]; then
  print_info "Pruning unused Docker images..."
  docker image prune -a -f
  print_success "Unused Docker images pruned"
fi

# Prune unused Docker volumes
if [ "$PRUNE_VOLUMES" = true ]; then
  print_info "Pruning unused Docker volumes..."
  docker volume prune -f
  print_success "Unused Docker volumes pruned"
fi

# Remove old backups
if [ "$REMOVE_BACKUPS" = true ]; then
  print_info "Removing backups older than $BACKUP_RETENTION_DAYS days..."
  
  # Check if backup directory exists
  BACKUP_DIR="$ROOT_DIR/backups"
  if [ -d "$BACKUP_DIR" ]; then
    # Find and remove old backup files
    find "$BACKUP_DIR" -name "backup_*.tar.gz" -type f -mtime +$BACKUP_RETENTION_DAYS -delete -print | while read -r file; do
      print_info "Removed old backup: $file"
    done
    
    print_success "Old backups removed"
  else
    print_warning "Backup directory not found: $BACKUP_DIR"
  fi
fi

# Clean up temporary files
print_info "Cleaning up temporary files..."
find "$ROOT_DIR" -name "*.tmp" -type f -delete
find "$ROOT_DIR" -name "*.log" -type f -size +10M -delete
find "$ROOT_DIR/tmp" -type f -mtime +7 -delete 2>/dev/null || true

print_success "Cleanup completed successfully"

# Show disk usage
print_info "Current disk usage:"
df -h | grep -E "Filesystem|/$"
