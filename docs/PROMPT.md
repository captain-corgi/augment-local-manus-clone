You are tasked with developing a local AI agent system, inspired by Manus.ai and agenticSeek, that runs entirely on the user’s hardware for privacy and cloud independence. The system must support autonomous task execution, including coding, web browsing, task planning, and filesystem interaction. The project must strictly follow Domain-Driven Design (DDD), use microservices with Clean Architecture for the backend, and implement a Micro Front End architecture for the frontend. All unit tests should employ Table-Driven Testing. Additionally, each functionality folder must include a README.md and a .windsurfrules file to guide AI assistants like Cursor or Windsurf.
Generate the entire project with the highest quality, adhering to the Software Development Lifecycle (SDLC) phases: planning, design, implementation, testing, deployment, and maintenance.
Project Structure
The project is organized into four main directories: backend, frontend, docs, and cicd. There is no root directory like ai-agent-system.

```shell
├── backend/
│   ├── task-service/
│   │   ├── README.md
│   │   ├── .windsurfrules
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── infrastructure/
│   │   └── delivery/
│   ├── ai-service/
│   │   ├── README.md
│   │   ├── .windsurfrules
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── infrastructure/
│   │   └── delivery/
│   ├── code-execution-service/
│   │   ├── README.md
│   │   ├── .windsurfrules
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── infrastructure/
│   │   └── delivery/
│   ├── web-browsing-service/
│   │   ├── README.md
│   │   ├── .windsurfrules
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── infrastructure/
│   │   └── delivery/
│   └── filesystem-service/
│       ├── README.md
│       ├── .windsurfrules
│       ├── go.mod
│       ├── main.go
│       ├── domain/
│       ├── usecase/
│       ├── infrastructure/
│       └── delivery/
├── frontend/
│   ├── shell/
│   │   ├── README.md
│   │   ├── .windsurfrules
│   │   ├── vite.config.ts
│   │   ├── src/
│   │   └── package.json
│   ├── task-management-app/
│   │   ├── README.md
│   │   ├── .windsurfrules
│   │   ├── vite.config.ts
│   │   ├── src/
│   │   └── package.json
│   └── result-viewer-app/
│       ├── README.md
│       ├── .windsurfrules
│       ├── vite.config.ts
│       ├── src/
│       └── package.json
├── docs/
│   ├── README.md
│   ├── .windsurfrules
│   ├── installation.md
│   ├── usage.md
│   └── architecture.md
└── cicd/
    ├── README.md
    ├── .windsurfrules
    ├── pipelines/
    └── scripts/
```

Backend (Golang)

Microservices: Implement the following services:
Task Service: Manages task lifecycle (submission, status, results).
AI Service: Interfaces with the local LLM (Ollama) for reasoning and planning.
Code Execution Service: Executes code in sandboxed Docker containers.
Web Browsing Service: Navigates web pages using chromedp.
Filesystem Service: Performs secure file operations within /app/workspace.


Clean Architecture: Each service should have:
Entities: Core domain models (e.g., Task).
Use Cases: Business logic (e.g., CreateTaskUseCase).
Interface Adapters: Repositories and controllers (e.g., SQLiteTaskRepository).
Frameworks/Drivers: External systems (e.g., Gin for APIs, SQLite for storage).


Domain-Driven Design (DDD):
Bounded Contexts: Task Management, AI Reasoning, Tool Execution (Code Execution, Web Browsing, Filesystem Operations).
Entities and Aggregates: Task as an aggregate root in Task Management.
Repositories: Interfaces for data access, implemented in infrastructure/.


Communication: Services use HTTP/JSON APIs.
Security: Sandbox code execution with Docker and restrict filesystem access.

Frontend (ReactJS with Vite, TypeScript, TailwindCSS)

Micro Front End Architecture:
Shell Application: Hosts micro frontends using Vite’s Module Federation plugin.
Task Management App: Handles task submission and listing.
Result Viewer App: Displays task results.


Integration: Apps communicate with backend services via REST APIs, using polling for task status updates.
Styling: Use TailwindCSS for responsive design.

AI Integration

Local LLM: Use Ollama with a model like DeepSeek-R1.
Integration: AI Service communicates with Ollama’s REST API to process tasks.

Testing

Unit Tests: Use Table-Driven Testing in Go for backend services.
Frontend Tests: Use Jest and React Testing Library, applying table-driven principles where possible.
Integration Tests: Validate microservice and frontend-backend interactions.

Documentation and CI/CD

docs/: Contains README.md, .windsurfrules, and guides for installation, usage, and architecture.
cicd/: Includes README.md, .windsurfrules, pipeline configurations, and automation scripts.

Implementation Instructions

Set Up the Project Structure:

Create the backend, frontend, docs, and cicd directories.
Inside backend, create folders for each microservice (task-service, ai-service, etc.), each containing README.md, .windsurfrules, go.mod, main.go, and subdirectories for domain/, usecase/, infrastructure/, and delivery/.
Inside frontend, create folders for shell, task-management-app, and result-viewer-app, each with README.md, .windsurfrules, vite.config.ts, src/, and package.json.


Backend Development:

For each microservice, initialize a Go module (e.g., go mod init github.com/yourusername/backend/task-service).
Implement Clean Architecture layers:
Define entities in domain/.
Implement use cases in usecase/.
Create repositories in infrastructure/.
Set up API handlers in delivery/http/.


Use Gin for API routing and SQLite for data storage in the Task Service.
Integrate Ollama in the AI Service using github.com/jmorganca/ollama.
Use the Docker Go client for the Code Execution Service.
Implement chromedp for the Web Browsing Service.
Use Golang’s os package for the Filesystem Service, ensuring operations are restricted to /app/workspace.


