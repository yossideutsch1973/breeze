package funcs

import (
	"fmt"

	"github.com/user/breeze"
)

func main() {
	// Define a team of agents
	agents := []breeze.Agent{
		{Name: "Alice", Role: "Senior Developer", Expertise: "Go programming", Personality: "pragmatic and efficient"},
		{Name: "Bob", Role: "Architect", Expertise: "system design", Personality: "thorough and methodical"},
		{Name: "Carol", Role: "QA Engineer", Expertise: "testing strategies", Personality: "detail-oriented and critical"},
	}

	challenge := "Design a scalable microservice for user authentication"

	fmt.Println("=== Testing New Collaboration Methods ===\n")

	// Test 1: Quick Parallel Collaboration
	fmt.Println("🚀 Test 1: Parallel Collaboration")
	results, err := breeze.QuickParallel(agents, challenge, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\nResults:")
	for agent, response := range results {
		fmt.Printf("- %s: %s\n", agent, response[:100]+"...")
	}

	// Test 2: Peer Review Collaboration
	fmt.Println("\n🔍 Test 2: Peer Review Collaboration")
	reviewResults, err := breeze.QuickPeerReview(agents, challenge, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nPeer review results: %d items\n", len(reviewResults))

	// Test 3: Consensus Building
	fmt.Println("\n🎯 Test 3: Consensus Building")
	consensus, err := breeze.QuickConsensus(agents, challenge, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nConsensus: %s\n", consensus[:150]+"...")

	// Test 4: Custom Collaboration Method
	fmt.Println("\n⚙️  Test 4: Custom Collaboration Method")

	// Create a custom collaboration method that does sequential with breaks
	customMethod := func(agents []breeze.Agent, collab *breeze.Collaboration, phase breeze.Phase, initialPrompt string) map[string]string {
		results := make(map[string]string)

		fmt.Println("  Custom method: Sequential with pauses...")
		for i, agent := range agents {
			if i > 0 {
				fmt.Printf("  Pausing before %s...\n", agent.Name)
				// In real use, might add delays or intermediate processing
			}

			prompt := collab.BuildAgentPrompt(agent, phase, initialPrompt)
			response := breeze.AI(prompt, breeze.WithConcise())
			results[agent.Name] = response

			fmt.Printf("  ✓ %s completed\n", agent.Name)
		}

		return results
	}

	// Use custom method
	phase := breeze.NewPhase(
		"Custom Analysis",
		"Custom sequential analysis with pauses",
		"Provide your expert analysis:",
		breeze.WithMethod(customMethod),
	)

	collab := breeze.NewCollaboration(agents, []breeze.Phase{phase})
	customResults, err := collab.Run(challenge)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nCustom collaboration completed with %d results\n", len(customResults[phase.Name]))

	fmt.Println("\n✅ All collaboration method tests completed!")
}
