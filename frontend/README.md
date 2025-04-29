# Frontend Applications

This directory contains the frontend applications for the Local AI Agent System.

## Overview

The frontend is built using a micro frontend architecture with the following applications:

- **Shell**: The main application shell that hosts the micro frontends
- **Task Management App**: Manages tasks and their status
- **Result Viewer App**: Displays results from various operations

## Architecture

The frontend uses Module Federation to load micro frontends at runtime. This allows for independent development and deployment of each application.

### Shell Application

The shell application is the main entry point for users. It provides:

- Navigation
- Layout
- Authentication
- Routing
- Error handling

### Task Management App

The Task Management app is responsible for:

- Creating tasks
- Updating task status
- Viewing task details
- Filtering tasks

### Result Viewer App

The Result Viewer app is responsible for:

- Displaying results from tasks
- Visualizing data
- Exporting results

## Getting Started

### Prerequisites

- Node.js 16 or later
- npm 8 or later

### Installation

1. Install dependencies for all applications:

```bash
cd shell && npm install
cd ../task-management-app && npm install
cd ../result-viewer-app && npm install
```

### Running the Applications

#### Option 1: Run all applications at once

From the project root:

```bash
make run-frontend
```

#### Option 2: Run each application individually

1. Run the Task Management application:

```bash
cd task-management-app
npm run dev
```

2. Run the Result Viewer application:

```bash
cd result-viewer-app
npm run dev
```

3. Run the Shell application:

```bash
cd shell
npm run dev
```

### Access the Applications

- Shell: http://localhost:3000
- Task Management: http://localhost:5001
- Result Viewer: http://localhost:5002

## Development

### Project Structure

```
frontend/
├── shell/                 # Main application shell
│   ├── src/
│   │   ├── components/    # Reusable components
│   │   ├── layouts/       # Page layouts
│   │   ├── types/         # TypeScript type definitions
│   │   └── App.tsx        # Main application component
│   ├── package.json
│   └── vite.config.ts     # Vite configuration with Module Federation
├── task-management-app/   # Task Management micro frontend
│   ├── src/
│   │   ├── components/    # Task-specific components
│   │   └── App.tsx        # Task Management application component
│   ├── package.json
│   └── vite.config.ts     # Vite configuration with Module Federation
└── result-viewer-app/     # Result Viewer micro frontend
    ├── src/
    │   ├── components/    # Result-specific components
    │   └── App.tsx        # Result Viewer application component
    ├── package.json
    └── vite.config.ts     # Vite configuration with Module Federation
```

### Module Federation Configuration

The applications use Module Federation to share components and libraries. The configuration is in the `vite.config.ts` files:

#### Shell Application

```typescript
federation({
  name: 'shell',
  remotes: {
    taskManagement: 'http://localhost:5001/remoteEntry.js',
    resultViewer: 'http://localhost:5002/remoteEntry.js',
  },
  shared: ['react', 'react-dom', 'react-router-dom'],
})
```

#### Task Management App

```typescript
federation({
  name: 'taskManagement',
  filename: 'remoteEntry.js',
  exposes: {
    './App': './src/App.tsx',
  },
  shared: ['react', 'react-dom', 'react-router-dom'],
})
```

#### Result Viewer App

```typescript
federation({
  name: 'resultViewer',
  filename: 'remoteEntry.js',
  exposes: {
    './App': './src/App.tsx',
  },
  shared: ['react', 'react-dom', 'react-router-dom'],
})
```

## Building for Production

To build all frontend applications for production:

```bash
cd shell && npm run build
cd ../task-management-app && npm run build
cd ../result-viewer-app && npm run build
```

## Troubleshooting

### Common Issues

1. **Module Federation Loading Errors**:
   - Make sure all applications are running
   - Check that the remoteEntry.js files are accessible
   - Verify that the shared dependencies have compatible versions

2. **CORS Issues**:
   - Ensure that CORS headers are properly set in the development server
   - Check that the applications are running on the expected ports

3. **React Version Mismatch**:
   - Make sure all applications use the same version of React
   - Configure shared dependencies correctly in Module Federation
