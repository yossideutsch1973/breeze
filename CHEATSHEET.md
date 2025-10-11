# Breeze API Cheatsheet

Quick reference guide for Breeze - organized by what you want to achieve.

---

## Basic AI Queries

### Simple one-time question
```go
response := breeze.AI("Explain quantum physics")
```

### Conversational chat (maintains context)
```go
breeze.Chat("Hello, I need help with Go")
breeze.Chat("Can you explain interfaces?")
breeze.Chat("Show me an example")
```

### Generate code
```go
code := breeze.Code("Write a factorial function in Go")
```

### Clear conversation history
```go
breeze.Clear()
```

---

## ⚙️ Configuration Options

### Choose a specific model
```go
response := breeze.AI("prompt", breeze.WithModel("codellama"))
response := breeze.AI("prompt", breeze.WithModel("mistral"))
```

### Control creativity (temperature)
```go
// More deterministic (0.0 - 0.5)
response := breeze.AI("prompt", breeze.WithTemp(0.1))

// More creative (0.7 - 1.0)
response := breeze.AI("prompt", breeze.WithTemp(0.9))
```

### Add context to your prompt
```go
context := "User is working on a Go web server project"
response := breeze.AI("How do I handle errors?", breeze.WithContext(context))
```

### Get concise responses (with streaming)
```go
response := breeze.AI("Explain machine learning", breeze.WithConcise())
```

### Combine multiple options
```go
response := breeze.AI(
    "Write a REST API",
    breeze.WithModel("codellama"),
    breeze.WithTemp(0.3),
    breeze.WithConcise(),
)
```

---

## 📄 Document Processing

### Process a single document
```go
response := breeze.AI("Summarize this document", breeze.WithDocs("report.pdf"))
```

### Process multiple documents
```go
response := breeze.AI(
    "Compare these documents",
    breeze.WithDocs("file1.pdf", "file2.docx", "file3.txt"),
)
```

### Supported formats
- PDF (`.pdf`)
- Word Documents (`.docx`)
- Text Files (`.txt`)

---

## 🌊 Streaming Responses

### Stream tokens as they arrive
```go
breeze.Stream("Write a story about space", func(token string) {
    fmt.Print(token)  // Print each token in real-time
})
```

### Stream with options
```go
breeze.Stream(
    "Explain AI",
    func(token string) { fmt.Print(token) },
    breeze.WithModel("mistral"),
    breeze.WithTemp(0.7),
)
```

---

## 📦 Batch Processing

### Process multiple prompts concurrently
```go
prompts := []string{
    "Explain AI",
    "Explain ML",
    "Explain deep learning",
}
results := breeze.Batch(prompts)

for i, result := range results {
    fmt.Printf("Result %d: %s\n", i+1, result)
}
```

### Batch with options
```go
results := breeze.Batch(prompts, breeze.WithModel("codellama"), breeze.WithTemp(0.5))
```

---

## 👥 Team Collaboration

### Basic multi-agent collaboration
```go
agents := []breeze.Agent{
    {
        Name:        "Alice",
        Role:        "Senior Developer",
        Expertise:   "Go programming",
        Personality: "pragmatic",
    },
    {
        Name:        "Bob",
        Role:        "Architect",
        Expertise:   "System design",
        Personality: "thorough",
    },
}

phases := []string{"Design", "Implementation", "Review"}

results, err := breeze.QuickCollab(agents, phases, "Build a REST API")
```

### Software development collaboration (SW + Testing teams)
```go
swTeam := []breeze.Agent{
    {Name: "Alex", Role: "Senior Engineer", Expertise: "Go development"},
    {Name: "Maria", Role: "Engineer", Expertise: "Database design"},
}

testTeam := []breeze.Agent{
    {Name: "David", Role: "QA Engineer", Expertise: "Testing"},
    {Name: "Sarah", Role: "Test Automation", Expertise: "Integration"},
}

results, err := breeze.TeamDevCollab(swTeam, testTeam, "Build a task manager CLI app")
```

### Quick team collaboration
```go
teams := []breeze.Team{
    {
        Name:        "Engineering",
        Description: "Development team",
        Agents: []breeze.Agent{
            {Name: "Alice", Role: "Developer", Expertise: "Backend"},
            {Name: "Bob", Role: "Developer", Expertise: "Frontend"},
        },
    },
    {
        Name:        "Design",
        Description: "UX/UI team",
        Agents: []breeze.Agent{
            {Name: "Carol", Role: "Designer", Expertise: "UI/UX"},
        },
    },
}

phases := []string{"Planning", "Design", "Implementation"}

results, err := breeze.QuickTeamCollab(teams, phases, "Build a mobile app")
```

