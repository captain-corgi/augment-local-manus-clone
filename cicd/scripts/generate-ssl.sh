#!/bin/bash

# Generate SSL certificates for development
# This script generates self-signed SSL certificates for local development

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

# Create SSL directory if it doesn't exist
mkdir -p "$ROOT_DIR/nginx/ssl"

print_info "Generating SSL certificates for development..."

# Check if OpenSSL is installed
if ! command -v openssl &> /dev/null; then
  print_error "OpenSSL is not installed. Please install OpenSSL and try again."
  exit 1
fi

# Generate SSL certificates
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout "$ROOT_DIR/nginx/ssl/server.key" \
  -out "$ROOT_DIR/nginx/ssl/server.crt" \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost" \
  -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

if [ $? -eq 0 ]; then
  print_success "SSL certificates generated successfully"
  print_info "Certificate: $ROOT_DIR/nginx/ssl/server.crt"
  print_info "Key: $ROOT_DIR/nginx/ssl/server.key"
  
  # Update Nginx configuration to use SSL
  print_info "Updating Nginx configuration to use SSL..."
  
  # Uncomment SSL configuration in default.conf
  sed -i.bak 's/# Uncomment the following lines when SSL is configured/# SSL is configured/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/# return 301 https:\/\/$host$request_uri;/return 301 https:\/\/$host$request_uri;/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/# server {/server {/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     listen 443 ssl http2;/    listen 443 ssl http2;/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     server_name localhost;/    server_name localhost;/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_certificate/    ssl_certificate/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_certificate_key/    ssl_certificate_key/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_protocols/    ssl_protocols/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_ciphers/    ssl_ciphers/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_prefer_server_ciphers/    ssl_prefer_server_ciphers/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_session_cache/    ssl_session_cache/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     ssl_session_timeout/    ssl_session_timeout/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#     location/    location/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#         proxy_pass/        proxy_pass/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/#         proxy_set_header/        proxy_set_header/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  sed -i.bak 's/# }/}/g' "$ROOT_DIR/nginx/conf.d/default.conf"
  
  # Remove backup file
  rm "$ROOT_DIR/nginx/conf.d/default.conf.bak"
  
  print_success "Nginx configuration updated successfully"
  print_info "You can now access the application at https://localhost"
else
  print_error "Failed to generate SSL certificates"
  exit 1
fi
