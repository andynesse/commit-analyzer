package analyzer

import (
	"github.com/andynesse/commit-analyzer/git"
)

type CommitScore struct {
	Commit      git.Commit
	Score       int
	Breakdown   map[string]RuleResult
	Suggestions []string
}

type RuleResult struct {
	Passed  bool
	Score   int
	Message string
}

type AnalysisSummary struct {
	TotalCommits int
	AverageScore float64
	BestScore    int
	WorstScore   int
}

func AnalyzeCommits(commits []git.Commit) []CommitScore {
	var results []CommitScore

	for _, commit := range commits {
		score := scoreCommit(commit)
		results = append(results, score)
	}

	return results
}

func CalculateSummary(results []CommitScore) AnalysisSummary {
	if len(results) == 0 {
		return AnalysisSummary{}
	}
	totalScore := 0
	bestScore := results[0].Score
	worstScore := results[0].Score

	for _, result := range results {
		totalScore += result.Score
		if result.Score > bestScore {
			bestScore = result.Score
		}
		if result.Score < worstScore {
			worstScore = result.Score
		}
	}
	return AnalysisSummary{
		TotalCommits: len(results),
		AverageScore: float64(totalScore) / float64(len(results)),
		BestScore:    bestScore,
		WorstScore:   worstScore,
	}
}
