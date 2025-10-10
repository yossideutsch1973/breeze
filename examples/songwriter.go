package main

import (
	"fmt"
	"strings"

	"github.com/user/breeze"
)

func main() {
	fmt.Println("🎵 === AI Songwriter Collaboration === 🎵")
	fmt.Println("Using Breeze's new collaborative AI framework!\n")

	// The songwriting challenge
	challenge := `Create a funny, creative song about programmers debugging code. Make it humorous, relatable, and entertaining for developers. Include verses about common programming frustrations, debugging adventures, and the joy of finally fixing that elusive bug.`

	// Define collaborative agents
	agents := []breeze.Agent{
		{
			Name:        "Lyricist",
			Role:        "Creative Poet",
			Expertise:   "song lyrics and poetic structure",
			Personality: "playful and creative, always finding the humor in everyday situations",
		},
		{
			Name:        "Composer",
			Role:        "Music Theorist",
			Expertise:   "rhythm, rhyme, and musical structure",
			Personality: "melodic and structured, ensuring the song flows beautifully",
		},
		{
			Name:        "Critic",
			Role:        "Entertainment Expert",
			Expertise:   "humor and audience engagement",
			Personality: "witty and constructive, always improving the fun factor",
		},
	}

	// Define collaborative phases (limited to 3 as requested)
	phases := []breeze.Phase{
		{
			Name:           "Concept Development",
			Description:    "Brainstorming funny programmer scenarios and song structure",
			PromptTemplate: "Brainstorm funny scenarios about programmers debugging code. Think of humorous situations, relatable frustrations, and entertaining metaphors.",
			IsParallel:     true,
			MaxConcurrency: 3,
		},
		{
			Name:           "Lyrics Creation",
			Description:    "Writing the actual song lyrics with verses and chorus",
			PromptTemplate: "Write creative song lyrics about programmer debugging adventures. Include verses about coding frustrations, debugging triumphs, and funny programming metaphors.",
			IsParallel:     true,
			MaxConcurrency: 3,
		},
		{
			Name:           "Final Polish",
			Description:    "Refining and perfecting the complete song",
			PromptTemplate: "Take all the contributions and create a polished, complete song about programmers debugging code. Make it funny, engaging, and ready to perform.",
			IsParallel:     false, // Sequential for final coherence
			MaxConcurrency: 1,
		},
	}

	// Create and run collaboration
	collab := breeze.NewCollaboration(agents, phases)

	// Add progress callbacks for fun user experience
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("\n✅ Phase '%s' completed with %d creative contributions!\n", phaseName, len(results))
		fmt.Println(strings.Repeat("🎵", len(results)))
	}

	collab.OnAgentResponse = func(agentName, response string) {
		fmt.Printf("🎤 %s shared their creative input!\n", agentName)
	}

	// Run the collaboration
	results, err := collab.Run(challenge)
	if err != nil {
		fmt.Printf("❌ Collaboration failed: %v\n", err)
		return
	}

	// Save results automatically
	err = collab.SaveResults(results, "songwriter_collaboration.md")
	if err != nil {
		fmt.Printf("❌ Failed to save results: %v\n", err)
	} else {
		fmt.Println("\n💾 Results saved to: songwriter_collaboration.md")
	}

	// Display final song
	if finalPhase, exists := results["Final Polish"]; exists {
		if song, exists := finalPhase["Critic"]; exists {
			fmt.Println("\n" + strings.Repeat("═", 80))
			fmt.Println("🎵 FINAL SONG: PROGRAMMER'S DEBUGGING BLUES")
			fmt.Println("════════════════════════════════════════════════════════════════════════════════")
			fmt.Println(song)
			fmt.Println("════════════════════════════════════════════════════════════════════════════════")
		}
	}

	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("🎉 SONGWRITER COLLABORATION COMPLETE!")
	fmt.Println("Three AI artists have collaboratively created a funny song about programmers!")
	fmt.Println("This demonstrates the power of Breeze's collaborative AI framework! 🎵✨")

	// Clear conversation
	breeze.Clear()
}
