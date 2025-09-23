package reporter

import (
	"fmt"

	"github.com/andynesse/commit-analyzer/analyzer"
)

func ConsoleReport(results []analyzer.CommitScore) {
	reportOutput := "Commit Message Analysis Report\n==============================\n\n"

	for _, res := range results {
		reportOutput += fmt.Sprintf("%s\n", commitResult(res))
	}
	reportOutput += "==============================\n"
	fmt.Print(reportOutput)
}

func commitResult(result analyzer.CommitScore) string {
	commitOutput := fmt.Sprintf("%s\n", result.Commit.Message)
	commitOutput += fmt.Sprintf("   Score: %d%%\n", result.Score)
	commitOutput += fmt.Sprintf("   Hash: %s\n   Author: %s | Date: %s\n\n", result.Commit.Hash, result.Commit.Author, result.Commit.Date)

	failedRules := 0
	for _, ruleResult := range result.Breakdown {
		if !ruleResult.Passed {
			failedRules++
		}
	}

	if failedRules > 0 {
		commitOutput += fmt.Sprintf("   %d improvement(s):\n", failedRules)
		for _, suggestion := range result.Suggestions {
			commitOutput += fmt.Sprintf("      * %s\n", suggestion)
		}
	} else {
		commitOutput += fmt.Sprintln("   All checks passed!")
	}
	return commitOutput
}
