package analyzer

import (
	"regexp"
	"strings"

	"github.com/andynesse/commit-analyzer/git"
)

type Rule struct {
	Name        string
	Description string
	Weight      int
	Check       func(string) bool
}

var Rules = []Rule{
	{
		Name:   "length_check",
		Weight: 10,
		Check:  checkLength,
	},
	{
		Name:   "type_check",
		Weight: 10,
		Check:  checkConventional,
	},
}

func scoreCommit(commit git.Commit) CommitScore {
	breakdown := make(map[string]RuleResult)
	totalScore := 0
	maxPossible := 0
	var suggestions []string

	for _, rule := range Rules {
		maxPossible += rule.Weight
		passed := rule.Check(commit.Message)

		score := 0
		if passed {
			score += rule.Weight
		} else {
			suggestions = append(suggestions, generateSuggestion(rule, commit.Message))
		}

		totalScore += score

		breakdown[rule.Name] = RuleResult{
			Passed:  passed,
			Score:   score,
			Message: rule.Description,
		}
	}

	normalizedScore := 0
	if maxPossible > 0 {
		normalizedScore = (totalScore * 100) / maxPossible
	}

	return CommitScore{
		Commit:      commit,
		Score:       normalizedScore,
		Breakdown:   breakdown,
		Suggestions: suggestions,
	}
}

func generateSuggestion(rule Rule, message string) string {
	switch rule.Name {
	case "length_check":
		if len(message) > 50 {
			return "Commit message is too long"
		}
		return "Commit message is too short"
	case "type_check":
		return "Use conventional commit format: 'type: description'  ('feat: Add new feature')"
	default:
		return "Improve commit message quality"
	}
}

func checkLength(msg string) bool {
	length := len(strings.TrimSpace(msg))
	return length > 10 && length <= 50
}

func checkConventional(msg string) bool {
	pattern := `^(feat|fix|docs|style|refactor|test|chore|perf|build|ci|revert)(\([^)]+\))?!?: .+`
	matched, _ := regexp.MatchString(pattern, msg)
	return matched
}
