# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`is` is a Go-based CLI tool that serves as an inspector for your environment. It allows users to check various system attributes like OS details, CLI versions, battery status, audio levels, environment variables, and more. The tool uses exit codes (0 for success, 1 for failure) to enable scripting and conditional execution.

## Development Commands

### Build and Run
```bash
# Build the project
go build -o is .

# Run directly with Go
go run . --help

# Build using the container script (requires Docker)
./bin/container-build
```

### Testing
```bash
# Run all Go tests
go test ./...

# Run specific tests
go test ./compare
go test ./parser

# Run BATS integration tests (requires bash and bats)
apk add bash  # if running in Alpine container
bats test/
```

### Code Quality
```bash
# Format code with gofumpt
gofumpt -w **/*.go

# Format imports
goimports -w **/*.go

# Run linter
golangci-lint run -c .golangci.yaml

# Run all code quality tools via precious
precious lint
precious tidy
```

### Container Development
```bash
# Start development environment
docker compose run --rm dev sh

# Build inside container
./bin/container-build
```

## Architecture

### Core Structure
- **main.go**: Entry point using Kong CLI parser, defines all top-level commands
- **api.go**: Type definitions for all CLI command structures and validation
- **types/types.go**: Common types including Context struct used throughout
- **Command handlers**: Each top-level command (arch, audio, battery, cli, fso, os, there, user, var, known) has its own .go file

### Key Packages
- **age/**: Time-based comparisons for files and commands
- **attr/**: Attribute extraction utilities  
- **audio/**: Audio level and mute status detection
- **battery/**: Battery information retrieval
- **command/**: Command execution with output capture
- **compare/**: Value comparison logic (string, numeric, version, etc.)
- **mac/**: macOS-specific functionality
- **ops/**: Operation definitions and mappings
- **os/**: Operating system detection and information
- **parser/**: CLI output parsing and version extraction
- **reader/**: File reading utilities
- **version/**: Version comparison utilities

### Command Categories
1. **System Info**: `arch`, `os`, `audio`, `battery`, `user`
2. **CLI Tools**: `cli version`, `cli age`, `cli output`, `there`
3. **Files**: `fso age`
4. **Environment**: `var`
5. **Information Display**: `known` (prints values without comparisons)

### Comparison Operators
The tool supports various comparison operators: `eq`, `ne`, `gt`, `gte`, `lt`, `lte`, `in`, `like` (regex), `unlike` (inverse regex)

### Testing Strategy
- Unit tests for individual packages (`*_test.go`)
- Integration tests using BATS framework in `test/` directory
- Test data in `testdata/` directory

### Key Design Patterns
- Kong-based CLI parsing with struct tags for command definitions
- Context pattern for passing debug flags and success state
- Modular architecture with separate packages for different functionalities
- Exit code based success/failure for shell scripting integration