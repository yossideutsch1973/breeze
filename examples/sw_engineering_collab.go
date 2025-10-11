// SW Engineering Team Collaboration with Peer Review
// Demonstrates a proper peer review system where agents:
// 1. Each do their own work (implementation)
// 2. Review others' work with scoring
// 3. Merge reviews based on scores
// 4. Only final report goes to file (no console spam)
// 5. Progress bar shows advancement

package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/user/breeze"
)

// Review represents a peer review with scoring
type Review struct {
	Reviewer    string
	Reviewee    string
	Score       int // 1-10 scale
	Comments    string
	Strengths   []string
	Weaknesses  []string
	Suggestions []string
}

// WorkItem represents a piece of work by an agent
type WorkItem struct {
	AgentName string
	Work      string
	Reviews   []Review
}

// SWEngineeringCollab manages SW engineering team collaboration with peer review
type SWEngineeringCollab struct {
	Agents     []breeze.Agent
	Challenge  string
	WorkItems  map[string]*WorkItem
	OnProgress func(current, total int, phase string)
	mu         sync.RWMutex
}

// NewSWEngineeringCollab creates a new SW engineering collaboration
func NewSWEngineeringCollab(agents []breeze.Agent, challenge string) *SWEngineeringCollab {
	return &SWEngineeringCollab{
		Agents:    agents,
		Challenge: challenge,
		WorkItems: make(map[string]*WorkItem),
	}
}

// Run executes the complete SW engineering workflow
func (sec *SWEngineeringCollab) Run() error {
	totalSteps := len(sec.Agents)*2 + len(sec.Agents)*(len(sec.Agents)-1) + 2 // work + reviews + merge + final
	currentStep := 0

	// Phase 1: Each agent does their work
	fmt.Println("Starting software engineering collaboration")
	fmt.Println("📋 Phase 1: Individual Implementation")
	sec.updateProgress(currentStep, totalSteps, "Individual Work")

	for _, agent := range sec.Agents {
		prompt := sec.buildWorkPrompt(agent)
		work := breeze.AI(prompt, breeze.WithConcise())

		sec.mu.Lock()
		sec.WorkItems[agent.Name] = &WorkItem{
			AgentName: agent.Name,
			Work:      work,
			Reviews:   []Review{},
		}
		sec.mu.Unlock()

		currentStep++
		sec.updateProgress(currentStep, totalSteps, fmt.Sprintf("Work: %s", agent.Name))
	}

	// Phase 2: Peer review (each agent reviews others' work)
	fmt.Println("\n📋 Phase 2: Peer Review Process")
	for _, reviewer := range sec.Agents {
		for _, reviewee := range sec.Agents {
			if reviewer.Name != reviewee.Name {
				review := sec.conductPeerReview(reviewer, reviewee)
				sec.mu.Lock()
				if item, exists := sec.WorkItems[reviewee.Name]; exists {
					item.Reviews = append(item.Reviews, review)
				}
				sec.mu.Unlock()

				currentStep++
				sec.updateProgress(currentStep, totalSteps, fmt.Sprintf("Review: %s → %s", reviewer.Name, reviewee.Name))
			}
		}
	}

	// Phase 3: Merge reviews and create final assessment
	fmt.Println("\n📋 Phase 3: Review Consolidation")
	finalAssessment := sec.consolidateReviews()

	currentStep++
	sec.updateProgress(currentStep, totalSteps, "Consolidating Reviews")

	// Phase 4: Generate final report
	fmt.Println("\n📋 Phase 4: Final Report Generation")
	report := sec.generateFinalReport(finalAssessment)

	currentStep++
	sec.updateProgress(currentStep, totalSteps, "Generating Final Report")

	// Save final report to file
	err := sec.saveReportToFile(report)
	if err != nil {
		return fmt.Errorf("failed to save report: %v", err)
	}

	fmt.Printf("\nCollaboration complete. Report saved to: sw_team_report_%s.md\n", time.Now().Format("2006-01-02_15-04-05"))
	return nil
}

