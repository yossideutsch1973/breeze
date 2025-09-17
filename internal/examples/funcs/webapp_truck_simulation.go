package funcs

import (
	"fmt"
	"strings"

	breeze "github.com/user/breeze"
)

// RunWebAppTruckSimulation demonstrates collaborative web app development for multi-chain truck movement simulation
func RunWebAppTruckSimulation() {
	project := `Develop a single page web application that simulates multi-chain truck movement.

Requirements:
- Interactive map showing truck routes across multiple supply chains
- Real-time truck position updates and status tracking
- Chain management (different logistics companies/routes)
- Performance metrics and analytics dashboard
- Responsive design for mobile and desktop
- Modern web technologies (HTML5, CSS3, JavaScript/React)

The app should visualize trucks moving along predefined routes, show chain interconnections, and provide insights into logistics efficiency.`

	// Define specialized agents for web development
	agents := []breeze.Agent{
		{Name: "Alice", Role: "UX/UI Designer", Expertise: "User experience, responsive design, visualization", Personality: "creative and user-focused"},
		{Name: "Bob", Role: "Frontend Developer", Expertise: "React, JavaScript, CSS, interactive maps", Personality: "technical and detail-oriented"},
		{Name: "Carol", Role: "Data Architect", Expertise: "data modeling, APIs, real-time updates", Personality: "systematic and analytical"},
		{Name: "Dave", Role: "DevOps Engineer", Expertise: "deployment, performance, scalability", Personality: "practical and efficiency-focused"},
		{Name: "Eve", Role: "Product Manager", Expertise: "requirements, user stories, project coordination", Personality: "organized and strategic"},
	}

	// Define development phases
	phases := []breeze.Phase{
		{
			Name:           "Requirements & Planning",
			Description:    "Analyze requirements, create user stories, and plan the technical architecture",
			PromptTemplate: "Analyze the truck simulation requirements. Create detailed user stories, technical specifications, and implementation plan for your expertise area.",
			IsParallel:     true,
			MaxConcurrency: 5,
		},
		{
			Name:           "System Design",
			Description:    "Design the overall system architecture, data models, and component structure",
			PromptTemplate: "Based on the requirements, design the system architecture for your domain. Specify APIs, data flow, component structure, and integration points.",
			IsParallel:     true,
			MaxConcurrency: 5,
		},
		{
			Name:           "Implementation Plan",
			Description:    "Create detailed implementation roadmap with code structure and key components",
			PromptTemplate: "Create a detailed implementation plan for your area. Include code structure, key algorithms, libraries/frameworks, and integration approach.",
			IsParallel:     false,
		},
		{
			Name:           "Code Generation",
			Description:    "Generate initial code, components, and configuration files",
			PromptTemplate: "Generate the initial code, components, and configuration files for your domain. Include working examples and starter templates.",
			IsParallel:     true,
			MaxConcurrency: 3,
		},
		{
			Name:           "Integration & Testing",
			Description:    "Plan integration strategy, testing approach, and deployment pipeline",
			PromptTemplate: "Define how your components integrate with others. Create testing strategy, deployment plan, and quality assurance approach.",
			IsParallel:     false,
		},
	}

	// Create collaboration with progress tracking
	collab := breeze.NewCollaboration(agents, phases)

	// Set up progress callbacks
	collab.OnPhaseComplete = func(phaseName string, results map[string]string) {
		fmt.Printf("\n🎯 PHASE COMPLETED: %s\n", phaseName)
		fmt.Println("Summary of deliverables:")
		for agent, result := range results {
			// Show first 100 chars of each result as summary
			summary := result
			if len(summary) > 100 {
				summary = summary[:100] + "..."
			}
			fmt.Printf("  %s: %s\n", agent, summary)
		}
	}

	fmt.Println("=== COLLABORATIVE WEB APP DEVELOPMENT ===")
	fmt.Println("Project: Multi-Chain Truck Movement Simulation")
	fmt.Println("Team: UX Designer, Frontend Dev, Data Architect, DevOps, Product Manager")
	fmt.Println(fmt.Sprintf("Phases: %d collaborative development phases", len(phases)))

	results, err := collab.Run(project)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("=== FINAL DELIVERABLES ===")

	for phase, res := range results {
		fmt.Printf("\n--- %s ---\n", phase)
		for agent, output := range res {
			fmt.Printf("\n**%s's Contribution:**\n%s\n", agent, output)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("=== PROJECT SUMMARY ===")
	fmt.Println("✅ Requirements analyzed and user stories created")
	fmt.Println("✅ System architecture designed with component specifications")
	fmt.Println("✅ Implementation roadmap with technical details")
	fmt.Println("✅ Code generated for frontend, backend, and infrastructure")
	fmt.Println("✅ Integration strategy and deployment pipeline defined")
	fmt.Println("\nReady for development team to begin implementation!")
}
