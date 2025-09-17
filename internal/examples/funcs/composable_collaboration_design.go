// Enhanced Collaboration Methods Design for Breeze
// This file demonstrates the proposed composable collaboration system

package funcs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/user/breeze"
)

// ===== ENHANCED COLLABORATION SYSTEM DESIGN =====

// CollaborationMethod defines how agents work together in a phase
// It's a pure function that takes agents, collaboration context, phase, and prompt
// and returns the results of their collaboration
type CollaborationMethod func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string

// Enhanced Phase struct with collaboration method
type EnhancedPhase struct {
	Name           string
	Description    string
	PromptTemplate string
	Method         CollaborationMethod // Replaces IsParallel flag
}

// ===== BUILT-IN COLLABORATION METHODS =====

// Sequential executes agents one after another, each building on previous work
func Sequential() CollaborationMethod {
	return func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string {
		results := make(map[string]string)

		for _, agent := range agents {
			// Build context-aware prompt that includes previous agents' work
			prompt := buildAgentPrompt(agent, phase, initialPrompt, collab.SharedKnowledge)
			response := breeze.AI(prompt, breeze.WithConcise())

			results[agent.Name] = response

			// Add to shared knowledge for next agent
			collab.SharedKnowledge[agent.Name] = response

			fmt.Printf("✓ %s completed sequential work\n", agent.Name)
		}

		return results
	}
}

// Parallel executes all agents simultaneously
func Parallel(maxConcurrency int) CollaborationMethod {
	return func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string {
		results := make(map[string]string)
		var wg sync.WaitGroup
		var mu sync.Mutex

		if maxConcurrency <= 0 {
			maxConcurrency = len(agents)
		}

		semaphore := make(chan struct{}, maxConcurrency)

		for _, agent := range agents {
			wg.Add(1)
			go func(agent breeze.Agent) {
				defer wg.Done()
				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				prompt := buildAgentPrompt(agent, phase, initialPrompt, collab.SharedKnowledge)
				response := breeze.AI(prompt, breeze.WithConcise())

				mu.Lock()
				results[agent.Name] = response
				mu.Unlock()

				fmt.Printf("✓ %s completed parallel work\n", agent.Name)
			}(agent)
		}

		wg.Wait()
		return results
	}
}

// PeerReview executes agents in parallel, then has each review others' work
func PeerReview(maxConcurrency int) CollaborationMethod {
	return func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string {
		// Phase 1: Initial work in parallel
		fmt.Println("  🚀 Phase 1: Initial work")
		initialResults := Parallel(maxConcurrency)(agents, collab, phase, initialPrompt)

		// Phase 2: Peer review
		fmt.Println("  🔍 Phase 2: Peer review")
		reviews := make(map[string]string)
		var wg sync.WaitGroup
		var mu sync.Mutex

		for _, reviewer := range agents {
			wg.Add(1)
			go func(reviewer breeze.Agent) {
				defer wg.Done()

				reviewPrompt := fmt.Sprintf("You are %s reviewing peer work. Original challenge: %s\n\nPEER CONTRIBUTIONS:\n",
					reviewer.Name, initialPrompt)

				for agentName, work := range initialResults {
					if agentName != reviewer.Name {
						reviewPrompt += fmt.Sprintf("- %s: %s\n", agentName, work)
					}
				}

				reviewPrompt += "\nProvide constructive feedback, strengths, weaknesses, and suggestions:"
				review := breeze.AI(reviewPrompt, breeze.WithConcise())

				mu.Lock()
				reviews[reviewer.Name+"_review"] = review
				mu.Unlock()

				fmt.Printf("✓ %s completed peer review\n", reviewer.Name)
			}(reviewer)
		}

		wg.Wait()

		// Combine initial work and reviews
		for k, v := range reviews {
			initialResults[k] = v
		}

		return initialResults
	}
}

// Consensus executes agents in parallel, then synthesizes a consensus view
func Consensus(maxConcurrency int) CollaborationMethod {
	return func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string {
		// Phase 1: Get individual perspectives
		fmt.Println("  🤖 Phase 1: Individual perspectives")
		individualResults := Parallel(maxConcurrency)(agents, collab, phase, initialPrompt)

		// Phase 2: Synthesize consensus
		fmt.Println("  🎯 Phase 2: Building consensus")
		consensusPrompt := fmt.Sprintf("Challenge: %s\n\nEXPERT OPINIONS:\n", initialPrompt)
		for agentName, opinion := range individualResults {
			consensusPrompt += fmt.Sprintf("- %s: %s\n", agentName, opinion)
		}
		consensusPrompt += "\nSynthesize these expert opinions into a unified consensus:"

		consensus := breeze.AI(consensusPrompt, breeze.WithConcise())

		// Return both individual results and consensus
		results := make(map[string]string)
		for k, v := range individualResults {
			results[k] = v
		}
		results["CONSENSUS"] = consensus

		fmt.Println("✓ Consensus reached")
		return results
	}
}

