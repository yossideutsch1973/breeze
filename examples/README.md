# Breeze Examples

This directory collects runnable examples demonstrating different collaboration patterns. Each example is intended to be small and approachable.

## Examples

### `main.go` - Hello World

A simple introduction to Breeze showing basic AI interaction.

```bash
go run main.go
```

### `sw_engineering_collab.go` - SW Engineering Team with Peer Review

Highlights:
- Generates a final report saved to a timestamped file.
- Emphasises file-based outputs, not console logging.
- Shows a review workflow across software engineering roles.

```bash
go run sw_engineering_collab.go
```

**Features:**

- Progress bar minimal console output
- Peer review process
- Scoring system for review consolidation
- File output stored under `sw_team_report_*.md`
- Homogeneous team (all software engineering roles)
- Work plus review workflow per agent

### `qwen3_mini_agents.go` - Qwen3 Mini-Agents Swarm

A powerful demonstration of **10 specialized mini-agents** using the qwen3:1.7b model! Shows:

- **10 Specialized Agents**: Each with unique personality and expertise (Logic Analyst, Code Architect, Memory Specialist, etc.)
- **Parallel Processing**: All agents working simultaneously on complex problems
- **Swarm Intelligence**: Collective problem-solving with shared knowledge
- **Real-time Progress**: Live tracking of agent contributions
- **Comprehensive Solutions**: Multi-phase collaborative approach
- **Qwen3:1.7b Model**: Efficient small model for specialized agent roles

```bash
go run qwen3_mini_agents.go
```

### `songwriter.go` - AI Collaboration

A fun example showing two AI instances collaborating on writing a funny song about programmers. Demonstrates:

- Creative brainstorming
- Critical review and feedback
- Iterative improvement
- **Concise responses with streaming output**
- How AI models can build on each other's ideas

```bash
go run songwriter.go
```

### `startup_founders.go` - Parallel AI Team

A jaw-dropping example showing **4 AI founders** building a complete startup from scratch in parallel! Demonstrates:

- **Parallel processing** with goroutines
- **Specialized AI roles** (CTO, CEO, CPO, Chief Scientist)
- **Inter-AI communication** through shared knowledge
- **Multi-phase collaboration** (analysis → architecture → strategy → synthesis)
- **Real-time concurrent execution**
- **Professional output** - complete business plan

```bash
go run startup_founders.go
```

## Running Examples

Make sure Ollama is installed and running, then:

```bash
# Build and run hello world
go run main.go

# Build and run songwriter collaboration
go run songwriter.go
```

## Features Demonstrated

What you can test quickly:
- Basic AI queries
- Conversation management
- Model selection
- Creative collaboration
- Iterative improvement
- Lightweight tooling exercises
