// Package breeze provides ultra-simple local LLM interactions via Ollama.
//
// Key Features:
//   - Zero-configuration AI queries and chat
//   - Team collaboration framework for multi-agent workflows
//   - Streaming responses and batch processing
//   - Document processing (PDF, DOCX, TXT)
//   - Cross-platform support with auto Ollama management
//   - Functional options pattern for clean configuration
//
// Example usage:
//
//	// Simple AI query
//	response := breeze.AI("Explain quantum physics")
//
//	// Team collaboration
//	result := breeze.TeamDevCollab(swTeam, testTeam, project)
//
//	// Streaming with concise mode
//	response := breeze.AI("Explain AI", breeze.WithConcise())
package breeze

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// Breeze represents the AI client
type Breeze struct {
	model     string
	ollamaURL string
	messages  []Message
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Option is a functional option for configuring requests
type Option func(*RequestOptions)

// RequestOptions holds configuration for AI requests
type RequestOptions struct {
	Model   string
	Temp    float64
	Stream  bool
	Context string
	Docs    []string
	Concise bool
}

// WithModel sets the model for the request
func WithModel(model string) Option {
	return func(opts *RequestOptions) {
		opts.Model = model
	}
}

// WithTemp sets the temperature
func WithTemp(temp float64) Option {
	return func(opts *RequestOptions) {
		opts.Temp = temp
	}
}

// WithContext adds context to the prompt
func WithContext(context string) Option {
	return func(opts *RequestOptions) {
		opts.Context = context
	}
}

// WithDocs adds document files to be processed and included in context
func WithDocs(filePaths ...string) Option {
	return func(opts *RequestOptions) {
		opts.Docs = append(opts.Docs, filePaths...)
	}
}

// WithConcise enables concise responses with streaming output
func WithConcise() Option {
	return func(opts *RequestOptions) {
		opts.Concise = true
		opts.Stream = true // Concise mode always streams
	}
}

// preferredModels in order of preference
var preferredModels = []string{"gpt-oss", "codellama", "llama2", "mistral"}

// defaultClient is the global client
var defaultClient *Breeze

func init() {
	defaultClient = &Breeze{
		model:     "",
		ollamaURL: "http://localhost:11434",
		messages:  []Message{},
	}
	// Ensure Ollama is running
	ensureOllamaRunning()
	// Select best model
	defaultClient.model = selectBestModel()
}

// ensureOllamaRunning checks if Ollama is running and starts it if not
func ensureOllamaRunning() {
	// Try to connect to Ollama
	_, err := http.Get("http://localhost:11434/api/tags")
	if err != nil {
		fmt.Println("Ollama not detected. Starting Ollama...")
		cmd := exec.Command("ollama", "serve")
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Failed to start Ollama: %v\n", err)
			fmt.Println("Please install Ollama from https://ollama.ai")
			return
		}
		// Wait a bit for it to start
		time.Sleep(2 * time.Second)
		fmt.Println("Ollama started successfully.")
	}
}

// extractTextFromFile extracts text content from various file formats
func extractTextFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	// Determine file type by extension
	if strings.HasSuffix(strings.ToLower(filePath), ".txt") {
		return string(data), nil
	}

	if strings.HasSuffix(strings.ToLower(filePath), ".pdf") {
		return extractTextFromPDF(data)
	}

	if strings.HasSuffix(strings.ToLower(filePath), ".docx") {
		return extractTextFromDOCX(data)
	}

	return "", fmt.Errorf("unsupported file format: %s", filePath)
}

// extractTextFromPDF extracts text from PDF files
func extractTextFromPDF(data []byte) (string, error) {
	var text strings.Builder

	// Simple PDF text extraction - look for text objects between BT/ET
	content := string(data)
	lines := strings.Split(content, "\n")

	inTextObject := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "BT" {
			inTextObject = true
			continue
		}
		if line == "ET" {
			inTextObject = false
			continue
		}
		if inTextObject && strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			// Extract text from PDF text showing operator
			if len(line) > 2 {
				text.WriteString(line[1 : len(line)-1])
				text.WriteString(" ")
			}
		}
	}

	return strings.TrimSpace(text.String()), nil
}

