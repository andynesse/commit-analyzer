package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andynesse/commit-analyzer/analyzer"
	"github.com/andynesse/commit-analyzer/git"
	"github.com/andynesse/commit-analyzer/reporter"
)

func main() {
	commits, err := git.GetCommitHistory(".")
	if err != nil {
		log.Fatal("Failed to get commit history: ", err)
	}
	if len(commits) == 0 {
		fmt.Println("No commits found")
		os.Exit(0)
	}

	results := analyzer.AnalyzeCommits(commits)

	reporter.ConsoleReport(results)
}
