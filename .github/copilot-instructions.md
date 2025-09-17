# Breeze AI Agent Coding Instructions

## Project Overview
Breeze is a Go library for ultra-simple local LLM interactions via Ollama, with a focus on developer experience, 1-line APIs, and team collaboration. It auto-manages Ollama, models, and conversation state, and supports document processing, streaming, and cross-platform builds.

## Architecture & Key Patterns
- **Single-file core**: All main logic (APIs, agent framework, document processing) is in `breeze.go`.
- **Functional options**: All configuration uses options like `WithModel`, `WithTemp`, `WithDocs`, etc.
- **Global state**: A single `defaultClient` manages Ollama, model selection, and shared context.
- **Auto-management**: Ollama is started/pulled as needed; models are auto-selected and downloaded.
- **Team collaboration**: Multi-agent workflows (see `TeamDevCollab`, `Phase`, and `Agent` in `breeze.go`).
- **Minimal dependencies**: Only Go stdlib and Ollama required for core; advanced examples may use `gorilla/mux` or `go-sqlite3`.

## Developer Workflows
- **Build**: `make build` (or `go build ./cmd/breeze`)
- **Test**: `make test` (integration tests require Ollama running; skipped in CI)
- **Format**: `make fmt`
- **Cross-compile**: `make cross` or `./build.sh` (outputs to `bin/`)
- **Run CLI**: `./breeze "prompt"` or `make run ARGS='chat "Hello"'`

## CLI & API Usage
- CLI commands are routed in `cmd/breeze/main.go` (e.g., `chat`, `code`, `clear`, or default to `AI`).
- Library usage is always 1-line, e.g.:
  - `breeze.AI("prompt")`
  - `breeze.Chat("prompt")`
  - `breeze.Code("generate code")`
  - With options: `breeze.AI("prompt", breeze.WithModel("codellama"), breeze.WithDocs("file.pdf"))`

## Team Collaboration & Agents
- Define agents with `Agent{Name, Role, Expertise, Personality}`.
- Use `TeamDevCollab` for multi-agent workflows (see `examples/sw_engineering_collab.go`).
- Collaboration phases can be parallel or sequential (`IsParallel` in `Phase`), or use composable methods (see `CollaborationMethod`).
- Progress callbacks: `OnPhaseComplete`, `OnAgentResponse`.

## Model & Document Handling
- Preferred models: `gpt-oss`, `codellama`, `llama2`, `mistral` (auto-selected, fallback if unavailable).
- Document processing: `WithDocs` supports PDF, DOCX, TXT (auto-extracts text).
- Streaming: Use `breeze.Stream` for token-wise output.

## Error Handling & Testing
- Functions return error strings, not panics; errors are user-friendly.
- Integration tests in `breeze_test.go` require Ollama (use `t.Skip` if not running).

## File/Directory Reference
- `breeze.go`: All core logic, APIs, agent/collab framework
- `cmd/breeze/main.go`: CLI routing
- `examples/`: Usage patterns, agent workflows, document processing
- `build.sh`, `Makefile`: Build/test/cross-compile commands
- `breeze_test.go`: Integration tests

## Project-specific Conventions
- All new features go in `breeze.go` (single-file core)
- Use functional options for all config
- Add CLI support in `cmd/breeze/main.go` if needed
- Add usage examples in `examples/` for new features
- Advanced dependencies (mux/sqlite3) only in examples, not core

## Example Patterns
```go
// 1-line API with options
resp := breeze.AI("Summarize", breeze.WithDocs("file.pdf"), breeze.WithConcise())

// Team collaboration
team := []breeze.Agent{{Name: "Alex", Role: "Engineer"}}
result := breeze.TeamDevCollab(team, nil, "Build a Go app")

// Streaming
breeze.Stream("prompt", func(token string) { fmt.Print(token) })
```