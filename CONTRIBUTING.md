# Contributing to Open Source Demo

Thank you for your interest in contributing to Open Source Demo! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

By participating in this project, you are expected to uphold our [Code of Conduct](./CODE_OF_CONDUCT.md). Please read it before participating.

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Git

### Development Environment

This project includes a devcontainer configuration, which allows you to develop in a containerized environment with all necessary dependencies pre-installed.

#### Using GitHub Codespaces

1. Navigate to the repository on GitHub
2. Click on the "Code" button
3. Select the "Codespaces" tab
4. Click "Create codespace on main"

#### Using VS Code with Dev Containers

1. Install the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension
2. Clone this repository: `git clone https://github.com/amenocal/open-source-demo.git`
3. Open the repository in VS Code
4. When prompted, click "Reopen in Container"

### Building and Testing

To build the project:

```bash
go build
```

To run tests:

```bash
go test ./...
```

## How to Contribute

### Reporting Bugs

Bugs are tracked as GitHub issues. Create an issue using the bug report template and provide the following information:

- Clear, descriptive title
- Steps to reproduce the problem
- Expected behavior
- Actual behavior
- Any additional context (screenshots, logs, etc.)

### Suggesting Enhancements

Enhancement suggestions are also tracked as GitHub issues. Create an issue using the feature request template and provide:

- Clear, descriptive title
- Detailed description of the proposed functionality
- Any relevant examples or mockups
- Explanation of why this enhancement would be useful

### Pull Requests

Follow these steps to submit a pull request:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes
4. Run tests to ensure they pass: `go test ./...`
5. Commit your changes using [conventional commit messages](https://www.conventionalcommits.org/)
6. Push to your fork: `git push origin feature/your-feature-name`
7. Create a pull request against the `main` branch

#### Pull Request Guidelines

- Use a clear, descriptive title
- Include a reference to related issues using GitHub keywords (e.g., "Fixes #123")
- Update documentation as needed
- Add tests for new features
- Follow Go code style and formatting conventions
- Apply appropriate labels to your PR
- Keep PRs focused on a single task

## Versioning and Release Process

This project uses [Semantic Versioning](https://semver.org/) and leverages GitHub's Release Drafter for automated release notes.

When creating a PR, add one of the following labels to indicate the type of change:

- `major` for breaking changes
- `minor` or `enhancement` for new features
- `patch` or `bug` for bug fixes

## Recognition of Contributors

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification to recognize all contributions. Contributors will be acknowledged in the README.md file.

## License

By contributing to this project, you agree that your contributions will be licensed under the project's [MIT License](./LICENSE).

## Questions?

If you have any questions or need help, please open an issue or contact the maintainers.
