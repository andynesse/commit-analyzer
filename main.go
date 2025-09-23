package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andynesse/commit-analyzer/analyzer"
	"github.com/andynesse/commit-analyzer/git"
	"github.com/andynesse/commit-analyzer/reporter"
)

func main() {
	repoPath := flag.String("repo", ".", "Path to git repository")
	since := flag.String("since", "", "Analyze commits since")
	until := flag.String("until", "", "Analyze commits until")
	limit := flag.Int("limit", 0, "Limit number of commits to analyze (0 = All)")
	flag.Parse()

	commits, err := git.GetCommitHistory(*repoPath, *since, *until, *limit)
	if err != nil {
		log.Fatal("Failed to get commit history: ", err)
	}
	if len(commits) == 0 {
		fmt.Println("No commits found")
		os.Exit(0)
	}

	results := analyzer.AnalyzeCommits(commits)
	reporter.ConsoleReport(results)

	analysisSummary := analyzer.CalculateSummary(results)
	reporter.ConsoleSummary(analysisSummary)
}