// extractTextFromDOCX extracts text from DOCX files (ZIP archive with XML)
func extractTextFromDOCX(data []byte) (string, error) {
	// DOCX is a ZIP file containing document.xml
	zipReader, err := zip.NewReader(strings.NewReader(string(data)), int64(len(data)))
	if err != nil {
		return "", fmt.Errorf("failed to read DOCX as ZIP: %v", err)
	}

	// Find document.xml
	var docFile *zip.File
	for _, file := range zipReader.File {
		if file.Name == "word/document.xml" {
			docFile = file
			break
		}
	}

	if docFile == nil {
		return "", fmt.Errorf("document.xml not found in DOCX")
	}

	// Read document.xml
	rc, err := docFile.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open document.xml: %v", err)
	}
	defer rc.Close()

	xmlData, err := io.ReadAll(rc)
	if err != nil {
		return "", fmt.Errorf("failed to read document.xml: %v", err)
	}

	// Simple XML text extraction - look for text between <w:t> tags
	content := string(xmlData)
	var text strings.Builder

	// Find all text elements
	parts := strings.Split(content, "<w:t")
	for _, part := range parts[1:] { // Skip first part before first <w:t>
		if endIdx := strings.Index(part, "</w:t>"); endIdx != -1 {
			textContent := part[:endIdx]
			// Remove XML entities and clean up
			textContent = strings.ReplaceAll(textContent, "&amp;", "&")
			textContent = strings.ReplaceAll(textContent, "&lt;", "<")
			textContent = strings.ReplaceAll(textContent, "&gt;", ">")
			textContent = strings.ReplaceAll(textContent, "&quot;", "\"")
			textContent = strings.ReplaceAll(textContent, "&apos;", "'")
			text.WriteString(textContent)
			text.WriteString(" ")
		}
	}

	return strings.TrimSpace(text.String()), nil
}

// processDocuments extracts text from all provided document files
func processDocuments(filePaths []string) (string, error) {
	if len(filePaths) == 0 {
		return "", nil
	}

	var allText strings.Builder
	allText.WriteString("\n--- Document Context ---\n")

	for _, filePath := range filePaths {
		text, err := extractTextFromFile(filePath)
		if err != nil {
			return "", fmt.Errorf("error processing %s: %v", filePath, err)
		}

		allText.WriteString(fmt.Sprintf("\nFile: %s\n", filePath))
		allText.WriteString(text)
		allText.WriteString("\n\n")
	}

	allText.WriteString("--- End Document Context ---\n")
	return allText.String(), nil
}

