package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/user/breeze"
)

func main() {
	fmt.Println("=== Team-Based Software Development Collaboration ===")
	fmt.Println("Using Breeze's enhanced team collaboration framework!\n")

	// Define SW Engineering team
	swTeam := []breeze.Agent{
		{
			Name:        "Alex Chen",
			Role:        "Senior Software Engineer",
			Expertise:   "Go development and system design",
			Personality: "pragmatic and detail-oriented, focuses on clean architecture",
		},
		{
			Name:        "Maria Rodriguez",
			Role:        "Software Engineer",
			Expertise:   "data structures and file I/O",
			Personality: "thorough and methodical, ensures robust implementation",
		},
	}

	// Define Testing team
	testTeam := []breeze.Agent{
		{
			Name:        "David Kim",
			Role:        "QA Engineer",
			Expertise:   "functional testing and edge cases",
			Personality: "critical and thorough, finds potential issues",
		},
		{
			Name:        "Sarah Johnson",
			Role:        "Test Automation Engineer",
			Expertise:   "integration testing and validation",
			Personality: "systematic and precise, ensures complete functionality",
		},
	}

	// The development challenge
	project := `Create a complete Go program that implements a task management system with:
- Add tasks with priority (high/medium/low)
- List tasks by priority
- Mark tasks as completed
- Remove completed tasks
- Save/load tasks to/from JSON file
Make it a command-line application with proper error handling.`

	// Use the enhanced TeamDevCollab function - much more concise!
	results, err := breeze.TeamDevCollab(swTeam, testTeam, project)
	if err != nil {
		fmt.Printf("❌ Development failed: %v\n", err)
		return
	}

	// Extract the final implementation
	var finalCode string
	if finalPhase, exists := results["Final Polish"]; exists {
		// Get the final polished version from SW team
		for agentName, code := range finalPhase {
			if agentName == "Alex Chen" || agentName == "Maria Rodriguez" {
				finalCode = code
				break
			}
		}
	}

	if finalCode == "" {
		fmt.Println("❌ Could not extract final implementation")
		return
	}

	// Save the final program
	err = os.WriteFile("task_manager.go", []byte(finalCode), 0644)
	if err != nil {
		fmt.Printf("❌ Failed to save program: %v\n", err)
		return
	}

	fmt.Printf("✅ Final program saved to: task_manager.go\n")
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("🎊 TEAM DEVELOPMENT COMPLETE!")
	fmt.Println("Two teams have collaboratively built a complete task management system!")
	fmt.Println("The final program is saved and ready to run!")
	fmt.Println("This demonstrates the power of Breeze's enhanced team collaboration framework! 🚀✨")

	// Clear conversation
	breeze.Clear()
}
