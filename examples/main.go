package main

import (
	"fmt"
	"os"

	ex "github.com/user/breeze/internal/examples/funcs"
)

var registry = map[string]struct {
	Desc string
	Run  func()
}{
	// "team_collab_minimal": {"Minimal team collaboration", ex.RunTeamCollabMinimal},
	// "ai_doc_summarize":    {"AI document summarization", ex.RunAIDocSummarize},
	// "product_review_team": {"Team review of requirements document", ex.RunProductReviewTeam},
	// "composable_collab_demo": {"Composable collaboration methods demo", ex.RunComposableCollaborationDemo},
	// "debate_competition": {"Debate competition with judges", ex.RunDebateCompetition},
	"coding_team_benchmark":       {"Team coding benchmark with peer review and iteration", ex.RunCodingTeamBenchmark},
	"symbolic_integration_collab": {"Collaborative symbolic integration of a complex expression", ex.RunSymbolicIntegrationCollab},
	"single_vs_collab_comparison": {"Single LLM vs Collaborative approach comparison", ex.RunSingleVsCollabComparison},
	"webapp_truck_simulation":     {"Collaborative web app development for truck simulation", ex.RunWebAppTruckSimulation},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run examples/main.go <example>")
		fmt.Println("Available examples:")
		for k, v := range registry {
			fmt.Printf("  %-20s %s\n", k, v.Desc)
		}
		os.Exit(1)
	}
	name := os.Args[1]
	e, ok := registry[name]
	if !ok {
		fmt.Printf("Unknown example: %s\n", name)
		os.Exit(1)
	}
	e.Run()
}
