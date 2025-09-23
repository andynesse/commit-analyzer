package git

import (
	"os/exec"
	"strconv"
	"strings"
)

type Commit struct {
	Hash    string
	Message string
	Author  string
	Date    string
}

func GetCommitHistory(repoPath, since, until string, limit int) ([]Commit, error) {
	args := []string{"log", "--pretty=format:%H|%an|%ad|%s", "--date=short"}

	if since != "" {
		args = append(args, "--since=\""+since+"\"")
	}
	if until != "" {
		args = append(args, "--until=\""+until+"\"")
	}
	if limit > 0 {
		args = append(args, "-n", strconv.Itoa(limit))
	}

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
