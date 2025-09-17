package funcs

import (
	"fmt"

	"github.com/user/breeze"
)

// RunTeamCollabMinimal demonstrates a minimal team collaboration example
func RunTeamCollabMinimal() {
	team := []breeze.Agent{
		{Name: "Alice", Role: "Product Manager", Expertise: "Requirements Analysis"},
		{Name: "Bob", Role: "Engineer", Expertise: "Technical Feasibility"},
		{Name: "Carol", Role: "Designer", Expertise: "UX/UI"},
	}

	result, err := breeze.TeamDevCollab(team, nil, "Design a new mobile app for note-taking.")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", result)
}