// ai generates a response for a single prompt
func AI(prompt string, opts ...Option) string {
	options := RequestOptions{
		Model:  defaultClient.model,
		Temp:   0.7,
		Stream: false,
	}
	for _, opt := range opts {
		opt(&options)
	}

	// Process documents if provided
	if len(options.Docs) > 0 {
		docText, err := processDocuments(options.Docs)
		if err != nil {
			return fmt.Sprintf("Error processing documents: %v", err)
		}
		if options.Context != "" {
			options.Context = options.Context + "\n\n" + docText
		} else {
			options.Context = docText
		}
	}

	if options.Context != "" {
		prompt = options.Context + "\n\n" + prompt
	}

	// Add concise instruction if enabled
	if options.Concise {
		prompt = "Be extremely concise and brief in your response. " + prompt
	}

	req := map[string]interface{}{
		"model":  options.Model,
		"prompt": prompt,
		"stream": options.Stream,
	}
	if options.Temp != 0.7 {
		req["options"] = map[string]interface{}{
			"temperature": options.Temp,
		}
	}

	jsonData, _ := json.Marshal(req)
	resp, err := http.Post(defaultClient.ollamaURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	defer resp.Body.Close()

	// Handle streaming response for concise mode
	if options.Concise && options.Stream {
		var fullResponse strings.Builder
		decoder := json.NewDecoder(resp.Body)
		for {
			var chunk map[string]interface{}
			if err := decoder.Decode(&chunk); err != nil {
				break
			}
			if token, ok := chunk["response"].(string); ok {
				fmt.Print(token) // Stream to stdout
				fullResponse.WriteString(token)
			}
			if done, ok := chunk["done"].(bool); ok && done {
				break
			}
		}
		fmt.Println() // New line after streaming
		return fullResponse.String()
	}

	// Regular non-streaming response
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result["response"].(string)
}

// chat maintains conversation context
func Chat(prompt string, opts ...Option) string {
	options := RequestOptions{
		Model:  defaultClient.model,
		Temp:   0.7,
		Stream: false,
	}
	for _, opt := range opts {
		opt(&options)
	}

	// Process documents if provided
	userMessage := prompt
	if len(options.Docs) > 0 {
		docText, err := processDocuments(options.Docs)
		if err != nil {
			return fmt.Sprintf("Error processing documents: %v", err)
		}
		userMessage = docText + "\n\n" + prompt
	}

	// Add concise instruction if enabled
	if options.Concise {
		userMessage = "Be extremely concise and brief in your response. " + userMessage
	}

	defaultClient.messages = append(defaultClient.messages, Message{Role: "user", Content: userMessage})

	req := map[string]interface{}{
		"model":    options.Model,
		"messages": defaultClient.messages,
		"stream":   options.Stream,
	}
	if options.Temp != 0.7 {
		req["options"] = map[string]interface{}{
			"temperature": options.Temp,
		}
	}

	jsonData, _ := json.Marshal(req)
	resp, err := http.Post(defaultClient.ollamaURL+"/api/chat", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	defer resp.Body.Close()

	// Handle streaming response for concise chat mode
	if options.Concise && options.Stream {
		var fullResponse strings.Builder
		decoder := json.NewDecoder(resp.Body)
		for {
			var chunk map[string]interface{}
			if err := decoder.Decode(&chunk); err != nil {
				break
			}
			if message, ok := chunk["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					fmt.Print(content) // Stream to stdout
					fullResponse.WriteString(content)
				}
			}
			if done, ok := chunk["done"].(bool); ok && done {
				break
			}
		}
		fmt.Println() // New line after streaming
		defaultClient.messages = append(defaultClient.messages, Message{Role: "assistant", Content: fullResponse.String()})
		return fullResponse.String()
	}

	// Regular non-streaming chat response
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	response := result["message"].(map[string]interface{})["content"].(string)
	defaultClient.messages = append(defaultClient.messages, Message{Role: "assistant", Content: response})

	return response
}

// code is optimized for code generation
func Code(prompt string, opts ...Option) string {
	// Use codellama if available, fallback to default
	model := "codellama"
	if !isModelAvailable(model) {
		model = defaultClient.model
	}
	opts = append(opts, WithModel(model))
	return AI("Write code for: "+prompt, opts...)
}

// isModelAvailable checks if a model is available
func isModelAvailable(model string) bool {
	resp, err := http.Get(defaultClient.ollamaURL + "/api/tags")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	models := result["models"].([]interface{})
	for _, m := range models {
		if m.(map[string]interface{})["name"].(string) == model {
			return true
		}
	}
	return false
}

// selectBestModel selects the best available model, pulling if necessary
func selectBestModel() string {
	for _, model := range preferredModels {
		if isModelAvailable(model) {
			return model
		}
	}
	// Pull the first preferred model
	pullModel(preferredModels[0])
	return preferredModels[0]
}

// pullModel pulls a model using Ollama
func pullModel(model string) {
	fmt.Printf("Pulling model %s...\n", model)
	cmd := exec.Command("ollama", "pull", model)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to pull model %s: %v\n", model, err)
	} else {
		fmt.Printf("Model %s pulled successfully.\n", model)
	}
}

// clear resets the conversation
func Clear() {
	defaultClient.messages = []Message{}
}

// StreamFunc is the callback for streaming
type StreamFunc func(token string)

