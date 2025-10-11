## Python Package Installation

Breeze can now be installed as a Python package using pip!

### Prerequisites

1. **Go 1.21+**: Required to build the Breeze binary
   - Download from [golang.org](https://golang.org/dl/)
   - Verify installation: `go version`

2. **Ollama**: Required for LLM interactions
   - Install from [ollama.ai](https://ollama.ai)
   - Verify installation: `ollama --version`

### Installation

#### Option 1: Install from source (development)

```bash
# Clone the repository
git clone https://github.com/yossideutsch1973/breeze.git
cd breeze

# Build the Go binary
go build ./cmd/breeze

# Install the Python package in development mode
pip install -e .
```

**Note**: If you encounter network issues during build, you can build without isolation:
```bash
python -m build --no-isolation
pip install dist/breeze_ai-2.0.0-py3-none-any.whl
```

#### Option 2: Install from PyPI (when published)

```bash
pip install breeze-ai
```

### Usage

#### Python API

```python
import breeze

# Simple AI query
response = breeze.ai("Explain quantum physics")
print(response)

# Conversational AI
breeze.chat("Hello!")
response = breeze.chat("Tell me about Python")
print(response)

# Code generation
code = breeze.code("Write a factorial function in Python")
print(code)

# Clear conversation history
breeze.clear()

# Batch processing
results = breeze.batch(["Explain AI", "Explain ML"])
for result in results:
    print(result)
```

#### Command Line

The Python package also provides a `breeze` command:

```bash
# Simple query
breeze "Explain recursion"

# Chat mode
breeze chat "Hello!"
breeze chat "Tell me more"

# Code generation
breeze code "Write a Python HTTP server"

# Clear history
breeze clear
```

### Python vs Go Usage

The Python package is a wrapper around the Go binary and provides:
- Installer via `pip install`.
- Pythonic API surface.
- Command-line interface mirroring the Go binary.
- Subset of features compared to the native Go library.

For full features (streaming, document processing, team collaboration), use the Go library directly:

```go
import "github.com/user/breeze"

response := breeze.AI("prompt", breeze.WithConcise(), breeze.WithDocs("file.pdf"))
```

### Development

To contribute to the Python package:

```bash
# Install development dependencies
pip install -e ".[dev]"

# Run tests
pytest python/tests/

# Format code
black python/
```

### Troubleshooting

**"Breeze binary not found"**
- Ensure Go is installed: `go version`
- Build the binary: `go build ./cmd/breeze`
- Verify binary exists in the package directory

**"Connection refused" or Ollama errors**
- Ensure Ollama is installed and running
- Start Ollama: `ollama serve` (runs automatically on most systems)
- Verify with: `ollama list`

### Notes

- The Python wrapper calls the Go binary via subprocess
- First run will download required AI models (may take time)
- Requires internet connection for initial model download
- After setup, works completely offline
