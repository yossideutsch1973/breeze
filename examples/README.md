# Breeze Examples

This folder contains example programs demonstrating Breeze's capabilities.

## Examples

### `main.go` - Hello World

A simple introduction to Breeze showing basic AI interaction.

```bash
go run main.go
```

### `sw_engineering_collab.go` - SW Engineering Team with Peer Review

A comprehensive demonstration of **professional software engineering collaboration** with proper peer review:

- **🎯 Real Peer Review System**: Each engineer does work AND reviews others' work
- **📊 Scoring & Consolidation**: Reviews are scored (1-10) and intelligently merged
- **📈 Progress Bar**: Clean progress indication without console spam
- **📁 File Output Only**: Final comprehensive report saved to timestamped file
- **👥 Homogeneous Team**: All agents are SW engineers with different specializations
- **🔄 Structured Workflow**: Individual work → Peer review → Consolidation → Final report

```bash
go run sw_engineering_collab.go
```

**Features:**

- ✅ Progress bar (no console clutter)
- ✅ Professional peer review process
- ✅ Scoring system for review consolidation
- ✅ Clean file output (no console spam)
- ✅ Homogeneous team (all SW engineers)
- ✅ Work + Review workflow per agent

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

- ✅ Basic AI queries
- ✅ Conversation management
- ✅ Model selection
- ✅ Creative collaboration
- ✅ Iterative improvement
- ✅ Fun AI interactions
