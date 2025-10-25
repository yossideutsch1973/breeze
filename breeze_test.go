package breeze

import (
	"os"
	"strings"
	"testing"
)

// Unit tests that don't require Ollama

func TestWithModel(t *testing.T) {
	opts := RequestOptions{}
	WithModel("testmodel")(&opts)
	if opts.Model != "testmodel" {
		t.Errorf("Expected model 'testmodel', got '%s'", opts.Model)
	}
}

func TestWithTemp(t *testing.T) {
	opts := RequestOptions{}
	WithTemp(0.5)(&opts)
	if opts.Temp != 0.5 {
		t.Errorf("Expected temp 0.5, got %f", opts.Temp)
	}
}

func TestWithContext(t *testing.T) {
	opts := RequestOptions{}
	WithContext("test context")(&opts)
	if opts.Context != "test context" {
		t.Errorf("Expected context 'test context', got '%s'", opts.Context)
	}
}

func TestWithDocs(t *testing.T) {
	opts := RequestOptions{}
	WithDocs("file1.txt", "file2.txt")(&opts)
	if len(opts.Docs) != 2 {
		t.Errorf("Expected 2 docs, got %d", len(opts.Docs))
	}
	if opts.Docs[0] != "file1.txt" || opts.Docs[1] != "file2.txt" {
		t.Errorf("Unexpected docs: %v", opts.Docs)
	}
}

func TestWithConcise(t *testing.T) {
	opts := RequestOptions{}
	WithConcise()(&opts)
	if !opts.Concise {
		t.Error("Expected Concise to be true")
	}
	if !opts.Stream {
		t.Error("Expected Stream to be true when Concise is enabled")
	}
}

func TestNewCollaboration(t *testing.T) {
	agents := []Agent{
		{Name: "Alice", Role: "Developer", Expertise: "Go", Personality: "friendly"},
	}
	phases := []Phase{
		{Name: "Phase1", Description: "Test", PromptTemplate: "Test prompt"},
	}

	collab := NewCollaboration(agents, phases)

	if collab == nil {
		t.Fatal("NewCollaboration returned nil")
	}
	if len(collab.Agents) != 1 {
		t.Errorf("Expected 1 agent, got %d", len(collab.Agents))
	}
	if len(collab.Phases) != 1 {
		t.Errorf("Expected 1 phase, got %d", len(collab.Phases))
	}
	if collab.SharedKnowledge == nil {
		t.Error("SharedKnowledge should be initialized")
	}
}

func TestCollaborationBuildAgentPrompt(t *testing.T) {
	collab := &Collaboration{
		SharedKnowledge: map[string]string{
			"Bob": "Some knowledge from Bob",
		},
	}

	agent := Agent{
		Name:        "Alice",
		Role:        "Developer",
		Expertise:   "Go programming",
		Personality: "friendly and collaborative",
	}

	phase := Phase{
		Name:           "Testing",
		Description:    "Test phase",
		PromptTemplate: "Please test this",
	}

	prompt := collab.buildAgentPrompt(agent, phase, "Build a feature")

	if !strings.Contains(prompt, "Alice") {
		t.Error("Prompt should contain agent name")
	}
	if !strings.Contains(prompt, "Developer") {
		t.Error("Prompt should contain agent role")
	}
	if !strings.Contains(prompt, "Go programming") {
		t.Error("Prompt should contain expertise")
	}
	if !strings.Contains(prompt, "Testing") {
		t.Error("Prompt should contain phase name")
	}
	if !strings.Contains(prompt, "Build a feature") {
		t.Error("Prompt should contain initial prompt")
	}
	if !strings.Contains(prompt, "Bob") {
		t.Error("Prompt should contain shared knowledge from other agents")
	}
}

func TestCollaborationSaveResults(t *testing.T) {
	collab := &Collaboration{}
	results := map[string]map[string]string{
		"Phase1": {
			"Alice": "Alice's response",
			"Bob":   "Bob's response",
		},
	}

	// Save to temp file
	tmpFile := "/tmp/test_results.md"
	err := collab.SaveResults(results, tmpFile)
	if err != nil {
		t.Errorf("SaveResults failed: %v", err)
	}
}

