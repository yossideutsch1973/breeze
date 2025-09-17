package funcs

import (
	"fmt"

	"github.com/user/breeze"
)

// RunAIDocSummarize demonstrates document summarization with AI
func RunAIDocSummarize() {
	resp := breeze.AI(
		"Summarize the attached requirements document.",
		breeze.WithDocs("requirements.txt"),
		breeze.WithConcise(),
	)
	fmt.Println(resp)
}
