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
	version := flag.Bool("version", false, "Show module version")
	help := flag.Bool("help", false, "Show help information")

	flag.Usage = func() {
		reporter.PrintUsage()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		reporter.PrintVersion()
		os.Exit(0)
	}

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