// DebateStyle creates opposing positions and synthesizes resolution
func DebateStyle(rounds int) CollaborationMethod {
	return func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string {
		results := make(map[string]string)

		for round := 1; round <= rounds; round++ {
			fmt.Printf("  🗣️  Debate Round %d\n", round)

			for i, agent := range agents {
				position := "advocate"
				if i%2 == 1 {
					position = "challenge"
				}

				prompt := fmt.Sprintf("You are %s in a structured debate. Your position: %s the proposal.\n",
					agent.Name, position)
				prompt += fmt.Sprintf("Challenge: %s\n\n", initialPrompt)

				if round > 1 {
					prompt += "PREVIOUS DEBATE POINTS:\n"
					for k, v := range results {
						if strings.Contains(k, fmt.Sprintf("round_%d", round-1)) {
							prompt += fmt.Sprintf("- %s\n", v)
						}
					}
				}

				prompt += fmt.Sprintf("Provide a strong %s argument:", position)

				response := breeze.AI(prompt, breeze.WithConcise())
				results[fmt.Sprintf("%s_%s_round_%d", agent.Name, position, round)] = response

				fmt.Printf("    ✓ %s (%s position)\n", agent.Name, position)
			}
		}

		return results
	}
}

// ===== CONVENIENCE FUNCTIONS =====

// WithMethod sets the collaboration method for a phase
func WithMethod(method CollaborationMethod) func(*EnhancedPhase) {
	return func(p *EnhancedPhase) {
		p.Method = method
	}
}

// NewPhase creates a phase with collaboration method
func NewPhase(name, description, promptTemplate string, options ...func(*EnhancedPhase)) EnhancedPhase {
	phase := EnhancedPhase{
		Name:           name,
		Description:    description,
		PromptTemplate: promptTemplate,
		Method:         Sequential(), // Default to sequential
	}

	for _, option := range options {
		option(&phase)
	}

	return phase
}

// Helper function to build agent prompts (would be exported from breeze package)
func buildAgentPrompt(agent breeze.Agent, phase breeze.Phase, initialPrompt string, sharedKnowledge map[string]string) string {
	prompt := fmt.Sprintf("You are %s, %s with expertise in %s. %s\n\n",
		agent.Name, agent.Role, agent.Expertise, agent.Personality)

	prompt += fmt.Sprintf("CHALLENGE: %s\n\n", initialPrompt)

	if len(sharedKnowledge) > 0 {
		prompt += "COLLABORATIVE INSIGHTS:\n"
		for name, knowledge := range sharedKnowledge {
			if name != agent.Name {
				prompt += fmt.Sprintf("🔹 %s: %s\n", name, knowledge)
			}
		}
		prompt += "\n"
	}

	prompt += phase.PromptTemplate
	return prompt
}

// ===== EXAMPLE USAGE =====

// RunComposableCollaborationDemo demonstrates composable collaboration methods in action
func RunComposableCollaborationDemo() {
	agents := []breeze.Agent{
		{Name: "Alice", Role: "Senior Developer", Expertise: "Go programming", Personality: "pragmatic"},
		{Name: "Bob", Role: "Architect", Expertise: "system design", Personality: "thorough"},
		{Name: "Carol", Role: "QA Engineer", Expertise: "testing", Personality: "detail-oriented"},
	}

	fmt.Println("=== Demonstrating Composable Collaboration Methods ===")

	// Example 1: Peer Review Collaboration
	phase1 := NewPhase(
		"System Design",
		"Design a scalable authentication system",
		"Provide your expert design approach:",
		WithMethod(PeerReview(2)),
	)

	// Example 2: Consensus Building
	phase2 := NewPhase(
		"Technology Stack",
		"Choose optimal technology stack",
		"Recommend technology choices:",
		WithMethod(Consensus(3)),
	)

	// Example 3: Structured Debate
	phase3 := NewPhase(
		"Architecture Decision",
		"Microservices vs Monolith",
		"Present your architectural position:",
		WithMethod(DebateStyle(2)),
	)

	phases := []EnhancedPhase{phase1, phase2, phase3}
	for i, phase := range phases {
		fmt.Printf("\n--- Phase %d: %s ---\n", i+1, phase.Name)
		collab := &breeze.Collaboration{SharedKnowledge: make(map[string]string)}
		results := phase.Method(agents, collab, breeze.Phase{
			Name:           phase.Name,
			Description:    phase.Description,
			PromptTemplate: phase.PromptTemplate,
		}, phase.PromptTemplate)
		for k, v := range results {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
}
