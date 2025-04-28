# Task Management App

This is the Task Management micro frontend for the local AI agent system. It handles task submission and listing.

## Architecture

The Task Management App follows the Micro Frontend architecture:

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

- Task submission form
- Task listing
- Task status display

## Dependencies

- React: UI library
- React Hook Form: Form handling
- TailwindCSS: Styling
- Vite: Build tool
- Module Federation: Micro frontend architecture
