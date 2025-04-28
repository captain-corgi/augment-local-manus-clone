# Installation Guide

This guide provides instructions for setting up the Local AI Agent System on your machine.

## Prerequisites

Before installing the system, ensure you have the following prerequisites:

- **Go**: Version 1.21 or later
- **Node.js**: Version 16 or later
- **npm**: Version 8 or later
- **Docker**: Latest version
- **Chrome/Chromium**: Required for web browsing service
- **Ollama**: Required for AI service

## Installation Steps

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/local-ai-agent-system.git
cd local-ai-agent-system
```

### 2. Install Ollama and Download the Model

Follow the instructions at [https://ollama.ai/](https://ollama.ai/) to install Ollama, then download the DeepSeek-R1 model:

```bash
ollama pull deepseek-r1
```

### 3. Set Up Backend Services

#### Task Service

```bash
cd backend/task-service
go mod tidy
go run main.go
```

#### AI Service

```bash
cd backend/ai-service
go mod tidy
go run main.go
```

#### Code Execution Service

```bash
cd backend/code-execution-service
go mod tidy
go run main.go
```

#### Web Browsing Service

```bash
cd backend/web-browsing-service
go mod tidy
go run main.go
```

#### Filesystem Service

```bash
cd backend/filesystem-service
go mod tidy
go run main.go
```

### 4. Set Up Frontend Applications

#### Shell Application

```bash
cd frontend/shell
npm install
npm run dev
```

#### Task Management App

```bash
cd frontend/task-management-app
npm install
npm run dev
```

#### Result Viewer App

```bash
cd frontend/result-viewer-app
npm install
npm run dev
```

## Configuration

### Backend Services

Each backend service can be configured by modifying its main.go file. The default configuration is:

- **Task Service**: Port 8081
- **AI Service**: Port 8082
- **Code Execution Service**: Port 8083
- **Web Browsing Service**: Port 8084
- **Filesystem Service**: Port 8085

### Frontend Applications

Frontend applications can be configured by modifying their vite.config.ts files. The default configuration is:

- **Shell**: Port 5000
- **Task Management App**: Port 5001
- **Result Viewer App**: Port 5002

### Workspace Directory

The Filesystem Service uses a workspace directory for file operations. By default, it uses:

- `/app/workspace` in production
- `./workspace` in development

## Running with Docker

A Docker Compose file is provided to run the entire system with a single command:

```bash
docker-compose up
```

This will start all backend services and frontend applications in separate containers.

## Troubleshooting

### Common Issues

- **Port conflicts**: If a port is already in use, change the port in the service's configuration.
- **Ollama connection issues**: Ensure Ollama is running and the DeepSeek-R1 model is downloaded.
- **Docker permission issues**: Ensure your user has permission to use Docker.

### Logs

Each service writes logs to stdout. You can redirect these to files if needed:

```bash
go run main.go > service.log 2>&1
```
