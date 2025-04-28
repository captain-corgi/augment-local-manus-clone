# Contributing to Local AI Agent System

Thank you for considering contributing to the Local AI Agent System! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please read it before contributing.

## How Can I Contribute?

### Reporting Bugs

- Check if the bug has already been reported in the [Issues](https://github.com/yourusername/local-ai-agent-system/issues)
- If not, create a new issue using the bug report template
- Include as much detail as possible:
  - Steps to reproduce
  - Expected behavior
  - Actual behavior
  - Screenshots if applicable
  - Environment details

### Suggesting Enhancements

- Check if the enhancement has already been suggested in the [Issues](https://github.com/yourusername/local-ai-agent-system/issues)
- If not, create a new issue using the feature request template
- Describe the enhancement in detail
- Explain why it would be useful

### Pull Requests

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes
4. Run tests: `make test`
5. Commit your changes using conventional commit format: `feat: add new feature`
6. Push to your branch: `git push origin feature/your-feature-name`
7. Create a pull request

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/local-ai-agent-system.git
cd local-ai-agent-system
```

2. Set up the development environment:
```bash
make setup
```

3. Run the system:
```bash
make run
```

## Coding Guidelines

### Go

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Write tests for your code
- Document your code with comments

### JavaScript/TypeScript

- Follow the ESLint configuration
- Use TypeScript for type safety
- Write tests for your code
- Document your code with JSDoc comments

### Git Commit Messages

We use [Conventional Commits](https://www.conventionalcommits.org/) for commit messages:

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: Code changes that neither fix a bug nor add a feature
- `perf`: Performance improvements
- `test`: Adding or fixing tests
- `build`: Changes to the build system
- `ci`: Changes to CI configuration
- `chore`: Other changes that don't modify source or test files

Example: `feat(backend): add new endpoint for task filtering`

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update the documentation if applicable
3. The PR should work on all supported platforms
4. The PR must pass all tests and CI checks
5. The PR must be reviewed by at least one maintainer

## License

By contributing to this project, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).
