# Architecture Overview

This document provides an overview of the Local AI Agent System architecture.

## System Architecture

The Local AI Agent System follows a microservices architecture for the backend and a micro frontend architecture for the frontend. It strictly adheres to Domain-Driven Design (DDD) principles and Clean Architecture.

### High-Level Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  Frontend Apps  │◄───►│ Backend Services│◄───►│  External Tools │
│                 │     │                 │     │                 │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

## Backend Architecture

The backend consists of five microservices, each responsible for a specific domain:

### Task Service

Manages the lifecycle of tasks, including submission, status tracking, and result storage.

- **Domain**: Task entities and repositories
- **Use Cases**: CreateTask, GetTask, ListTasks, UpdateTask, DeleteTask
- **Infrastructure**: SQLite repository
- **Delivery**: HTTP API with Gin

### AI Service

Interfaces with the local LLM (Ollama) for reasoning and planning.

- **Domain**: AIRequest, AIResponse
- **Use Cases**: ProcessAIRequest
- **Infrastructure**: Ollama client
- **Delivery**: HTTP API with Gin

### Code Execution Service

Executes code in sandboxed Docker containers.

- **Domain**: CodeExecution
- **Use Cases**: ExecuteCode
- **Infrastructure**: Docker client
- **Delivery**: HTTP API with Gin

### Web Browsing Service

Navigates web pages using chromedp.

- **Domain**: WebBrowsingRequest, WebBrowsingResult
- **Use Cases**: BrowseWeb, SearchWeb, InteractWithWeb
- **Infrastructure**: chromedp client
- **Delivery**: HTTP API with Gin

### Filesystem Service

Performs secure file operations within a restricted workspace directory.

- **Domain**: FileOperation, FileInfo
- **Use Cases**: ReadFile, WriteFile, ListFiles, DeleteFile, MakeDirectory
- **Infrastructure**: OS file operations
- **Delivery**: HTTP API with Gin

## Frontend Architecture

The frontend follows a Micro Frontend architecture using Vite's Module Federation:

### Shell Application

Hosts the micro frontends and provides the main layout and navigation.

- **Components**: MainLayout, Navbar, Sidebar
- **Routing**: React Router
- **Styling**: TailwindCSS

### Task Management App

Handles task submission and listing.

- **Components**: TaskForm, TaskList
- **API Integration**: Task Service
- **State Management**: React Hooks

### Result Viewer App

Displays task results in a user-friendly format.

- **Components**: TaskResultList, TaskResultDetail
- **API Integration**: Task Service
- **Features**: Syntax highlighting for code

## Clean Architecture

Each backend service follows Clean Architecture with the following layers:

1. **Domain Layer**: Contains the core business entities and interfaces
2. **Use Case Layer**: Implements the business logic
3. **Infrastructure Layer**: Provides implementations for repositories and external services
4. **Delivery Layer**: Handles HTTP requests and responses

```
┌─────────────────────────────────────────┐
│                                         │
│  Delivery Layer (HTTP, CLI, etc.)       │
│                                         │
├─────────────────────────────────────────┤
│                                         │
│  Use Case Layer (Application Logic)     │
│                                         │
├─────────────────────────────────────────┤
│                                         │
│  Domain Layer (Entities, Repositories)  │
│                                         │
├─────────────────────────────────────────┤
│                                         │
│  Infrastructure Layer (DB, External)    │
│                                         │
└─────────────────────────────────────────┘
```

## Domain-Driven Design (DDD)

The system follows DDD principles:

- **Bounded Contexts**: Each microservice represents a bounded context
- **Entities**: Core domain objects with identity and lifecycle
- **Value Objects**: Immutable objects without identity
- **Aggregates**: Clusters of entities and value objects
- **Repositories**: Interfaces for data access
- **Services**: Domain operations that don't belong to entities

## Communication

Services communicate with each other via HTTP/JSON APIs. The frontend communicates with the backend using the same APIs.

## Security

The system implements several security measures:

- **Code Execution**: Sandboxed in Docker containers with limited resources and no network access
- **Filesystem Operations**: Restricted to a workspace directory with path traversal prevention
- **Web Browsing**: Controlled browser automation with restricted capabilities
- **Input Validation**: All inputs are validated before processing

## Testing

All components include comprehensive tests using Table-Driven Testing:

- **Backend**: Go's testing package with table-driven tests
- **Frontend**: Jest and React Testing Library

## Deployment

The system can be deployed using Docker Compose, which starts all services in separate containers.
