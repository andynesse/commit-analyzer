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

func AnalyzeCommits(commits []git.Commit) []CommitScore {
	var results []CommitScore

	for _, commit := range commits {
		score := scoreCommit(commit)
		results = append(results, score)
	}

	return results
}
