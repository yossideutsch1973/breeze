package main

import (
	"fmt"
	"strings"

	"github.com/user/breeze"
)

func main() {
	fmt.Println("🚀 === AI Startup Founders Collaboration === 🚀")
	fmt.Println("Using Breeze's new collaborative AI framework!\n")

	// The startup challenge
	challenge := "Design and implement an AI-powered delivery route optimization system for a food delivery service. The system must handle 50 delivery locations with time windows (e.g., 'deliver between 12:00-13:00'), vehicle capacity constraints (max 25kg per vehicle), and traffic patterns. Optimize for minimum total delivery time while respecting all constraints. Provide a complete algorithmic solution with pseudocode, complexity analysis, and implementation considerations."

	// Define collaborative agents
	agents := []breeze.Agent{
		{
			Name:        "Alex Chen",
			Role:        "CTO",
			Expertise:   "technical architecture and scalability",
			Personality: "technical and pragmatic, always focused on building robust systems",
		},
		{
			Name:        "Sarah Johnson",
			Role:        "CEO",
			Expertise:   "business strategy and market analysis",
			Personality: "strategic and visionary, always thinking about market opportunities",
		},
		{
			Name:        "Marcus Rodriguez",
			Role:        "CPO",
			Expertise:   "product design and user experience",
			Personality: "user-centric and creative, always designing for the customer",
		},
		{
			Name:        "Dr. Emily Watson",
			Role:        "Chief Scientist",
			Expertise:   "AI/ML innovation and research",
			Personality: "innovative and research-driven, always pushing technological boundaries",
		},
	}

	// Define collaborative phases (limited to 3 as requested)
	phases := []breeze.Phase{
		{
			Name:           "Market Analysis",
			Description:    "Analyzing the market opportunity and customer needs",
			PromptTemplate: "Analyze the market opportunity for this delivery route optimization system. Consider target customers, competitive landscape, and key value propositions.",
			IsParallel:     true,
			MaxConcurrency: 4,
		},
		{
			Name:           "Technical Design",
			Description:    "Designing the technical architecture and algorithms",
			PromptTemplate: "Design the technical architecture for this delivery optimization system. Focus on algorithms, data structures, scalability, and implementation approach.",
			IsParallel:     true,
			MaxConcurrency: 4,
		},
		{
			Name:           "Business Strategy",
			Description:    "Developing business model and go-to-market strategy",
			PromptTemplate: "Develop a comprehensive business strategy including revenue model, market positioning, and growth plan for this delivery optimization startup.",
			IsParallel:     false, // Sequential for final coherence
			MaxConcurrency: 1,
		},
	}

	// Create and run collaboration
	collab := breeze.NewCollaboration(agents, phases)

	// Add progress callbacks for fun user experience
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("\n✅ Phase '%s' completed with %d founder insights!\n", phaseName, len(results))
		fmt.Println(strings.Repeat("🚀", len(results)))
	}

	collab.OnAgentResponse = func(agentName, response string) {
		fmt.Printf("👔 %s shared their expertise!\n", agentName)
	}

	// Run the collaboration
	results, err := collab.Run(challenge)
	if err != nil {
		fmt.Printf("❌ Collaboration failed: %v\n", err)
		return
	}

	// Save results automatically
	err = collab.SaveResults(results, "startup_founders_collaboration.md")
	if err != nil {
		fmt.Printf("❌ Failed to save results: %v\n", err)
	} else {
		fmt.Println("\n💾 Results saved to: startup_founders_collaboration.md")
	}

	// Display final business strategy
	if finalPhase, exists := results["Business Strategy"]; exists {
		if strategy, exists := finalPhase["Sarah Johnson"]; exists {
			fmt.Println("\n" + strings.Repeat("═", 60))
			fmt.Println("📋 FINAL STARTUP BUSINESS STRATEGY")
			fmt.Println("════════════════════════════════════════════════════════════════════════════")
			fmt.Println(strategy)
			fmt.Println("════════════════════════════════════════════════════════════════════════════")
		}
	}

	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("🎊 STARTUP LAUNCH COMPLETE!")
	fmt.Println("Four AI founders have collaboratively built a complete startup in minutes!")
	fmt.Println("This demonstrates the power of Breeze's collaborative AI framework! 🚀✨")

	// Clear conversation
	breeze.Clear()
}
