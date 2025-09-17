# Repository Guidelines

## Project Structure & Module Organization
- `breeze.go`: core library (agents, AI, chat, code, batch).
- `cmd/breeze/`: CLI entrypoint (`main.go`). Build target for the binary.
- `examples/`: runnable samples (team collaboration, apps, utilities).
- `breeze_test.go`: integration-oriented tests (skipped unless Ollama is available).
- `bin/`: build artifacts created by `build.sh` or manual builds.
- `go.mod`, `go.sum`: module metadata (Go 1.21).

## Build, Test, and Development Commands
- Quick tasks: `make build`, `make test`, `make fmt`, `make vet`, `make lint`, `make cross`
- Build CLI: `go build ./cmd/breeze` (binary `./breeze`)
- Run CLI: `./breeze "Explain recursion"` or `./breeze chat "Hello"` (or `make run ARGS='chat "Hello"'`)
- Cross-compile: `bash build.sh` (artifacts in `bin/`)
- Tests: `go test ./...` (Ollama-dependent tests are `t.Skip` by default)

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

## Commit & Pull Request Guidelines
- Conventional Commits required. Example: `feat(cli): add concise mode flag`.
- Enable commit hook: `git config core.hooksPath .githooks` (enforces message format).
- PRs must include: clear description, linked issue (if any), usage/output example, and test notes.
- Keep diffs minimal and focused; update README or examples when behavior changes.

## Security & Configuration Tips
- Do not hardcode secrets or model endpoints; prefer environment variables when needed.
- Don’t commit local artifacts; only binaries produced by release workflows should live in `bin/`.
- Validate inputs in CLI; fail fast with actionable messages.
