# Breeze AI Agent Instructions

## Project Overview
Breeze is a Go library for ultra-simple local LLM interactions via Ollama. It provides a 1-line API with zero configuration, focusing on developer experience and simplicity. Features include team collaboration frameworks, document processing, streaming responses, and cross-platform support.

## Architecture Patterns

### Core Design Philosophy
- **Single Responsibility**: Each function (`AI`, `Chat`, `Code`, `Stream`, `Batch`, `TeamDevCollab`) handles one specific use case
- **Functional Options**: Use `WithModel()`, `WithTemp()`, `WithContext()`, `WithDocs()`, `WithConcise()` for configuration (see `breeze.go:28-45`)
- **Global State Management**: Single `defaultClient` instance manages Ollama connection, conversation state, and shared knowledge
- **Auto-Management**: Automatically starts Ollama, pulls models, and selects best available model
- **Team Collaboration**: Multi-agent workflows with phases, parallel execution, and shared knowledge

### Key Components
- `breeze.go`: Core library with all public APIs, team collaboration framework, and document processing
- `cmd/breeze/main.go`: Minimal CLI wrapper (simple command routing)
- `build.sh`: Cross-platform compilation script with GOOS/GOARCH patterns
- `examples/`: Comprehensive examples including team development and document processing
- Minimal dependencies: Go standard library + Ollama (auto-managed) + optional gorilla/mux for advanced features

## Critical Workflows

### Building & Distribution
```bash
# Single platform build
go build ./cmd/breeze

# Cross-platform compilation (see build.sh)
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/breeze-linux-amd64 ./cmd/breeze

# All platforms at once
./build.sh
```

### Development Setup
1. Install Ollama: https://ollama.ai (auto-managed by library)
2. Clone repo and build: `go build ./cmd/breeze`
3. Run: `./breeze "test prompt"` or use library functions directly

### Team Collaboration Workflow
```go
// Define specialized agents
swTeam := []breeze.Agent{
    {Name: "Alex", Role: "Senior Engineer", Expertise: "Go development"},
    {Name: "Maria", Role: "Engineer", Expertise: "Data structures"},
}

// Use TeamDevCollab for complete development cycles
result := breeze.TeamDevCollab(swTeam, testTeam, projectDescription)
```

## Code Patterns & Conventions

### API Design
```go
// Simple usage (global functions)
response := breeze.AI("Explain quantum physics")
breeze.Chat("conversational prompt")
code := breeze.Code("generate code")

// With functional options
response := breeze.AI("prompt", breeze.WithModel("codellama"), breeze.WithTemp(0.1))
response := breeze.AI("prompt", breeze.WithConcise()) // Concise responses with streaming
response := breeze.AI("prompt", breeze.WithDocs("file.pdf")) // Document processing

// Team collaboration
results := breeze.TeamDevCollab(swTeam, testTeam, project)

// Streaming
breeze.Stream("prompt", func(token string) { fmt.Print(token) })

// Batch processing
results := breeze.Batch([]string{"prompt1", "prompt2"})
```

### Model Management
- **Preferred Models**: `gpt-oss`, `codellama`, `llama2`, `mistral` (in order of preference)
- **Auto-Pull**: Models are automatically pulled if not available
- **Smart Selection**: `selectBestModel()` chooses best available model
- **Code-Specific**: `Code()` function prefers `codellama` when available
- **Fallback Chain**: Graceful degradation to available models

### Document Processing
```go
// Process various file formats
response := breeze.AI("Summarize this", breeze.WithDocs("document.pdf"))
response := breeze.AI("Analyze these files", breeze.WithDocs("file1.txt", "file2.docx"))

// Automatic text extraction from:
// - PDF files (parses text objects between BT/ET)
// - DOCX files (XML parsing of word/document.xml)
// - TXT files (direct reading)
```

### Error Handling
- **Silent Failures**: Functions return error strings rather than panicking
- **Graceful Degradation**: Falls back to default model if preferred models unavailable
- **User-Friendly**: Error messages guide users (e.g., "Please install Ollama")
- **Non-blocking**: Auto-starts Ollama if not running

### HTTP Integration
- **Direct API Calls**: Raw HTTP requests to `http://localhost:11434`
- **JSON Marshaling**: Request/response handled with `encoding/json`
- **Connection Management**: Auto-detects and starts Ollama if not running
- **Dual Endpoints**: Uses `/api/generate` for single responses, `/api/chat` for conversations

### Team Collaboration Framework
```go
// Agent definition with personality
agent := breeze.Agent{
    Name: "Alex",
    Role: "Senior Engineer",
    Expertise: "Go development",
    Personality: "pragmatic and detail-oriented",
}

// Phase-based workflow
phases := []breeze.Phase{
    {
        Name: "Requirements Analysis",
        Description: "Analyze requirements",
        PromptTemplate: "Provide your expert analysis...",
        IsParallel: true,
        MaxConcurrency: 4,
    },
}

// Parallel execution with shared knowledge
collab := breeze.NewCollaboration(agents, phases)
results, _ := collab.Run("Build a task manager")
```

## Testing Approach
- **Integration Tests**: Tests require running Ollama (skipped in CI with `t.Skip()`)
- **Simple Assertions**: Basic non-empty response checks
- **Mock-Friendly**: Architecture supports dependency injection for testing
- **Real Environment**: Tests validate actual Ollama integration

## File Organization
- `breeze.go`: All public APIs, core logic, team collaboration, document processing
- `cmd/breeze/main.go`: Minimal CLI wrapper with command routing
- `examples/`: Complete usage examples including team development and document processing
- `build.sh`: Cross-compilation automation for Linux/macOS/Windows
- `bin/`: Platform-specific binaries
- `breeze_test.go`: Integration tests (require Ollama)

## Common Patterns to Follow

### When Adding New Features
1. Add to `breeze.go` (single file architecture)
2. Use functional options pattern for configuration
3. Handle errors gracefully with user-friendly messages
4. Update `cmd/breeze/main.go` for CLI access if needed
5. Add example usage in `examples/` folder
6. Consider team collaboration integration if applicable

### When Modifying Core Logic
- Preserve the 1-line API simplicity
- Maintain minimal dependencies (Go stdlib + Ollama)
- Test with real Ollama instance
- Update cross-platform build compatibility
- Consider impact on team collaboration features

### When Working with Ollama Integration
- Use direct HTTP calls (no SDK dependencies)
- Handle connection failures gracefully
- Support both `/api/generate` and `/api/chat` endpoints
- Implement streaming with JSON decoder pattern
- Auto-manage model pulling and selection

### When Adding Team Collaboration Features
- Use the `Agent` struct with Name, Role, Expertise, Personality fields
- Implement phases with parallel execution support
- Maintain shared knowledge across agents
- Provide progress callbacks (`OnPhaseComplete`, `OnAgentResponse`)
- Support both sequential and parallel execution modes

## Key Files to Reference
- `breeze.go:60-90`: Ollama auto-management and model selection
- `breeze.go:91-130`: Core AI/generation logic with document processing
- `breeze.go:131-170`: Chat conversation management
- `breeze.go:220-250`: Streaming implementation with JSON decoding
- `breeze.go:550-650`: Team collaboration framework and agent management
- `breeze.go:900-950`: TeamDevCollab convenience function
- `build.sh`: Cross-compilation patterns with GOOS/GOARCH
- `examples/team_development.go`: Complete team collaboration example
- `examples/`: Complete usage examples for all features