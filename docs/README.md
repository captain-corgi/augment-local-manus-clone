# Local AI Agent System Documentation

This directory contains documentation for the Local AI Agent System, a self-hosted AI agent that runs entirely on the user's hardware for privacy and cloud independence.

## Overview

The Local AI Agent System is designed to provide autonomous task execution capabilities, including coding, web browsing, task planning, and filesystem interaction. It follows Domain-Driven Design (DDD) principles, uses microservices with Clean Architecture for the backend, and implements a Micro Front End architecture for the frontend.

## Contents

- [Installation Guide](installation.md): Instructions for setting up the system
- [Usage Guide](usage.md): How to use the system
- [Architecture Overview](architecture.md): System architecture and design

## System Components

### Backend Services

- **Task Service**: Manages task lifecycle (submission, status, results)
- **AI Service**: Interfaces with the local LLM (Ollama) for reasoning and planning
- **Code Execution Service**: Executes code in sandboxed Docker containers
- **Web Browsing Service**: Navigates web pages using chromedp
- **Filesystem Service**: Performs secure file operations within /app/workspace

### Frontend Applications

- **Shell**: Hosts micro frontends and provides main layout and navigation
- **Task Management App**: Handles task submission and listing
- **Result Viewer App**: Displays task results

## Development

The system is built using the following technologies:

- **Backend**: Go, Gin, SQLite, Docker, chromedp
- **Frontend**: React, TypeScript, TailwindCSS, Vite, Module Federation
- **AI**: Ollama with DeepSeek-R1 model

All components follow strict architectural patterns and include comprehensive tests using Table-Driven Testing.