// Stream streams the response
func Stream(prompt string, fn StreamFunc, opts ...Option) {
	options := RequestOptions{
		Model:  defaultClient.model,
		Temp:   0.7,
		Stream: true,
	}
	for _, opt := range opts {
		opt(&options)
	}

	req := map[string]interface{}{
		"model":  options.Model,
		"prompt": prompt,
		"stream": true,
	}
	if options.Temp != 0.7 {
		req["options"] = map[string]interface{}{
			"temperature": options.Temp,
		}
	}

	jsonData, _ := json.Marshal(req)
	resp, err := http.Post(defaultClient.ollamaURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fn(fmt.Sprintf("Error: %v", err))
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var chunk map[string]interface{}
		if err := decoder.Decode(&chunk); err != nil {
			break
		}
		if token, ok := chunk["response"].(string); ok {
			fn(token)
		}
		if done, ok := chunk["done"].(bool); ok && done {
			break
		}
	}
}

// Batch processes multiple prompts concurrently
func Batch(prompts []string, opts ...Option) []string {
	results := make([]string, len(prompts))
	for i, prompt := range prompts {
		go func(idx int, p string) {
			results[idx] = AI(p, opts...)
		}(i, prompt)
	}
	// Wait for all to complete (simple implementation)
	time.Sleep(5 * time.Second) // TODO: better synchronization
	return results
}

// ===== COLLABORATIVE AI FRAMEWORK =====

// Agent represents a collaborative AI agent with specific expertise
type Agent struct {
	Name        string
	Role        string
	Expertise   string
	Personality string
}

// Phase represents a collaborative phase with specific instructions
type Phase struct {
	Name           string
	Description    string
	PromptTemplate string
	IsParallel     bool
	MaxConcurrency int
}

// Collaboration manages multi-agent collaborative workflows
type Collaboration struct {
	Agents          []Agent
	Phases          []Phase
	SharedKnowledge map[string]string
	OnPhaseComplete func(phaseName string, results map[string]string)
	OnAgentResponse func(agentName, response string)
}

// NewCollaboration creates a new collaborative workflow
func NewCollaboration(agents []Agent, phases []Phase) *Collaboration {
	return &Collaboration{
		Agents:          agents,
		Phases:          phases,
		SharedKnowledge: make(map[string]string),
	}
}

// Run executes the entire collaborative workflow
func (c *Collaboration) Run(initialPrompt string) (map[string]map[string]string, error) {
	results := make(map[string]map[string]string)

	for _, phase := range c.Phases {
		fmt.Printf("\n🔄 PHASE: %s\n%s\n", phase.Name, phase.Description)

		phaseResults := c.runPhase(phase, initialPrompt)
		results[phase.Name] = phaseResults

		// Update shared knowledge
		for agentName, response := range phaseResults {
			c.SharedKnowledge[agentName] = response
		}

		if c.OnPhaseComplete != nil {
			c.OnPhaseComplete(phase.Name, phaseResults)
		}
	}

	return results, nil
}

// runPhase executes a single collaborative phase
func (c *Collaboration) runPhase(phase Phase, initialPrompt string) map[string]string {
	results := make(map[string]string)

	if phase.IsParallel {
		return c.runParallelPhase(phase, initialPrompt)
	}

	// Sequential execution
	for _, agent := range c.Agents {
		prompt := c.buildAgentPrompt(agent, phase, initialPrompt)
		response := AI(prompt, WithConcise())

		results[agent.Name] = response

		if c.OnAgentResponse != nil {
			c.OnAgentResponse(agent.Name, response)
		}
	}

	return results
}

// runParallelPhase executes agents in parallel
func (c *Collaboration) runParallelPhase(phase Phase, initialPrompt string) map[string]string {
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	maxConcurrency := phase.MaxConcurrency
	if maxConcurrency <= 0 {
		maxConcurrency = len(c.Agents)
	}

	semaphore := make(chan struct{}, maxConcurrency)

	for _, agent := range c.Agents {
		wg.Add(1)
		go func(agent Agent) {
			defer wg.Done()

			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			prompt := c.buildAgentPrompt(agent, phase, initialPrompt)
			response := AI(prompt, WithConcise())

			mu.Lock()
			results[agent.Name] = response
			mu.Unlock()

			if c.OnAgentResponse != nil {
				c.OnAgentResponse(agent.Name, response)
			}
		}(agent)
	}

	wg.Wait()
	return results
}