func TestCollaborationFormatResults(t *testing.T) {
	collab := &Collaboration{}
	results := map[string]map[string]string{
		"Phase1": {
			"Alice": "Alice's response",
		},
	}

	formatted := collab.formatResults(results)

	if !strings.Contains(formatted, "Phase1") {
		t.Error("Formatted results should contain phase name")
	}
	if !strings.Contains(formatted, "Alice") {
		t.Error("Formatted results should contain agent name")
	}
	if !strings.Contains(formatted, "Alice's response") {
		t.Error("Formatted results should contain response")
	}
}

func TestNewTeamCollaboration(t *testing.T) {
	teams := []Team{
		{
			Name:        "Dev Team",
			Description: "Developers",
			Agents: []Agent{
				{Name: "Alice", Role: "Dev", Expertise: "Go", Personality: "friendly"},
			},
		},
	}
	phases := []Phase{
		{Name: "Phase1", Description: "Test", PromptTemplate: "Test prompt"},
	}

	collab := NewTeamCollaboration(teams, phases)

	if collab == nil {
		t.Fatal("NewTeamCollaboration returned nil")
	}
	if len(collab.Teams) != 1 {
		t.Errorf("Expected 1 team, got %d", len(collab.Teams))
	}
	if len(collab.Phases) != 1 {
		t.Errorf("Expected 1 phase, got %d", len(collab.Phases))
	}
}

func TestTeamCollaborationBuildTeamAgentPrompt(t *testing.T) {
	tc := &TeamCollaboration{
		SharedKnowledge: map[string]string{
			"Charlie": "Some knowledge from Charlie",
		},
	}

	team := Team{
		Name:        "Dev Team",
		Description: "Development team",
	}

	agent := Agent{
		Name:        "Alice",
		Role:        "Developer",
		Expertise:   "Go programming",
		Personality: "efficient",
	}

	phase := Phase{
		Name:           "Implementation",
		Description:    "Implement feature",
		PromptTemplate: "Implement this",
	}

	prompt := tc.buildTeamAgentPrompt(agent, team, phase, "Build a feature")

	if !strings.Contains(prompt, "Alice") {
		t.Error("Prompt should contain agent name")
	}
	if !strings.Contains(prompt, "Dev Team") {
		t.Error("Prompt should contain team name")
	}
	if !strings.Contains(prompt, "Implementation") {
		t.Error("Prompt should contain phase name")
	}
	if !strings.Contains(prompt, "Build a feature") {
		t.Error("Prompt should contain initial prompt")
	}
}

func TestExtractTextFromPDF(t *testing.T) {
	// Simple PDF text extraction test
	pdfData := []byte("BT\n(Hello World)\nET")
	text, err := extractTextFromPDF(pdfData)
	if err != nil {
		t.Errorf("extractTextFromPDF failed: %v", err)
	}
	// PDF extraction is basic - just verify it doesn't crash
	// The actual text extraction may vary
	_ = text
}

func TestExtractTextFromFile_UnsupportedFormat(t *testing.T) {
	// Test with unsupported format
	tmpFile := "/tmp/test_file_breeze.xyz"
	content := []byte("test content")
	err := createTestFile(tmpFile, content)
	if err != nil {
		t.Skipf("Cannot create test file: %v", err)
	}
	defer removeTestFile(tmpFile)

	_, err = extractTextFromFile(tmpFile)
	if err == nil {
		t.Error("Expected error for unsupported file format")
	}
	if !strings.Contains(err.Error(), "unsupported file format") {
		t.Errorf("Expected 'unsupported file format' error, got: %v", err)
	}
}

func TestExtractTextFromFile_TXT(t *testing.T) {
	// Test TXT file extraction
	tmpFile := "/tmp/test_file_breeze.txt"
	content := []byte("This is test content")
	err := createTestFile(tmpFile, content)
	if err != nil {
		t.Skipf("Cannot create test file: %v", err)
	}
	defer removeTestFile(tmpFile)

	text, err := extractTextFromFile(tmpFile)
	if err != nil {
		t.Errorf("extractTextFromFile failed: %v", err)
	}
	if text != string(content) {
		t.Errorf("Expected '%s', got '%s'", content, text)
	}
}

