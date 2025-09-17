package funcs

import (
	"fmt"
	"strings"

	"github.com/user/breeze"
)

// RunDebateCompetition runs a debate between two agents and judges the result
func RunDebateCompetition() {
	topic := "Should AI be allowed to make medical diagnoses without human oversight?"
	rounds := 3

	debaters := []breeze.Agent{
		{Name: "Debater A", Role: "Proponent", Expertise: "AI Ethics", Personality: "logical and persuasive"},
		{Name: "Debater B", Role: "Opponent", Expertise: "Medical Safety", Personality: "cautious and critical"},
	}
	judges := []breeze.Agent{
		{Name: "Judge 1", Role: "Ethics Professor", Expertise: "AI & Society"},
		{Name: "Judge 2", Role: "Medical Doctor", Expertise: "Patient Safety"},
	}

	transcript := []string{}
	lastStatement := fmt.Sprintf("Debate topic: %s\nDebater A, please open with your argument.", topic)

	for i := 0; i < rounds; i++ {
		for _, debater := range debaters {
			prompt := fmt.Sprintf("You are %s. The topic is: %s\nPrevious statement: %s\nRespond concisely with your argument.", debater.Name, topic, lastStatement)
			response := breeze.AI(prompt, breeze.WithConcise())
			transcript = append(transcript, fmt.Sprintf("%s: %s", debater.Name, response))
			lastStatement = response
		}
	}

	// Judges review the transcript and score
	judgeVerdicts := make(map[string]string)
	for _, judge := range judges {
		prompt := fmt.Sprintf("You are %s. Here is the debate transcript:\n%s\n\nAssign a score (0-10) to each debater and provide a one-sentence verdict.", judge.Name, strings.Join(transcript, "\n"))
		verdict := breeze.AI(prompt, breeze.WithConcise())
		judgeVerdicts[judge.Name] = verdict
		// (Optional: parse score from verdict if needed)
	}

	fmt.Println("\n=== Debate Transcript ===")
	for _, line := range transcript {
		fmt.Println(line)
	}
	fmt.Println("\n=== Judges' Verdicts ===")
	for judge, verdict := range judgeVerdicts {
		fmt.Printf("%s: %s\n", judge, verdict)
	}
}