// buildAgentPrompt constructs the prompt for an agent in a specific phase
func (c *Collaboration) buildAgentPrompt(agent Agent, phase Phase, initialPrompt string) string {
	prompt := fmt.Sprintf("You are %s, %s with expertise in %s. %s\n\n",
		agent.Name, agent.Role, agent.Expertise, agent.Personality)

	prompt += fmt.Sprintf("PHASE: %s\n%s\n\n", phase.Name, phase.Description)
	prompt += fmt.Sprintf("ORIGINAL CHALLENGE: %s\n\n", initialPrompt)

	// Add shared knowledge from other agents
	if len(c.SharedKnowledge) > 0 {
		prompt += "COLLABORATIVE INSIGHTS FROM OTHER EXPERTS:\n"
		for name, knowledge := range c.SharedKnowledge {
			if name != agent.Name {
				prompt += fmt.Sprintf("🔹 %s: %s\n", name, knowledge)
			}
		}
		prompt += "\n"
	}

	prompt += phase.PromptTemplate
	return prompt
}

// ===== CONVENIENCE FUNCTIONS =====

// QuickCollab creates a simple collaborative workflow
func QuickCollab(agents []Agent, phases []string, challenge string) (map[string]map[string]string, error) {
	phaseConfigs := make([]Phase, len(phases))
	for i, phaseName := range phases {
		phaseConfigs[i] = Phase{
			Name:           phaseName,
			Description:    fmt.Sprintf("Working on %s", phaseName),
			PromptTemplate: fmt.Sprintf("Provide your expert contribution to %s. Be specific and actionable.", phaseName),
			IsParallel:     true,
			MaxConcurrency: 4,
		}
	}

	collab := NewCollaboration(agents, phaseConfigs)

	// Add nice progress reporting
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("✅ Completed phase: %s (%d contributions)\n", phaseName, len(results))
	}

	collab.OnAgentResponse = func(agentName, response string) {
		fmt.Printf("🤖 %s contributed to the discussion\n", agentName)
	}

	return collab.Run(challenge)
}

// SaveResults automatically saves collaboration results to files
func (c *Collaboration) SaveResults(results map[string]map[string]string, filename string) error {
	content := c.formatResults(results)
	return os.WriteFile(filename, []byte(content), 0644)
}

// formatResults creates a nicely formatted summary of all results
func (c *Collaboration) formatResults(results map[string]map[string]string) string {
	output := "# Collaborative AI Results\n\n"
	output += fmt.Sprintf("Generated on: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	for phaseName, phaseResults := range results {
		output += fmt.Sprintf("## Phase: %s\n\n", phaseName)
		for agentName, response := range phaseResults {
			output += fmt.Sprintf("### %s\n%s\n\n", agentName, response)
		}
	}

	return output
}

// ===== ENHANCED TEAM COLLABORATION FRAMEWORK =====

// Team represents a group of agents working together
type Team struct {
	Name        string
	Description string
	Agents      []Agent
}

// TeamCollaboration manages multi-team collaborative workflows with automatic phase alternation
type TeamCollaboration struct {
	Teams           []Team
	Phases          []Phase
	SharedKnowledge map[string]string
	OnPhaseComplete func(phaseName string, results map[string]string)
	OnAgentResponse func(agentName, response string)
	OnTeamComplete  func(teamName string, results map[string]string)
}

// NewTeamCollaboration creates a new team-based collaborative workflow
func NewTeamCollaboration(teams []Team, phases []Phase) *TeamCollaboration {
	return &TeamCollaboration{
		Teams:           teams,
		Phases:          phases,
		SharedKnowledge: make(map[string]string),
	}
}

