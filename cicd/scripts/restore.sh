#!/bin/bash

# Restore script for Local AI Agent System
# This script restores the system from a backup

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
BACKUP_DIR="$ROOT_DIR/backups"
BACKUP_FILE=""
RESTORE_WORKSPACE=true
RESTORE_DATABASE=true

while [[ $# -gt 0 ]]; do
  case $1 in
    --backup-dir|-b)
      BACKUP_DIR="$2"
      shift 2
      ;;
    --backup-file|-f)
      BACKUP_FILE="$2"
      shift 2
      ;;
    --no-workspace)
      RESTORE_WORKSPACE=false
      shift
      ;;
    --no-database)
      RESTORE_DATABASE=false
      shift
      ;;
    --help|-h)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  --backup-dir, -b <dir>   Backup directory (default: ./backups)"
      echo "  --backup-file, -f <file> Specific backup file to restore"
      echo "  --no-workspace           Exclude workspace directory from restore"
      echo "  --no-database            Exclude database files from restore"
      echo "  --help, -h               Show this help message"
      exit 0
      ;;
    *)
      print_error "Unknown option: $1"
      exit 1
      ;;
  esac
done

# Check if backup directory exists
if [ ! -d "$BACKUP_DIR" ]; then
  print_error "Backup directory not found: $BACKUP_DIR"
  exit 1
fi

# If no specific backup file is provided, list available backups and ask user to select one
if [ -z "$BACKUP_FILE" ]; then
  print_info "Available backups:"
  
  # Get list of backup files
  BACKUP_FILES=($(ls -t "$BACKUP_DIR" | grep "^backup_.*\.tar\.gz$"))
  
  if [ ${#BACKUP_FILES[@]} -eq 0 ]; then
    print_error "No backup files found in $BACKUP_DIR"
    exit 1
  fi
  
  # Display backup files with numbers
  for i in "${!BACKUP_FILES[@]}"; do
    backup_size=$(du -h "$BACKUP_DIR/${BACKUP_FILES[$i]}" | cut -f1)
    backup_date=$(echo "${BACKUP_FILES[$i]}" | sed 's/backup_\([0-9]\{8\}_[0-9]\{6\}\).*/\1/')
    backup_date_formatted=$(date -d "${backup_date:0:8} ${backup_date:9:2}:${backup_date:11:2}:${backup_date:13:2}" "+%Y-%m-%d %H:%M:%S" 2>/dev/null || echo "${backup_date:0:4}-${backup_date:4:2}-${backup_date:6:2} ${backup_date:9:2}:${backup_date:11:2}:${backup_date:13:2}")
    echo "  $((i+1)). ${BACKUP_FILES[$i]} ($backup_size) - $backup_date_formatted"
  done
  
  # Ask user to select a backup
  read -p "Select a backup to restore (1-${#BACKUP_FILES[@]}): " selection
  
  # Validate selection
  if ! [[ "$selection" =~ ^[0-9]+$ ]] || [ "$selection" -lt 1 ] || [ "$selection" -gt ${#BACKUP_FILES[@]} ]; then
    print_error "Invalid selection: $selection"
    exit 1
  fi
  
  # Set backup file
  BACKUP_FILE="$BACKUP_DIR/${BACKUP_FILES[$((selection-1))]}"
else
  # Check if the provided backup file exists
  if [ ! -f "$BACKUP_FILE" ]; then
    # Check if it's a relative path in the backup directory
    if [ -f "$BACKUP_DIR/$BACKUP_FILE" ]; then
      BACKUP_FILE="$BACKUP_DIR/$BACKUP_FILE"
    else
      print_error "Backup file not found: $BACKUP_FILE"
      exit 1
    fi
  fi
fi

print_info "Restoring from backup: $BACKUP_FILE"

# Confirm restore
read -p "This will overwrite existing data. Are you sure you want to continue? (y/n): " confirm
if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
  print_info "Restore cancelled"
  exit 0
fi

# Create temporary directory for restore
TEMP_DIR=$(mktemp -d)

# Extract backup archive
print_info "Extracting backup archive..."
tar -xzf "$BACKUP_FILE" -C "$TEMP_DIR"

# Restore database files
if [ "$RESTORE_DATABASE" = true ]; then
  print_info "Restoring database files..."
  
  # Check if data directory exists in the backup
  if [ -d "$TEMP_DIR/data" ]; then
    # Stop services to ensure data consistency
    print_info "Stopping services to ensure data consistency..."
    if command -v docker-compose &> /dev/null && [ -f "$ROOT_DIR/docker-compose.yml" ]; then
      docker-compose stop task-service
    fi
    
    # Create data directory if it doesn't exist
    mkdir -p "$ROOT_DIR/data"
    
    # Copy database files
    cp -r "$TEMP_DIR/data/"* "$ROOT_DIR/data/"
    
    # Restart services
    print_info "Restarting services..."
    if command -v docker-compose &> /dev/null && [ -f "$ROOT_DIR/docker-compose.yml" ]; then
      docker-compose start task-service
    fi
    
    print_success "Database files restored successfully"
  else
    print_warning "Data directory not found in backup, skipping database restore"
  fi
fi

# Restore workspace directory
if [ "$RESTORE_WORKSPACE" = true ]; then
  print_info "Restoring workspace directory..."
  
  # Check if workspace directory exists in the backup
  if [ -d "$TEMP_DIR/workspace" ]; then
    # Create workspace directory if it doesn't exist
    mkdir -p "$ROOT_DIR/workspace"
    
    # Copy workspace files
    cp -r "$TEMP_DIR/workspace/"* "$ROOT_DIR/workspace/"
    
    print_success "Workspace directory restored successfully"
  else
    print_warning "Workspace directory not found in backup, skipping workspace restore"
  fi
fi

# Clean up temporary directory
rm -rf "$TEMP_DIR"

print_success "Restore completed successfully"
