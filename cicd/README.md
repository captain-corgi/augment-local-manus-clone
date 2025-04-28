# CI/CD

This directory contains CI/CD pipelines and scripts for the Local AI Agent System.

## Overview

The CI/CD setup automates the building, testing, and deployment of the system. It ensures that all changes are properly tested before being deployed to production.

## Pipelines

The `pipelines` directory contains CI/CD pipeline configurations:

- `backend-pipeline.yml`: Pipeline for backend services
- `frontend-pipeline.yml`: Pipeline for frontend applications
- `integration-pipeline.yml`: Pipeline for integration tests

## Scripts

The `scripts` directory contains automation scripts:

- `build.sh`: Builds all components
- `test.sh`: Runs all tests
- `deploy.sh`: Deploys the system
- `setup-dev.sh`: Sets up the development environment

## Usage

### Setting Up Development Environment

```bash
./cicd/scripts/setup-dev.sh
```

### Running Tests

```bash
./cicd/scripts/test.sh
```

### Building the System

```bash
./cicd/scripts/build.sh
```

### Deploying the System

```bash
./cicd/scripts/deploy.sh
```

## CI/CD Flow

1. **Commit**: Developer commits code to the repository
2. **Build**: CI/CD system builds the code
3. **Test**: CI/CD system runs unit tests and integration tests
4. **Deploy**: If tests pass, the code is deployed to the target environment

## Environments

- **Development**: For development and testing
- **Production**: For production use

## Monitoring

The CI/CD system monitors the health of the deployed services and can automatically roll back if issues are detected.
