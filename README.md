# Breeze v2.0 - Ultra-Concise Local LLM Library

Breeze makes local LLM development a breeze! Get AI responses in **1 line of code** with zero configuration.

## 🎯 **Project Status: Production Ready**

✅ **Core Features Complete**: AI queries, chat, code generation, streaming, batch processing
✅ **Advanced Features**: Team collaboration, document processing, concise mode
✅ **Cross-Platform**: Linux, macOS, Windows binaries available
✅ **Examples**: 10+ comprehensive examples including AI team collaboration
✅ **Documentation**: Complete API docs and usage examples
✅ **Architecture**: Clean, minimal dependencies, production-ready

## 🚀 Quick Start

### Installation

1. Install [Ollama](https://ollama.ai)
2. Clone this repo
3. Build: `go build ./cmd/breeze`

### Usage

```bash
# Simple AI query
./breeze "Explain quantum physics"

# Conversational AI
./breeze chat "Hello!"
./breeze chat "Tell me more"

# Code generation
./breeze code "Write a Go HTTP server"

# Clear conversation
./breeze clear
```

## 📚 Library Usage

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

## ✨ Features

- **Zero Configuration**: Auto-detects and starts Ollama
- **Intelligent Model Selection**: Automatically chooses the best available model
- **Conversation Memory**: Remembers context across chat calls
- **Streaming Support**: Real-time token streaming
- **Batch Processing**: Concurrent request processing
- **Document Processing**: Process PDF, DOCX, and TXT files
- **Concise Mode**: Brief, focused responses with streaming
- **Team Collaboration**: Multi-agent AI collaboration framework
- **Cross-Platform**: Single binary for Linux, macOS, Windows

## 🤝 **Team Collaboration Framework**

Breeze includes a powerful team collaboration system for complex AI workflows:

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

### **Real Examples Built:**

- **Task Manager**: Complete CLI app with priorities, persistence, and testing
- **Finance Tracker**: Financial analysis and reporting system
- **Super Resolution**: Image processing algorithm implementation
- **Startup Founders**: Multi-agent business planning simulation

## 📁 **Project Structure**

```
breeze/
├── breeze.go              # Main library with team collaboration
├── breeze_test.go         # Comprehensive test suite
├── README.md              # This documentation
├── build.sh               # Cross-platform build script
├── go.mod & go.sum        # Go module dependencies
├── bin/                   # Pre-built binaries (Linux, macOS, Windows)
├── cmd/breeze/            # CLI application
├── examples/              # 10+ example applications
│   ├── team_development.go    # AI team collaboration demo
│   ├── task_manager/          # Complete working task manager app
│   ├── finance_tracker.go     # Financial analysis system
│   └── ...                    # More examples
└── .github/               # GitHub Actions and configuration
```

## 🏗️ Architecture

- **Minimal Dependencies**: Only Go standard library + Ollama (auto-managed)
- **Single Binary**: Cross-compiled for all platforms
- **Goroutines**: Concurrent processing for team collaboration
- **HTTP Client**: Direct Ollama API integration
- **Functional Options**: Clean configuration pattern
- **Team Framework**: Multi-agent collaboration system

## 🔧 Advanced Usage

### Options

```go
// Model selection
breeze.AI("prompt", breeze.WithModel("mistral"))

// Temperature control
breeze.AI("prompt", breeze.WithTemp(0.5))

// Add context
breeze.AI("Explain this code", breeze.WithContext(codeSnippet))

// Document processing
breeze.AI("Summarize this document", breeze.WithDocs("document.pdf"))
breeze.AI("Analyze these files", breeze.WithDocs("file1.txt", "file2.docx"))

// Concise responses with streaming
breeze.AI("Explain quantum physics", breeze.WithConcise())
breeze.Chat("Help me debug this code", breeze.WithConcise())
```

### Streaming

```go
breeze.Stream("Write a novel", func(token string) {
    fmt.Print(token) // Real-time output
})
```

### Batch Processing

```go
prompts := []string{"Task 1", "Task 2", "Task 3"}
results := breeze.Batch(prompts)
```

## 🛠️ Building

```bash
# Build for current platform
go build ./cmd/breeze

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" ./cmd/breeze

# Cross-compile for Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" ./cmd/breeze

# Cross-compile for macOS
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" ./cmd/breeze
```

## 📋 Requirements

- Go 1.21+
- Ollama (auto-managed)

## 🤝 Contributing

Contributions welcome! Focus on developer experience and simplicity.

## � **Handoff Status: Complete**

This Breeze project is **production-ready** and fully documented:

✅ **Core Library**: Complete with team collaboration, streaming, batch processing
✅ **CLI Tool**: Cross-platform binaries available
✅ **Examples**: 10+ working examples including AI team collaboration
✅ **Documentation**: Comprehensive README and inline code docs
✅ **Architecture**: Clean, minimal dependencies, well-organized
✅ **Testing**: Test suite included
✅ **Build System**: Cross-compilation scripts ready

### **Key Achievements:**
- **Team Collaboration Framework**: Multi-agent AI workflows
- **Task Manager App**: Complete working application generated by AI teams
- **Cross-Platform Support**: Linux, macOS, Windows binaries
- **Advanced Features**: Document processing, concise mode, streaming
- **Clean Architecture**: Minimal dependencies, functional options pattern

### **Ready for Use:**
- Clone repository
- Run `go build ./cmd/breeze`
- Use pre-built binaries in `bin/` directory
- Explore examples in `examples/` directory

## �📄 License

MIT License
