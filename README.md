# Local AI Agent System

A local AI agent system inspired by Manus.ai with Go microservices backend and React micro frontends.

## Overview

This system provides a local AI agent that can perform various tasks such as:

- Task management
- Code execution
- Web browsing
- File system operations

The system is built using a microservices architecture with Go for the backend services and React for the frontend applications.

## Architecture

### Backend Services

- **Task Service**: Manages tasks and their lifecycle
- **AI Service**: Handles AI model interactions
- **Code Execution Service**: Executes code in various languages
- **Web Browsing Service**: Provides web browsing capabilities
- **Filesystem Service**: Manages file system operations

### Frontend Applications

- **Shell**: The main application shell that hosts the micro frontends
- **Task Management App**: Manages tasks and their status
- **Result Viewer App**: Displays results from various operations

## Getting Started

### Prerequisites

- Go 1.21 or later
- Node.js 16 or later
- npm 8 or later
- Docker and Docker Compose (optional, for containerized deployment)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/local-ai-agent-system.git
cd local-ai-agent-system
```

2. Set up the development environment:

```bash
make setup
```

### Running the Frontend Applications

You can run the frontend applications in development mode using the following commands:

#### Option 1: Run all frontend applications at once

```bash
make run-frontend
```

This will start all three frontend applications:
- Shell application at http://localhost:3000
- Task Management application at http://localhost:5001
- Result Viewer application at http://localhost:5002

#### Option 2: Run each application individually

1. Run the Task Management application:

```bash
cd frontend/task-management-app
npm install
npm run dev
```

2. Run the Result Viewer application:

```bash
cd frontend/result-viewer-app
npm install
npm run dev
```

3. Run the Shell application:

```bash
cd frontend/shell
npm install
npm run dev
```

### Running the Backend Services

You can run the backend services using the following command:

```bash
make run
```

This will start all the backend services using Docker Compose.

## Development

### Project Structure

```
.
├── backend/
│   ├── task-service/
│   ├── ai-service/
│   ├── code-execution-service/
│   ├── web-browsing-service/
│   └── filesystem-service/
├── frontend/
│   ├── shell/
│   ├── task-management-app/
│   └── result-viewer-app/
├── cicd/
│   └── scripts/
├── docs/
└── docker-compose.yml
```

### Building

To build all components:

```bash
make build
```

### Testing

To run all tests:

```bash
make test
```

### Linting

To run linters on the codebase:

```bash
make lint
```

## Deployment

### Local Deployment

To deploy the system locally:

```bash
make run
```

### Production Deployment

For production deployment, use the production Docker Compose file:

```bash
docker-compose -f docker-compose.prod.yml up -d
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