Frontend Development:

Set up Vite projects for the shell and apps with React, TypeScript, and TailwindCSS.
Configure Module Federation in vite.config.ts to load apps dynamically in the shell.
Implement components for task submission, listing, and result viewing.
Use fetch or Axios to communicate with backend APIs.


AI Integration:

Install Ollama and download the DeepSeek-R1 model.
Start the Ollama server and ensure the AI Service can communicate with it.


Testing:

Write unit tests for backend services using Go’s testing package with table-driven tests.
Test frontend components with Jest and React Testing Library.
Perform integration tests to validate end-to-end functionality.


Security:

Ensure code execution is sandboxed in Docker containers.
Restrict filesystem operations to /app/workspace.
Implement input validation on APIs and frontend.


Documentation:

Write README.md for each folder, explaining its purpose and usage.
Create .windsurfrules for each folder to guide AI assistants, tailored to the specific technology stack and constraints.


CI/CD:

Set up pipeline configurations in cicd/pipelines/ for automated testing and deployment.
Write automation scripts in cicd/scripts/ for build and deployment processes.



Example README.md (for backend/task-service/)
# Task Service
This microservice manages the lifecycle of tasks, including submission, status tracking, and result storage. It uses SQLite for persistence and communicates with other services via HTTP/JSON APIs.

## Setup
1. Install Go: `go version >= 1.21`
2. Initialize: `go mod tidy`
3. Run: `go run main.go`

## Usage
- Submit a task: `POST /tasks`
- Check status: `GET /tasks/{id}`

Example .windsurfrules (for backend/task-service/)
You are Windsurf Cascade, an AI assistant with advanced problem-solving capabilities. Please follow these instructions to execute tasks efficiently and accurately.

## Core Operating Principles
1. **Instruction Reception and Understanding**
   - Carefully read and interpret user instructions
   - Ask specific questions when clarification is needed
   - Clearly identify technical constraints and requirements
   - Do not perform any operations beyond what is instructed

2. **In-depth Analysis and Planning**
   ## Task Analysis
   - Purpose: [Final goal of the task]
   - Technical Requirements: [Technology stack and constraints]
   - Implementation Steps: [Specific steps]
   - Risks: [Potential issues]
   - Quality Standards: [Requirements to meet]

3. **Implementation Planning**
   ## Implementation Plan
   1. [Specific step 1]
      - Detailed implementation content
      - Expected challenges and countermeasures
   2. [Specific step 2]
      ...

4. **Comprehensive Implementation and Verification**
   - Execute file operations and related processes in optimized complete sequences
   - Continuously verify against quality standards throughout implementation
   - Address issues promptly with integrated solutions
   - Execute processes only within the scope of instructions, without adding extra features or operations

5. **Continuous Feedback**
   - Regularly report implementation progress
   - Confirm at critical decision points
   - Promptly report issues with proposed solutions

## Technology Stack and Constraints
### Core Technologies
- Go: ^1.21
- SQLite: ^3.0.0
- Gin: ^1.9.0

## Quality Management Protocol
### 1. Code Quality
- Adhere to Go linting standards
- Use Clean Architecture and DDD
- Maintain consistency in naming and structure
### 2. Performance
- Optimize database queries
- Efficient HTTP handling
### 3. Security
- Validate all inputs
- Use parameterized queries for SQLite
### 4. Testing
- Implement table-driven unit tests
- Ensure full coverage of use cases

## Project Structure Convention

task-service/├── domain/          # Core entities and models├── usecase/         # Business logic├── infrastructure/  # Repositories and external integrations└── delivery/        # API handlers and controllers

## Important Constraints
1. **Restricted Operations**
   - No direct filesystem access outside SQLite
   - HTTP/JSON communication only
2. **Version Management**
   - Technology stack versions are fixed unless approved
3. **Code Placement**
   - Entities in `domain/`
   - Business logic in `usecase/`
   - API routes in `delivery/http/`

## Implementation Process
### 1. Initial Analysis Phase
   ### Requirements Analysis
   - Identify functional requirements
   - Confirm technical constraints
   - Check consistency with existing code
   ### Risk Assessment
   - Potential technical challenges
   - Performance impacts
   - Security risks
### 2. Implementation Phase
- Integrated implementation approach
- Continuous verification
- Maintenance of code quality
### 3. Verification Phase
- Unit testing
- Integration testing
- Performance testing
### 4. Final Confirmation
- Consistency with requirements
- Code quality
- Documentation completeness

## Error Handling Protocol
1. **Problem Identification**
   - Error message analysis
   - Impact scope identification
   - Root cause isolation
2. **Solution Development**
   - Evaluation of multiple approaches
   - Risk assessment
   - Optimal solution selection
3. **Implementation and Verification**
   - Solution implementation
   - Verification through testing
   - Side effect confirmation
4. **Documentation**
   - Record of problem and solution
   - Preventive measure proposals
   - Sharing of learning points

I will follow these instructions to deliver high-quality implementations. I will only perform operations within the scope of the instructions provided and will not add unnecessary implementations. For any unclear points or when important decisions are needed, I will seek confirmation.

Final Instructions

Ensure all components are implemented according to the specified architectural patterns (DDD, Clean Architecture, Micro Front Ends).
Prioritize security, especially in code execution and filesystem operations.
Follow the SDLC phases to maintain a structured development process.
Use Table-Driven Testing to ensure comprehensive unit test coverage.
Document the system thoroughly for ease of installation and use, including README.md and 2 rule files with same content: '.windsurfrules', '.augment-guidelines' in each folder.
Generate the full project, including all source code, configuration files, and documentation, with the highest quality and attention to detail.