// Run executes the team collaborative workflow with automatic team alternation
func (tc *TeamCollaboration) Run(initialPrompt string) (map[string]map[string]string, error) {
	results := make(map[string]map[string]string)

	for _, phase := range tc.Phases {
		fmt.Printf("\n🔄 PHASE: %s\n%s\n", phase.Name, phase.Description)

		// Execute phase for all teams
		phaseResults := tc.runTeamPhase(phase, initialPrompt)
		results[phase.Name] = phaseResults

		// Update shared knowledge
		for agentName, response := range phaseResults {
			tc.SharedKnowledge[agentName] = response
		}

		if tc.OnPhaseComplete != nil {
			tc.OnPhaseComplete(phase.Name, phaseResults)
		}
	}

	return results, nil
}

// runTeamPhase executes a phase across all teams
func (tc *TeamCollaboration) runTeamPhase(phase Phase, initialPrompt string) map[string]string {
	results := make(map[string]string)

	for _, team := range tc.Teams {
		fmt.Printf("👥 %s team working...\n", team.Name)

		teamResults := tc.runTeamAgents(team, phase, initialPrompt)

		// Merge team results
		for agentName, response := range teamResults {
			results[agentName] = response
		}

		if tc.OnTeamComplete != nil {
			tc.OnTeamComplete(team.Name, teamResults)
		}
	}

	return results
}

// runTeamAgents executes all agents in a team for a phase
func (tc *TeamCollaboration) runTeamAgents(team Team, phase Phase, initialPrompt string) map[string]string {
	results := make(map[string]string)

	if phase.IsParallel {
		return tc.runParallelTeamAgents(team, phase, initialPrompt)
	}

	// Sequential execution within team
	for _, agent := range team.Agents {
		prompt := tc.buildTeamAgentPrompt(agent, team, phase, initialPrompt)
		response := AI(prompt, WithConcise())

		results[agent.Name] = response

		if tc.OnAgentResponse != nil {
			tc.OnAgentResponse(agent.Name, response)
		}
	}

	return results
}

// runParallelTeamAgents executes team agents in parallel
func (tc *TeamCollaboration) runParallelTeamAgents(team Team, phase Phase, initialPrompt string) map[string]string {
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	maxConcurrency := phase.MaxConcurrency
	if maxConcurrency <= 0 {
		maxConcurrency = len(team.Agents)
	}

	semaphore := make(chan struct{}, maxConcurrency)

	for _, agent := range team.Agents {
		wg.Add(1)
		go func(agent Agent) {
			defer wg.Done()

			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			prompt := tc.buildTeamAgentPrompt(agent, team, phase, initialPrompt)
			response := AI(prompt, WithConcise())

			mu.Lock()
			results[agent.Name] = response
			mu.Unlock()

			if tc.OnAgentResponse != nil {
				tc.OnAgentResponse(agent.Name, response)
			}
		}(agent)
	}

	wg.Wait()
	return results
}

// buildTeamAgentPrompt constructs the prompt for a team agent
func (tc *TeamCollaboration) buildTeamAgentPrompt(agent Agent, team Team, phase Phase, initialPrompt string) string {
	prompt := fmt.Sprintf("You are %s, %s with expertise in %s. %s\n\n",
		agent.Name, agent.Role, agent.Expertise, agent.Personality)

	prompt += fmt.Sprintf("TEAM: %s - %s\n", team.Name, team.Description)
	prompt += fmt.Sprintf("PHASE: %s\n%s\n\n", phase.Name, phase.Description)
	prompt += fmt.Sprintf("ORIGINAL CHALLENGE: %s\n\n", initialPrompt)

	// Add shared knowledge from other teams/agents
	if len(tc.SharedKnowledge) > 0 {
		prompt += "COLLABORATIVE INSIGHTS FROM OTHER TEAMS:\n"
		for name, knowledge := range tc.SharedKnowledge {
			if name != agent.Name {
				prompt += fmt.Sprintf("🔹 %s: %s\n", name, knowledge)
			}
		}
		prompt += "\n"
	}

	prompt += phase.PromptTemplate
	return prompt
}

