package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andynesse/commit-analyzer/analyzer"
	"github.com/andynesse/commit-analyzer/git"
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

	fmt.Printf("Found %d commits\n", len(commits))
	output := ""
	for _, res := range results {
		output += fmt.Sprintf("Score: %v%%\nMessage: %s\nDate: %v\nSuggestions: %v\n-------\n", res.Score, res.Commit.Message, res.Commit.Date, res.Suggestions)
	}
	fmt.Print(output)
}
