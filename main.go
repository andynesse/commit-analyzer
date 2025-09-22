package main

import (
	"fmt"

	"github.com/andynesse/commit-analyzer/git"
)

func main() {
	history, err := git.GetCommitHistory(".")
	if err != nil {
		fmt.Printf("Error %w", err)
		return
	}
	for _, commit := range history {
		fmt.Println(commit.Message)
	}
}
