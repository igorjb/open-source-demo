# gh-license

A GitHub CLI extension that helps you retrieve license information from open source projects.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/amenocal/open-source-demo)](https://github.com/amenocal/open-source-demo/releases)
[![CI Status](https://github.com/amenocal/open-source-demo/actions/workflows/ci.yml/badge.svg)](https://github.com/amenocal/open-source-demo/actions/workflows/ci.yml)

## 🚀 About

This repository serves as both a demonstration of open source project best practices and a functional GitHub CLI extension. The extension allows you to quickly retrieve and analyze license information from any GitHub repository.

This project showcases:

- 📄 Proper open source licensing (MIT)
- 🛠️ Development using GitHub Codespaces / Dev Containers
- 🔄 CI/CD automation and security scanning
- 🔍 Advanced security with CodeQL analysis and Dependabot
- 📝 Comprehensive documentation

## 📋 Prerequisites

- [GitHub CLI](https://cli.github.com/) version 2.0.0 or higher

## 📦 Installation

```bash
gh extension install amenocal/open-source-demo
```

## 🔧 Usage

Retrieve license information for a GitHub repository:

```bash
gh license [repository]
```

Example:

```bash
# Get license for the current repository
gh license .

# Get license for a specific repository
gh license amenocal/open-source-demo

# Show detailed license information
gh license amenocal/open-source-demo --detailed
```

### Options

```console
--detailed    Display detailed license information including permissions, conditions, and limitations
--format      Output format (text, json, yaml) [default: "text"]
--help        Show help for command
```

## ✨ Features

- Identify the license type of any GitHub repository
- Show permissions, conditions, and limitations of different licenses
- Compare license compatibility between projects
- Output in different formats (text, JSON, YAML)

## 🧪 Development

For detailed instructions on setting up your development environment, building the project, and using GitHub Codespaces or Dev Containers, please refer to our [Contributing Guidelines](CONTRIBUTING.md).

## 🤝 Contributing

Contributions of any kind are welcome! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details on:

- Setting up your development environment
- Building and testing the project
- Submitting bug reports and feature requests
- Pull request process and guidelines
- Release process and versioning

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification.

## 🔒 Security

If you discover any security issues, please follow our [Security Policy](SECURITY.md) for reporting.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.