// buildWorkPrompt creates the work prompt for an agent
func (sec *SWEngineeringCollab) buildWorkPrompt(agent breeze.Agent) string {
	return fmt.Sprintf(`You are %s, a %s with expertise in %s. %s

CHALLENGE: %s

Your task is to provide a comprehensive solution to this challenge. Focus on:
1. Technical implementation details
2. Architecture considerations
3. Code structure and patterns
4. Potential challenges and solutions
5. Best practices and standards

Provide a detailed, actionable implementation plan.`, agent.Name, agent.Role, agent.Expertise, agent.Personality, sec.Challenge)
}

// conductPeerReview has one agent review another's work
func (sec *SWEngineeringCollab) conductPeerReview(reviewer, reviewee breeze.Agent) Review {
	sec.mu.RLock()
	workItem := sec.WorkItems[reviewee.Name]
	sec.mu.RUnlock()

	prompt := fmt.Sprintf(`You are %s, a %s conducting a peer code review.

REVIEWEE: %s (%s)
WORK TO REVIEW:
%s

Please provide a structured peer review with:
1. Overall Score (1-10, where 10 is excellent)
2. Key Strengths (list 2-3)
3. Areas for Improvement (list 1-3)
4. Specific Suggestions (list 2-3)
5. Overall Assessment

Be constructive, specific, and professional in your feedback.`, reviewer.Name, reviewer.Role, reviewee.Name, reviewee.Role, workItem.Work)

	reviewText := breeze.AI(prompt, breeze.WithConcise())

	// Parse the review (simplified parsing)
	review := Review{
		Reviewer:    reviewer.Name,
		Reviewee:    reviewee.Name,
		Score:       sec.extractScore(reviewText),
		Comments:    reviewText,
		Strengths:   sec.extractListItems(reviewText, "strengths", "strength"),
		Weaknesses:  sec.extractListItems(reviewText, "weaknesses", "improvement"),
		Suggestions: sec.extractListItems(reviewText, "suggestions", "recommendations"),
	}

	return review
}

// extractScore extracts a numeric score from review text
func (sec *SWEngineeringCollab) extractScore(text string) int {
	// Simple score extraction - look for numbers 1-10
	for i := 10; i >= 1; i-- {
		if strings.Contains(text, fmt.Sprintf("%d", i)) {
			return i
		}
	}
	return 7 // Default neutral score
}

// extractListItems extracts list items from review text
func (sec *SWEngineeringCollab) extractListItems(text, keyword1, keyword2 string) []string {
	var items []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(strings.ToLower(line), keyword1) ||
			(keyword2 != "" && strings.Contains(strings.ToLower(line), keyword2)) {
			// Extract bullet points or numbered items
			if strings.HasPrefix(line, "-") || strings.HasPrefix(line, "*") ||
				(len(line) > 0 && line[0] >= '1' && line[0] <= '9' && strings.Contains(line, ".")) {
				items = append(items, strings.TrimSpace(line[1:]))
			}
		}
	}

	return items
}

// consolidateReviews merges all reviews with scoring
func (sec *SWEngineeringCollab) consolidateReviews() map[string]map[string]interface{} {
	assessment := make(map[string]map[string]interface{})

	for _, agent := range sec.Agents {
		sec.mu.RLock()
		workItem := sec.WorkItems[agent.Name]
		sec.mu.RUnlock()

		if workItem == nil {
			continue
		}

		// Calculate average score
		totalScore := 0
		for _, review := range workItem.Reviews {
			totalScore += review.Score
		}
		avgScore := float64(totalScore) / float64(len(workItem.Reviews))

		// Collect all feedback
		var allStrengths, allWeaknesses, allSuggestions []string
		for _, review := range workItem.Reviews {
			allStrengths = append(allStrengths, review.Strengths...)
			allWeaknesses = append(allWeaknesses, review.Weaknesses...)
			allSuggestions = append(allSuggestions, review.Suggestions...)
		}

		assessment[agent.Name] = map[string]interface{}{
			"work":         workItem.Work,
			"avg_score":    avgScore,
			"review_count": len(workItem.Reviews),
			"strengths":    sec.deduplicateList(allStrengths),
			"weaknesses":   sec.deduplicateList(allWeaknesses),
			"suggestions":  sec.deduplicateList(allSuggestions),
			"reviews":      workItem.Reviews,
		}
	}

	return assessment
}

