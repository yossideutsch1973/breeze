# Breeze AI Agent Instructions

## Project Overview
Breeze is a Go library for ultra-simple local LLM interactions via Ollama. It provides a 1-line API with zero configuration, focusing on developer experience and simplicity.

## Architecture Patterns

### Core Design Philosophy
- **Single Responsibility**: Each function (`AI`, `Chat`, `Code`, `Stream`, `Batch`) handles one specific use case
- **Functional Options**: Use `WithModel()`, `WithTemp()`, `WithContext()` for configuration (see `breeze.go:28-45`)
- **Global State Management**: Single `defaultClient` instance manages Ollama connection and conversation state
- **Auto-Management**: Automatically starts Ollama and pulls models as needed

### Key Components
- `breeze.go`: Core library with all public APIs
- `cmd/breeze/main.go`: CLI wrapper (simple command routing)
- `build.sh`: Cross-platform compilation script
- `examples/`: Example programs demonstrating capabilities
- Minimal dependencies: Go standard library + Ollama (auto-managed)

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
1. Install Ollama: https://ollama.ai
2. Clone repo
3. Build: `go build ./cmd/breeze`
4. Run: `./breeze "test prompt"`

## Code Patterns & Conventions

### API Design
```go
// Simple usage (global functions)
response := breeze.AI("prompt")
breeze.Chat("conversational prompt")
code := breeze.Code("generate code")

// With functional options
response := breeze.AI("prompt", breeze.WithModel("codellama"), breeze.WithTemp(0.1))
response := breeze.AI("prompt", breeze.WithConcise()) // Concise responses with streaming
response := breeze.AI("prompt", breeze.WithDocs("file.pdf")) // Document processing

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

### Error Handling
- **Silent Failures**: Functions return error strings rather than panicking
- **Graceful Degradation**: Falls back to default model if preferred models unavailable
- **User-Friendly**: Error messages guide users (e.g., "Please install Ollama")

### HTTP Integration
- **Direct API Calls**: Raw HTTP requests to `http://localhost:11434`
- **JSON Marshaling**: Request/response handled with `encoding/json`
- **Connection Management**: Auto-detects and starts Ollama if not running

## Testing Approach
- **Integration Tests**: Tests require running Ollama (skipped in CI)
- **Simple Assertions**: Basic non-empty response checks
- **Mock-Friendly**: Architecture supports dependency injection for testing

## File Organization
- `breeze.go`: All public APIs and core logic
- `cmd/breeze/main.go`: Minimal CLI wrapper
- `example/main.go`: Comprehensive usage examples
- `build.sh`: Cross-compilation automation
- `bin/`: Platform-specific binaries

## Common Patterns to Follow

### When Adding New Features
1. Add to `breeze.go` (single file architecture)
2. Use functional options pattern for configuration
3. Handle errors gracefully with user-friendly messages
4. Update `cmd/breeze/main.go` for CLI access
5. Add example usage in `examples/` folder

### When Modifying Core Logic
- Preserve the 1-line API simplicity
- Maintain minimal dependencies (Go stdlib + Ollama)
- Test with real Ollama instance
- Update cross-platform build compatibility

### When Working with Ollama Integration
- Use direct HTTP calls (no SDK dependencies)
- Handle connection failures gracefully
- Support both `/api/generate` and `/api/chat` endpoints
- Implement streaming with JSON decoder pattern

## Key Files to Reference
- `breeze.go:60-90`: Ollama auto-management and model selection
- `breeze.go:91-130`: Core AI/ generation logic
- `breeze.go:131-170`: Chat conversation management
- `breeze.go:220-250`: Streaming implementation
- `build.sh`: Cross-compilation patterns
- `examples/`: Complete usage examples