// ===== CONVENIENCE FUNCTIONS FOR TEAM COLLABORATIONS =====

// TeamDevCollab creates a development collaboration between SW and Testing teams
func TeamDevCollab(swTeam, testTeam []Agent, project string) (map[string]map[string]string, error) {
	teams := []Team{
		{
			Name:        "SW Engineering",
			Description: "Software development and implementation",
			Agents:      swTeam,
		},
		{
			Name:        "Testing",
			Description: "Quality assurance and validation",
			Agents:      testTeam,
		},
	}

	phases := []Phase{
		{
			Name:           "Requirements Analysis",
			Description:    "Both teams analyze requirements and plan approach",
			PromptTemplate: "Analyze the project requirements and provide your team's perspective on implementation approach and quality considerations.",
			IsParallel:     true,
			MaxConcurrency: 4,
		},
		{
			Name:           "SW Implementation",
			Description:    "SW team creates initial implementation",
			PromptTemplate: "As the software engineering team, create the initial implementation focusing on core functionality and clean architecture.",
			IsParallel:     true,
			MaxConcurrency: 2,
		},
		{
			Name:           "Testing & Feedback",
			Description:    "Testing team evaluates implementation and provides feedback",
			PromptTemplate: "As the testing team, thoroughly evaluate the implementation. Test all features, identify bugs, edge cases, and suggest improvements.",
			IsParallel:     true,
			MaxConcurrency: 2,
		},
		{
			Name:           "SW Refinement",
			Description:    "SW team addresses testing feedback",
			PromptTemplate: "As the software engineering team, address the testing feedback. Fix identified issues and improve the implementation.",
			IsParallel:     true,
			MaxConcurrency: 2,
		},
		{
			Name:           "Final Testing",
			Description:    "Testing team validates final version",
			PromptTemplate: "As the testing team, perform final validation. Ensure all issues are resolved and the product is production-ready.",
			IsParallel:     true,
			MaxConcurrency: 2,
		},
		{
			Name:           "Final Polish",
			Description:    "SW team adds final touches",
			PromptTemplate: "As the software engineering team, add final polish and optimizations. Ensure the code is clean, well-documented, and ready for production.",
			IsParallel:     false,
			MaxConcurrency: 1,
		},
	}

	collab := NewTeamCollaboration(teams, phases)

	// Add nice progress reporting
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("✅ Phase '%s' completed with %d contributions!\n", phaseName, len(results))
	}

	collab.OnTeamComplete = func(teamName string, results map[string]string) {
		if teamName == "SW Engineering" {
			fmt.Println("💻 SW team delivered implementation!")
		} else {
			fmt.Println("🧪 Testing team provided critical feedback!")
		}
	}

	collab.OnAgentResponse = func(agentName, response string) {
		fmt.Printf("🤖 %s contributed!\n", agentName)
	}

	return collab.Run(project)
}

// QuickTeamCollab creates a simple team-based collaboration
func QuickTeamCollab(teams []Team, phases []string, challenge string) (map[string]map[string]string, error) {
	phaseConfigs := make([]Phase, len(phases))
	for i, phaseName := range phases {
		phaseConfigs[i] = Phase{
			Name:           phaseName,
			Description:    fmt.Sprintf("Working on %s", phaseName),
			PromptTemplate: fmt.Sprintf("Provide your team's expert contribution to %s. Be specific and actionable.", phaseName),
			IsParallel:     true,
			MaxConcurrency: 4,
		}
	}

	collab := NewTeamCollaboration(teams, phaseConfigs)

	// Add nice progress reporting
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("✅ Completed phase: %s (%d contributions)\n", phaseName, len(results))
	}

	collab.OnTeamComplete = func(teamName string, results map[string]string) {
		fmt.Printf("👥 %s team completed their work\n", teamName)
	}

	collab.OnAgentResponse = func(agentName, response string) {
		fmt.Printf("🤖 %s contributed to the discussion\n", agentName)
	}

	return collab.Run(challenge)
}
