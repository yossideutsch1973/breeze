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

## Project Structure & Module Organization
- `breeze.go`: core library (agents, AI, chat, code, batch).
- `cmd/breeze/`: CLI entrypoint (`main.go`). Build target for the binary.
- `examples/`: runnable samples (team collaboration, apps, utilities).
- `breeze_test.go`: integration-oriented tests (skipped unless Ollama is available).
- `bin/`: build artifacts created by `build.sh` or manual builds.
- `go.mod`, `go.sum`: module metadata (Go 1.21).

## Developer Workflows
- **Build**: `make build` (or `go build ./cmd/breeze`)
- **Test**: `make test` (integration tests require Ollama running; skipped in CI)
- **Format**: `make fmt`
- **Vet**: `make vet`
- **Lint**: `make lint` (uses golangci-lint)
- **Cross-compile**: `make cross` or `./build.sh` (outputs to `bin/`)
- **Run CLI**: `./breeze "prompt"` or `./breeze chat "Hello"` or `make run ARGS='chat "Hello"'`

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

## Coding Style & Naming Conventions
- Go formatting is canonical: run `go fmt ./...` before committing.
- Use clear, exported names (PascalCase) for public API, unexported (camelCase) for internals.
- Keep packages small and focused; keep CLI-only logic under `cmd/breeze`.
- Errors: return `error` values; prefer wrapping with context via `%w`.
- Files: `*_test.go` for tests; one responsibility per file when practical.

## Testing Guidelines
- Write table-driven tests for pure logic; mock or skip external calls.
- Run tests selectively: `go test -run TestAI ./...`
- Integration tests that hit Ollama should detect availability and skip when absent.
- Add examples in `examples/` for new features to document behavior.
- All tests must pass before committing.
- Maintain or improve test coverage (current minimum: 14%).
- Use race detection: `go test -race ./...`

## Commit & Pull Request Guidelines
- Conventional Commits required. Example: `feat(cli): add concise mode flag`.
- Enable commit hook: `git config core.hooksPath .githooks` (enforces message format).
- PRs must include: clear description, linked issue (if any), usage/output example, and test notes.
- Keep diffs minimal and focused; update README or examples when behavior changes.
- Code must be formatted with `go fmt`.
- No lint warnings.

## Security & Configuration Tips
- Do not hardcode secrets or model endpoints; prefer environment variables when needed.
- Don't commit local artifacts; only binaries produced by release workflows should live in `bin/`.
- Validate inputs in CLI; fail fast with actionable messages.
- Input sanitization: validate all user inputs.
- No secrets in source code.

## CI/CD Pipeline
- **Linting**: golangci-lint with 50+ linters
- **Testing**: Unit tests with race detection
- **Coverage**: Minimum 14% threshold enforced (integration tests skipped in CI)
- **Building**: Cross-platform compilation for Linux, macOS, Windows
- **Formatting**: `go fmt` compliance check
- See `.github/workflows/ci.yml` for full pipeline configuration

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