# Shell Application

This is the shell application for the local AI agent system. It hosts the micro frontends and provides the main layout and navigation.

## Architecture

The Shell Application follows the Micro Frontend architecture:

- It uses Vite's Module Federation plugin to load micro frontends dynamically
- It provides shared dependencies to the micro frontends
- It handles routing and navigation between micro frontends

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

## Micro Frontends

The Shell Application loads the following micro frontends:

- Task Management App: Handles task submission and listing
- Result Viewer App: Displays task results

## Dependencies

- React: UI library
- React Router: Routing
- TailwindCSS: Styling
- Vite: Build tool
- Module Federation: Micro frontend architecture
