package git

import (
	"os/exec"
	"strings"
)

type Commit struct {
	Hash    string
	Message string
	Author  string
	Date    string
}

func GetCommitHistory(repoPath string) ([]Commit, error) {
	args := []string{"log", "--pretty=format:%H|%an|%ad|%s", "--date=short"}
	cmd := exec.Command("git", args...)
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseGitOutput(string(output)), nil
}

func parseGitOutput(output string) []Commit {
	var commits []Commit
	if output == "" {
		return commits
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")

	for _, line := range lines {
		fields := strings.SplitN(line, "|", 4)
		if len(fields) == 4 {
			commits = append(commits, Commit{
				Hash:    fields[0],
				Author:  fields[1],
				Date:    fields[2],
				Message: fields[3],
			})
		}
	}
	return commits
}