// deduplicateList removes duplicate items from a list
func (sec *SWEngineeringCollab) deduplicateList(items []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, item := range items {
		if item != "" && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// generateFinalReport creates the comprehensive final report
func (sec *SWEngineeringCollab) generateFinalReport(assessment map[string]map[string]interface{}) string {
	report := fmt.Sprintf(`# SW Engineering Team Collaboration Report
**Generated:** %s
**Challenge:** %s
**Team Size:** %d engineers

## Executive Summary

This report presents the results of a collaborative software engineering effort where %d team members worked on the challenge, conducted peer reviews, and consolidated feedback through a structured scoring system.

## Team Performance Overview

`, time.Now().Format("2006-01-02 15:04:05"), sec.Challenge, len(sec.Agents), len(sec.Agents))

	// Team statistics
	totalReviews := 0
	totalScore := 0.0
	for _, data := range assessment {
		if reviews, ok := data["review_count"].(int); ok {
			totalReviews += reviews
		}
		if score, ok := data["avg_score"].(float64); ok {
			totalScore += score
		}
	}
	avgTeamScore := totalScore / float64(len(sec.Agents))

	report += fmt.Sprintf(`- **Total Reviews Conducted:** %d
- **Average Team Score:** %.1f/10
- **Review Coverage:** %.1f reviews per team member

## Individual Performance & Feedback

`, totalReviews, avgTeamScore, float64(totalReviews)/float64(len(sec.Agents)))

	// Individual assessments
	for _, agent := range sec.Agents {
		data := assessment[agent.Name]
		report += fmt.Sprintf(`### %s (%s)
**Average Peer Review Score:** %.1f/10 (%d reviews)

#### Key Strengths
`, agent.Name, agent.Role, data["avg_score"], data["review_count"])

		if strengths, ok := data["strengths"].([]string); ok && len(strengths) > 0 {
			for _, strength := range strengths {
				report += fmt.Sprintf("- %s\n", strength)
			}
		} else {
			report += "- No specific strengths identified\n"
		}

		report += "\n#### Areas for Improvement\n"
		if weaknesses, ok := data["weaknesses"].([]string); ok && len(weaknesses) > 0 {
			for _, weakness := range weaknesses {
				report += fmt.Sprintf("- %s\n", weakness)
			}
		} else {
			report += "- No major weaknesses identified\n"
		}

		report += "\n#### Suggested Improvements\n"
		if suggestions, ok := data["suggestions"].([]string); ok && len(suggestions) > 0 {
			for _, suggestion := range suggestions {
				report += fmt.Sprintf("- %s\n", suggestion)
			}
		} else {
			report += "- No specific suggestions provided\n"
		}

		report += "\n#### Original Implementation\n"
		if work, ok := data["work"].(string); ok {
			// Truncate very long work descriptions
			if len(work) > 500 {
				report += work[:500] + "...\n"
			} else {
				report += work + "\n"
			}
		}
		report += "\n---\n\n"
	}

	report += `## Recommendations

### For Individual Contributors
1. **Focus on Strengths:** Continue developing areas where you excel
2. **Address Weaknesses:** Prioritize improvements in identified areas
3. **Implement Suggestions:** Apply peer recommendations to enhance work quality

### For Team Process
1. **Maintain Review Quality:** Continue the peer review process for consistent improvement
2. **Knowledge Sharing:** Leverage team strengths across projects
3. **Process Refinement:** Use feedback to improve collaboration workflows

### For Future Projects
1. **Apply Lessons Learned:** Incorporate successful patterns from this collaboration
2. **Risk Mitigation:** Address common weaknesses proactively
3. **Quality Assurance:** Maintain high standards through peer review

## Conclusion

This collaborative effort demonstrates the value of structured peer review and team-based software engineering. The scoring system and consolidated feedback provide actionable insights for continuous improvement.

**Collaboration Score:** ` + fmt.Sprintf("%.1f/10", avgTeamScore) + `

---
*Report generated by Breeze AI Collaboration Framework*
`

	return report
}

// saveReportToFile saves the final report to a timestamped file
func (sec *SWEngineeringCollab) saveReportToFile(report string) error {
	filename := fmt.Sprintf("examples/sw_team_report_%s.md", time.Now().Format("2006-01-02_15-04-05"))
	return os.WriteFile(filename, []byte(report), 0644)
}

// updateProgress shows progress bar and updates
func (sec *SWEngineeringCollab) updateProgress(current, total int, phase string) {
	if sec.OnProgress != nil {
		sec.OnProgress(current, total, phase)
		return
	}

	// Default progress bar
	percentage := float64(current) / float64(total) * 100
	barWidth := 40
	filled := int(percentage / 100 * float64(barWidth))

	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	fmt.Printf("\r[%s] %.1f%% (%d/%d) - %s", bar, percentage, current, total, phase)
	if current == total {
		fmt.Println() // New line when complete
	}
}

// runSWEngineeringCollabDemo demonstrates the improved collaboration system
func runSWEngineeringCollabDemo() {
	fmt.Println("🔧 SW Engineering Team Collaboration with Peer Review")
	fmt.Println("Features:")
	fmt.Println("  ✅ Progress bar (no console spam)")
	fmt.Println("  ✅ All agents are SW engineers + code reviewers")
	fmt.Println("  ✅ Each agent does work AND reviews others")
	fmt.Println("  ✅ Scoring system for merging reviews")
	fmt.Println("  ✅ Only final report to file")
	fmt.Println()

	// Define SW engineering team (all same type: engineers who also review)
	agents := []breeze.Agent{
		{
			Name:        "Alex Chen",
			Role:        "Senior Software Engineer",
			Expertise:   "Go development and system architecture",
			Personality: "meticulous, focuses on scalable solutions",
		},
		{
			Name:        "Maria Rodriguez",
			Role:        "Software Engineer",
			Expertise:   "data structures and algorithms",
			Personality: "analytical, emphasizes code efficiency",
		},
		{
			Name:        "David Kim",
			Role:        "Software Engineer",
			Expertise:   "testing and quality assurance",
			Personality: "thorough, prioritizes reliability",
		},
		{
			Name:        "Sarah Johnson",
			Role:        "Software Engineer",
			Expertise:   "user experience and API design",
			Personality: "user-focused, emphasizes usability",
		},
		{
			Name:        "Mike Wilson",
			Role:        "Software Engineer",
			Expertise:   "performance optimization",
			Personality: "performance-driven, optimizes for speed",
		},
	}

	// The engineering challenge
	challenge := `Design and implement a REST API for a task management system with the following requirements:

1. CRUD operations for tasks (create, read, update, delete)
2. Task prioritization (high, medium, low)
3. Due date tracking and overdue detection
4. User assignment and collaboration features
5. Search and filtering capabilities
6. Data persistence with proper error handling
7. RESTful design principles
8. Input validation and security considerations

Provide a complete implementation plan including:
- API endpoint design
- Data models and relationships
- Error handling strategy
- Security measures
- Testing approach`

	// Create and run collaboration
	collab := NewSWEngineeringCollab(agents, challenge)

	// Custom progress callback
	collab.OnProgress = func(current, total int, phase string) {
		percentage := float64(current) / float64(total) * 100
		barWidth := 50
		filled := int(percentage / 100 * float64(barWidth))

		bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

		fmt.Printf("\r[%s] %.1f%% (%d/%d) %s", bar, percentage, current, total, phase)
		if current == total {
			fmt.Println() // New line when complete
		}
	}

	// Run the collaboration
	err := collab.Run()
	if err != nil {
		fmt.Printf("❌ Collaboration failed: %v\n", err)
		return
	}

	// Clear conversation
	breeze.Clear()
}
