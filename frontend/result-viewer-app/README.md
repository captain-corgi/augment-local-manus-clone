# Result Viewer App

This is the Result Viewer micro frontend for the local AI agent system. It displays task results in a user-friendly format.

## Architecture

The Result Viewer App follows the Micro Frontend architecture:

- It is exposed as a remote module using Vite's Module Federation plugin
- It is loaded by the Shell Application
- It communicates with the Task Service API

## Setup

1. Install dependencies:
```bash
npm install
```

2. Run the development server:
```bash
npm run dev
```

3. Build for production:
```bash
npm run build
```

## Features

- Task result display
- Syntax highlighting for code
- Task details view

## Dependencies

- React: UI library
- React Syntax Highlighter: Code highlighting
- TailwindCSS: Styling
- Vite: Build tool
- Module Federation: Micro frontend architecture
