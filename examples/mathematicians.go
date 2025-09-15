package main

import (
	"fmt"
	"strings"

	"github.com/user/breeze"
)

func main() {
	fmt.Println("🧮 === Mathematical Collaboration === 🧮")
	fmt.Println("Using Breeze's new collaborative AI framework!\n")

	// The challenging mathematical problem (Putnam-style)
	problem := `Find all continuous functions f: [0,1] → [0,1] such that f(f(x)) = x for all x ∈ [0,1].

This is a functional equation problem requiring multiple mathematical approaches. Consider:
- Fixed points and iterations
- Continuity properties
- Possible forms of such functions
- Rigorous proof techniques`

	// Define collaborative agents
	agents := []breeze.Agent{
		{
			Name:       "Dr. Elena Vasquez",
			Role:       "Algebraist",
			Expertise:  "algebraic structures and equation solving",
			Personality: "rigorous and systematic, always seeking complete classifications",
		},
		{
			Name:       "Prof. Marcus Chen",
			Role:       "Analyst",
			Expertise:  "real and complex analysis",
			Personality: "deeply theoretical, passionate about continuity and convergence",
		},
		{
			Name:       "Dr. Sofia Patel",
			Role:       "Topologist",
			Expertise:  "topological spaces and continuity",
			Personality: "intuitive and geometric, always thinking about spaces and mappings",
		},
		{
			Name:       "Prof. David Kim",
			Role:       "Number Theorist",
			Expertise:  "Diophantine equations and modular arithmetic",
			Personality: "creative and pattern-seeking, always looking for elegant solutions",
		},
	}

	// Define collaborative phases (limited to 3 phases as requested)
	phases := []breeze.Phase{
		{
			Name:           "Initial Analysis",
			Description:    "Analyzing the functional equation from different mathematical perspectives",
			PromptTemplate: "Analyze this functional equation from your mathematical domain. Consider fixed points, continuity properties, and possible function forms.",
			IsParallel:     true,
			MaxConcurrency: 4,
		},
		{
			Name:           "Technical Development",
			Description:    "Developing rigorous solution approaches and proofs",
			PromptTemplate: "Develop a technical approach to solve this functional equation. Consider proofs, counterexamples, and domain-specific techniques.",
			IsParallel:     true,
			MaxConcurrency: 4,
		},
		{
			Name:           "Final Synthesis",
			Description:    "Synthesizing all contributions into a complete mathematical solution",
			PromptTemplate: `Synthesize all mathematical contributions into a complete, rigorous proof for this functional equation.

Create a comprehensive mathematical solution including:
- Statement of the problem
- Analysis of possible function forms
- Rigorous proof of the solution
- Verification and edge cases
- Mathematical rigor and completeness

Present this as a formal mathematical proof suitable for publication.`,
			IsParallel:     false, // Sequential for coherence
			MaxConcurrency: 1,
		},
	}

	// Create and run collaboration
	collab := breeze.NewCollaboration(agents, phases)

	// Add progress callbacks for fun user experience
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("\n✅ Phase '%s' completed with %d mathematical insights!\n", phaseName, len(results))
		fmt.Println(strings.Repeat("🧮", len(results)))
	}

	collab.OnAgentResponse = func(agentName, response string) {
		fmt.Printf("🤖 %s shared their mathematical expertise!\n", agentName)
	}

	// Run the collaboration
	results, err := collab.Run(problem)
	if err != nil {
		fmt.Printf("❌ Collaboration failed: %v\n", err)
		return
	}

	// Save results automatically
	err = collab.SaveResults(results, "mathematical_collaboration.md")
	if err != nil {
		fmt.Printf("❌ Failed to save results: %v\n", err)
	} else {
		fmt.Println("\n💾 Results saved to: mathematical_collaboration.md")
	}

	// Display final synthesis
	if finalPhase, exists := results["Final Synthesis"]; exists {
		if synthesis, exists := finalPhase["Prof. David Kim"]; exists {
			fmt.Println("\n" + strings.Repeat("═", 80))
			fmt.Println("📋 FINAL MATHEMATICAL PROOF")
			fmt.Println("════════════════════════════════════════════════════════════════════════════════")
			fmt.Println(synthesis)
			fmt.Println("════════════════════════════════════════════════════════════════════════════════")
		}
	}

	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("🎊 MATHEMATICAL COLLABORATION COMPLETE!")
	fmt.Println("Four mathematicians have collaboratively solved a complex functional equation!")
	fmt.Println("This demonstrates the power of Breeze's collaborative AI framework! 🧮✨")

	// Clear conversation
	breeze.Clear()
}