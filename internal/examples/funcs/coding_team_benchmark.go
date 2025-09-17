package funcs

import (
	"fmt"

	"github.com/user/breeze"
)

// RunCodingTeamBenchmark demonstrates 3 coders collaborating on a coding benchmark problem with peer review and iteration
func RunCodingTeamBenchmark() {
	problem := `LeetCode 2: Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.`

	coders := []breeze.Agent{
		{Name: "Alice", Role: "Senior Go Developer", Expertise: "algorithms, Go", Personality: "pragmatic and concise"},
		{Name: "Bob", Role: "Systems Engineer", Expertise: "data structures, C++", Personality: "thorough and analytical"},
		{Name: "Carol", Role: "Python Specialist", Expertise: "competitive programming", Personality: "creative and efficient"},
	}

	// Round 1: Each coder submits an initial solution
	fmt.Println("\n=== Round 1: Initial Solutions ===")
	solutions := make(map[string]string)
	for _, coder := range coders {
		prompt := fmt.Sprintf("You are %s. Solve the following problem in your preferred language.\nProblem: %s\nReturn only the code, no explanation.", coder.Name, problem)
		solutions[coder.Name] = breeze.Code(prompt)
		fmt.Printf("\n%s's solution:\n%s\n", coder.Name, solutions[coder.Name])
	}

	// Round 2: Peer review (each coder reviews another's code)
	fmt.Println("\n=== Round 2: Peer Review ===")
	reviews := make(map[string]string)
	for i, coder := range coders {
		reviewee := coders[(i+1)%len(coders)]
		prompt := fmt.Sprintf("You are %s. Review the following code written by %s for the problem: %s\nCode:\n%s\nProvide a concise review: strengths, weaknesses, and suggestions for improvement.", coder.Name, reviewee.Name, problem, solutions[reviewee.Name])
		reviews[coder.Name] = breeze.AI(prompt, breeze.WithConcise())
		fmt.Printf("\n%s reviews %s:\n%s\n", coder.Name, reviewee.Name, reviews[coder.Name])
	}

	// Round 3: Iteration (each coder improves their code based on peer feedback)
	fmt.Println("\n=== Round 3: Iterated Solutions ===")
	improved := make(map[string]string)
	for _, coder := range coders {
		review := reviews[coder.Name]
		prompt := fmt.Sprintf("You are %s. Here is your original code for the problem: %s\n%s\nPeer review feedback: %s\nImprove your code based on the feedback. Return only the improved code, no explanation.", coder.Name, problem, solutions[coder.Name], review)
		improved[coder.Name] = breeze.Code(prompt)
		fmt.Printf("\n%s's improved solution:\n%s\n", coder.Name, improved[coder.Name])
	}

	// Judges: Evaluate the final solutions
	judges := []breeze.Agent{
		{Name: "Judge 1", Role: "Algorithm Expert", Expertise: "code quality, correctness"},
		{Name: "Judge 2", Role: "Performance Engineer", Expertise: "efficiency, style"},
	}
	fmt.Println("\n=== Judges' Evaluation ===")
	for _, judge := range judges {
		for _, coder := range coders {
			prompt := fmt.Sprintf("You are %s. Evaluate the following solution to the problem: %s\nCode:\n%s\nScore the code from 0-10 and provide a one-sentence justification.", judge.Name, problem, improved[coder.Name])
			verdict := breeze.AI(prompt, breeze.WithConcise())
			fmt.Printf("%s on %s: %s\n", judge.Name, coder.Name, verdict)
		}
	}
}
