package funcs

import (
	"fmt"

	breeze "github.com/user/breeze"
)

// RunProductReviewTeam demonstrates a team review of a requirements document
func RunProductReviewTeam() {
	team := []breeze.Agent{
		{Name: "Alice", Role: "Product Manager", Expertise: "Requirements Analysis"},
		{Name: "Bob", Role: "Engineer", Expertise: "Technical Feasibility"},
		{Name: "Carol", Role: "Designer", Expertise: "UX/UI"},
	}

	result, err := breeze.TeamDevCollab(
		team,
		nil, // No test team
		"Review the attached product requirements and provide your expert feedback.",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\n\n--- Consensus Summary ---")
	fmt.Println(result)
}