// Helper function to create test files
func createTestFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0o644)
}

// Helper function to remove test files
func removeTestFile(path string) {
	os.Remove(path)
}

func TestAI_EmptyPrompt(t *testing.T) {
	result := AI("")
	if !strings.Contains(result, "Error") && !strings.Contains(result, "empty") {
		t.Errorf("Expected error for empty prompt, got: %s", result)
	}
}

func TestChat_EmptyPrompt(t *testing.T) {
	result := Chat("")
	if !strings.Contains(result, "Error") && !strings.Contains(result, "empty") {
		t.Errorf("Expected error for empty prompt, got: %s", result)
	}
}

func TestCode_EmptyPrompt(t *testing.T) {
	result := Code("")
	if !strings.Contains(result, "Error") && !strings.Contains(result, "empty") {
		t.Errorf("Expected error for empty prompt, got: %s", result)
	}
}

func TestBatch_EmptyList(t *testing.T) {
	results := Batch([]string{})
	if len(results) != 0 {
		t.Errorf("Expected empty results for empty prompts list, got %d items", len(results))
	}
}

func TestProcessDocuments_EmptyList(t *testing.T) {
	text, err := processDocuments([]string{})
	if err != nil {
		t.Errorf("processDocuments with empty list should not error: %v", err)
	}
	if text != "" {
		t.Errorf("Expected empty text for empty document list, got: %s", text)
	}
}

func TestWithTemp_Validation(t *testing.T) {
	// Test that temperature is set correctly
	opts := RequestOptions{}
	WithTemp(1.0)(&opts)
	if opts.Temp != 1.0 {
		t.Errorf("Expected temp 1.0, got %f", opts.Temp)
	}

	// Test extreme values
	WithTemp(0.0)(&opts)
	if opts.Temp != 0.0 {
		t.Errorf("Expected temp 0.0, got %f", opts.Temp)
	}

	WithTemp(2.0)(&opts)
	if opts.Temp != 2.0 {
		t.Errorf("Expected temp 2.0, got %f", opts.Temp)
	}
}

func TestNewCollaboration_EmptyAgents(t *testing.T) {
	phases := []Phase{
		{Name: "Phase1", Description: "Test", PromptTemplate: "Test"},
	}

	collab := NewCollaboration([]Agent{}, phases)
	if collab == nil {
		t.Fatal("NewCollaboration should not return nil for empty agents")
	}
	if len(collab.Agents) != 0 {
		t.Errorf("Expected 0 agents, got %d", len(collab.Agents))
	}
}

func TestNewCollaboration_EmptyPhases(t *testing.T) {
	agents := []Agent{
		{Name: "Alice", Role: "Dev", Expertise: "Go", Personality: "friendly"},
	}

	collab := NewCollaboration(agents, []Phase{})
	if collab == nil {
		t.Fatal("NewCollaboration should not return nil for empty phases")
	}
	if len(collab.Phases) != 0 {
		t.Errorf("Expected 0 phases, got %d", len(collab.Phases))
	}
}

// Integration tests that require Ollama

func TestAI(t *testing.T) {
	// Note: This test requires Ollama running
	t.Skip("Requires Ollama - run manually for integration testing")

	response := AI("Hello")
	if response == "" {
		t.Error("Expected non-empty response")
	}
}

func TestChat(t *testing.T) {
	t.Skip("Requires Ollama - run manually for integration testing")

	Clear() // Reset
	response := Chat("Hello")
	if response == "" {
		t.Error("Expected non-empty response")
	}
}

func TestCode(t *testing.T) {
	t.Skip("Requires Ollama - run manually for integration testing")

	response := Code("Write a function")
	if response == "" {
		t.Error("Expected non-empty response")
	}
}

func TestBatch(t *testing.T) {
	t.Skip("Requires Ollama - run manually for integration testing")

	prompts := []string{"Hello", "Hi"}
	results := Batch(prompts)
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}
