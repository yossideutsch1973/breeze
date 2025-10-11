# Breeze v2.0 - Local LLM Toolkit

Breeze is a local LLM toolkit written in Go with a Python wrapper. It focuses on practical agent orchestration, CLI tooling, and runnable examples against Ollama installs.

📋 **[View the API Cheatsheet](CHEATSHEET.md)** - Quick reference for all features

## Status & Scope

- Built for local experimentation with Ollama-backed models.
- Actively maintained by a single developer.
- Uses Go for core runtime/agents and Python for a lightweight wrapper.
- Tested on macOS and Linux; Windows builds exist but are not exercised regularly.

## 🚀 Quick Start

### Installation

#### Option 1: Python Package (pip install)

```bash
# Prerequisites: Go 1.21+ and Ollama installed
pip install breeze-ai

# Or install from source
git clone https://github.com/yossideutsch1973/breeze.git
cd breeze
go build ./cmd/breeze  # Build the Go binary
pip install -e .        # Install Python package
```

See [PYTHON_INSTALL.md](PYTHON_INSTALL.md) for detailed Python installation instructions.

#### Option 2: Go Binary (native)

1. Install [Ollama](https://ollama.ai)
2. Clone this repo
3. Build: `go build ./cmd/breeze`

### Usage

#### Python API

```python
import breeze

# Simple AI query
response = breeze.ai("Explain quantum physics")
print(response)

# Conversational AI
breeze.chat("Hello!")
breeze.chat("Tell me more")

# Code generation
code = breeze.code("Write a Python function")
print(code)

# Clear conversation
breeze.clear()
```

#### Command Line (works for both pip install and Go binary)

```bash
# Simple AI query
breeze "Explain quantum physics"

# Conversational AI
breeze chat "Hello!"
breeze chat "Tell me more"

# Code generation
breeze code "Write a Go HTTP server"

# Clear conversation
breeze clear
```

## 📚 Library Usage

### Python Library

```python
import breeze

# Ultra-simple API
response = breeze.ai("Explain recursion")
print(response)

# Conversational
breeze.chat("Hello AI!")
breeze.chat("Help me with Python")

# Code-focused
code = breeze.code("Write a factorial function")
print(code)

# Batch processing
results = breeze.batch(["Explain AI", "Explain ML"])
for result in results:
    print(result)

# Clear conversation
breeze.clear()
```

Note: The Python wrapper provides basic functionality. For advanced features like streaming, document processing, and team collaboration, use the Go library directly.

### Go Library

```go
package main

import "github.com/user/breeze"

func main() {
    // Ultra-simple API
    response := breeze.AI("Explain recursion")
    println(response)

    // Conversational
    breeze.Chat("Hello AI!")
    breeze.Chat("Help me with Go")

    // Code-focused
    code := breeze.Code("Write a factorial function")
    println(code)

    // With options
    response := breeze.AI("Complex prompt", breeze.WithModel("codellama"), breeze.WithTemp(0.1))

    // Streaming
    breeze.Stream("Write a story", func(token string) {
        fmt.Print(token)
    })

    // Batch processing
    results := breeze.Batch([]string{"Explain AI", "Explain ML"})

    // Document processing
    response := breeze.AI("Summarize this document", breeze.WithDocs("document.pdf"))
    response := breeze.AI("Analyze these files", breeze.WithDocs("file1.txt", "file2.docx"))

    // Concise responses with streaming
    response := breeze.AI("Explain quantum physics", breeze.WithConcise())
    response := breeze.Chat("Help me debug this code", breeze.WithConcise())
}
```

## Capabilities

- Detects running Ollama instance and fails with actionable errors when missing.
- Provides conversational helpers that maintain short-lived session state.
- Exposes streaming, batch evaluation, and document parsing entry points.
- Includes multi-agent collaboration patterns centered on software tasks.
- Ships with Go/Python examples and a CLI for local workflows.

## Collaboration Framework

Breeze includes a team collaboration module used in the examples to script multi-agent workflows:

```go
// Define specialized AI agents
swTeam := []breeze.Agent{
    {Name: "Alex", Role: "Senior Engineer", Expertise: "Go development"},
    {Name: "Maria", Role: "Engineer", Expertise: "Data structures"},
}

testTeam := []breeze.Agent{
    {Name: "David", Role: "QA Engineer", Expertise: "Testing"},
    {Name: "Sarah", Role: "Test Automation", Expertise: "Integration"},
}

// Run collaborative development
result := breeze.TeamDevCollab(project, swTeam, testTeam)
```

### Example Scenarios

- Task manager CLI with lightweight persistence.
- Finance tracker that summarizes CSV files.
- Super-resolution proof of concept using external tools.
- Startup planning prompts demonstrating agent hand-offs.

## 📁 **Project Structure**

```
breeze/
├── breeze.go              # Core library including team collaboration helpers
├── breeze_test.go         # Unit tests
├── README.md              # Project documentation
├── build.sh               # Cross-platform build script
├── go.mod & go.sum        # Go module metadata
├── cmd/breeze/            # CLI entry point
├── examples/              # Example applications
│   ├── team_development.go
│   ├── finance_tracker.go
│   └── ...
└── .github/               # GitHub Actions and configuration
```

## Architecture Notes

- Minimal dependencies: standard library plus Ollama.
- Single binary build targets.
- Goroutines handle collaboration workflows.
- Direct HTTP integration with Ollama.
- Functional options configure runtime behaviour.

## 🔧 Advanced Usage

### Options

```