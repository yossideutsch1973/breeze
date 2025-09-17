package funcs

import (
	"fmt"
	"strings"

	breeze "github.com/user/breeze"
)

// RunSingleVsCollabComparison compares single LLM vs collaborative approach on symbolic integration
func RunSingleVsCollabComparison() {
	problem := `Compute the indefinite integral:

    ∫ x^4 * sin(x) / (1 + x^2)^2 dx

Show all steps, use integration by parts and substitution as needed.`

	fmt.Println("=== SINGLE LLM APPROACH ===")
	singleResult := breeze.AI(problem)
	fmt.Println(singleResult)

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("=== COLLABORATIVE APPROACH ===")

	agents := []breeze.Agent{
		{Name: "Planner", Role: "Strategy", Expertise: "Symbolic Math", Personality: "Analytical"},
		{Name: "Integrator1", Role: "First Integration by Parts", Expertise: "Calculus", Personality: "Methodical"},
		{Name: "Simplifier", Role: "Simplification", Expertise: "Algebra", Personality: "Detail-oriented"},
		{Name: "Integrator2", Role: "Second Integration by Parts", Expertise: "Calculus", Personality: "Thorough"},
		{Name: "Assembler", Role: "Final Assembly", Expertise: "Math Formatting", Personality: "Clear"},
	}

	phases := []breeze.Phase{
		{
			Name:           "Plan Solution",
			Description:    "Devise a high-level plan for solving the integral",
			PromptTemplate: "You are the planner. Devise a high-level plan for solving the integral, specifying which techniques to use and the order of steps.",
		},
		{
			Name:           "First Integration by Parts",
			Description:    "Carry out the first integration by parts as planned",
			PromptTemplate: "You are the first integrator. Carry out the first integration by parts as planned. Show all work.",
		},
		{
			Name:           "Simplify Result",
			Description:    "Simplify the result from the previous step",
			PromptTemplate: "You are the simplifier. Simplify the result from the previous step. Combine like terms and prepare for the next integration.",
		},
		{
			Name:           "Second Integration by Parts",
			Description:    "Perform the next integration by parts or substitution as needed",
			PromptTemplate: "You are the second integrator. Perform the next integration by parts or substitution as needed. Show all work.",
		},
		{
			Name:           "Assemble Final Answer",
			Description:    "Assemble the final answer, check correctness, and format the solution",
			PromptTemplate: "You are the assembler. Assemble the final answer, check correctness, and format the solution in clear LaTeX.",
		},
	}

	collab := breeze.NewCollaboration(agents, phases)
	results, err := collab.Run(problem)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for phase, res := range results {
		fmt.Printf("\n--- %s ---\n", phase)
		for agent, output := range res {
			fmt.Printf("%s: %s\n", agent, output)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("=== COMPARISON SUMMARY ===")
	fmt.Println("Single LLM: Likely to provide a basic solution, may skip steps or make errors")
	fmt.Println("Collaborative: Structured approach, peer review, detailed steps, higher accuracy")
}