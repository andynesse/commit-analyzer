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
	commitOutput := fmt.Sprintf("%s %s\n", GetScoreEmoji(result.Score), result.Commit.Message)
	commitOutput += fmt.Sprintf("%s Score: %d%%\n", getScoreColor(result.Score), result.Score)
	commitOutput += fmt.Sprintf("   Hash: %s\n   Author: %s | Date: %s\n\n", result.Commit.Hash, result.Commit.Author, result.Commit.Date)

	failedRules := 0
	for _, ruleResult := range result.Breakdown {
		if !ruleResult.Passed {
			failedRules++
		}
	}

	if failedRules > 0 {
		commitOutput += fmt.Sprintf("   💡 %d improvement(s):\n", failedRules)
		for _, suggestion := range result.Suggestions {
			commitOutput += fmt.Sprintf("      • %s\n", suggestion)
		}
	} else {
		commitOutput += fmt.Sprintln("   ✅ All checks passed!")
	}
	return commitOutput
}

func ConsoleSummary(summary analyzer.AnalysisSummary) {
	summaryOutput := "📝 Summary:\n"
	summaryOutput += fmt.Sprintf("   Total Commits: %d\n", summary.TotalCommits)
	summaryOutput += fmt.Sprintf("   Average Score: %.1f%%\t %s\n", summary.AverageScore, getScoreBar(int(summary.AverageScore)))
	summaryOutput += fmt.Sprintf("   Best Score:    %d%%\t %s\n", summary.BestScore, GetScoreEmoji(summary.BestScore))
	summaryOutput += fmt.Sprintf("   Worst Score:   %d%%\t %s\n", summary.WorstScore, GetScoreEmoji(summary.WorstScore))
	fmt.Println(summaryOutput)
}
