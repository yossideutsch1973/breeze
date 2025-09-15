package breeze

import (
	"testing"
)

func TestAI(t *testing.T) {
	// Note: This test requires Ollama running
	// For CI, mock or skip
	t.Skip("Requires Ollama")

	response := AI("Hello")
	if response == "" {
		t.Error("Expected non-empty response")
	}
}

func TestChat(t *testing.T) {
	t.Skip("Requires Ollama")

	Clear() // Reset
	response := Chat("Hello")
	if response == "" {
		t.Error("Expected non-empty response")
	}
}

func TestCode(t *testing.T) {
	t.Skip("Requires Ollama")

	response := Code("Write a function")
	if response == "" {
		t.Error("Expected non-empty response")
	}
}

func TestBatch(t *testing.T) {
	t.Skip("Requires Ollama")

	prompts := []string{"Hello", "Hi"}
	results := Batch(prompts)
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}