### Advanced: Custom collaboration workflow
```go
agents := []breeze.Agent{
    {Name: "Alice", Role: "Analyst", Expertise: "Requirements"},
    {Name: "Bob", Role: "Developer", Expertise: "Implementation"},
}

phases := []breeze.Phase{
    {
        Name:           "Analysis",
        Description:    "Analyze requirements",
        PromptTemplate: "Provide detailed analysis for: {{.Prompt}}",
        IsParallel:     false,  // Sequential execution
        MaxConcurrency: 1,
    },
    {
        Name:           "Implementation",
        Description:    "Build the solution",
        PromptTemplate: "Implement the solution: {{.Prompt}}",
        IsParallel:     true,   // Parallel execution
        MaxConcurrency: 2,
    },
}

collab := breeze.NewCollaboration(agents, phases)

// Add progress callbacks (optional)
collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
    fmt.Printf("✅ Completed: %s\n", phaseName)
}

collab.OnAgentResponse = func(agentName, response string) {
    fmt.Printf("🤖 %s responded\n", agentName)
}

results, err := collab.Run("Build a web server")
```

### Save collaboration results
```go
results, _ := breeze.QuickCollab(agents, phases, "Project task")
collab := breeze.NewCollaboration(agents, phases)
collab.SaveResults(results, "collaboration_results.md")
```

---

## 🖥️ CLI Usage

### Simple query
```bash
./breeze "Explain quantum physics"
```

### Chat mode (with context)
```bash
./breeze chat "Hello, I need help with Go"
./breeze chat "Can you explain interfaces?"
```

### Code generation
```bash
./breeze code "Write a Go HTTP server"
```

### Clear conversation history
```bash
./breeze clear
```

---

## 📊 Types Reference

### Agent
```go
type Agent struct {
    Name        string  // Agent's name
    Role        string  // Job role (e.g., "Engineer", "Designer")
    Expertise   string  // Area of expertise
    Personality string  // Personality trait (e.g., "pragmatic", "creative")
}
```

### Phase
```go
type Phase struct {
    Name           string  // Phase name
    Description    string  // Phase description
    PromptTemplate string  // Template for agent prompts
    IsParallel     bool    // Execute agents in parallel?
    MaxConcurrency int     // Max concurrent agents
}
```

### Team
```go
type Team struct {
    Name        string   // Team name
    Description string   // Team description
    Agents      []Agent  // Team members
}
```

### Options
```go
type RequestOptions struct {
    Model   string    // LLM model to use
    Temp    float64   // Temperature (0.0-1.0)
    Stream  bool      // Enable streaming
    Context string    // Additional context
    Docs    []string  // Document file paths
    Concise bool      // Enable concise mode
}
```

---

## 🎯 Common Patterns

### Question & Answer
```go
answer := breeze.AI("What is recursion?")
```

### Code Review
```go
review := breeze.AI(
    "Review this code for issues",
    breeze.WithContext(codeSnippet),
    breeze.WithModel("codellama"),
)
```

### Document Analysis
```go
summary := breeze.AI(
    "What are the key findings?",
    breeze.WithDocs("research.pdf"),
)
```

### Multi-step conversation
```go
breeze.Chat("I'm building a web API")
breeze.Chat("What framework should I use?")
breeze.Chat("Show me an example with gorilla/mux")
response := breeze.Chat("How do I add authentication?")
```

### Parallel processing
```go
questions := []string{
    "Explain REST",
    "Explain GraphQL",
    "Explain gRPC",
}
answers := breeze.Batch(questions)
```

### Team brainstorming
```go
team := []breeze.Agent{
    {Name: "PM", Role: "Product Manager", Expertise: "User needs"},
    {Name: "Dev", Role: "Developer", Expertise: "Technical feasibility"},
    {Name: "Designer", Role: "UX Designer", Expertise: "User experience"},
}

results, _ := breeze.QuickCollab(
    team,
    []string{"Ideation", "Feasibility", "Design"},
    "Build a note-taking mobile app",
)
```

---

## 🔧 Setup & Requirements

### Prerequisites
```bash
# Install Ollama (auto-managed by Breeze)
# Download from: https://ollama.ai

# Build Breeze
go build ./cmd/breeze
```

### Preferred Models (auto-selected)
1. `gpt-oss` - General purpose
2. `codellama` - Code generation
3. `llama2` - General purpose
4. `mistral` - General purpose

Models are automatically downloaded if not available.

---

## 💡 Tips

1. **Use `WithConcise()` for quick answers** - Automatically enables streaming
2. **Use `Code()` for programming tasks** - Automatically selects codellama
3. **Use `Chat()` for conversations** - Maintains context across calls
4. **Use `Clear()` to reset context** - Start fresh conversations
5. **Team collaboration works best with 2-5 agents** - More agents = slower but diverse
6. **Documents are extracted automatically** - Supports PDF, DOCX, TXT
7. **Temperature affects creativity** - Lower (0.1) = consistent, Higher (0.9) = creative
8. **Streaming provides real-time feedback** - Great for long responses

---

## 📚 Examples Directory

Check out `/examples` for complete working examples:
- `team_development.go` - AI team collaboration demo
- `task_manager/` - Complete task manager app built by AI teams
- `finance_tracker.go` - Financial analysis system
- `composable_collaboration_design.go` - Advanced collaboration patterns
- And more...

---

## 🔗 Quick Links

- **GitHub**: https://github.com/yossideutsch1973/breeze
- **Full README**: [README.md](README.md)
- **License**: MIT

Troubleshooting
1. **Ollama must be running**: start with `ollama serve`.
2. **Model availability**: use `ollama list` to confirm the model is downloaded.
3. **Python wrapper**: use `pip install breeze-ai` to install the wrapper before calling from Python.
