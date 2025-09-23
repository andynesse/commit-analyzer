package analyzer

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

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
		Name:        "length_check",
		Description: "Keep message length between 10-50 characters",
		Weight:      10,
		Check:       checkLength,
	},
	{
		Name:        "required_type_check",
		Description: "Start the message with a type followed by a colon and space",
		Weight:      20,
		Check:       requiredTypeCheck,
	},
	{
		Name:        "conventional_type_check",
		Description: "Use a type listed by the Conventional Commits",
		Weight:      10,
		Check:       conventionalTypeCheck,
	},
	{
		Name:        "structure_check",
		Description: "Follow Conventional Commits structure: <type>[optional scope]: <description>",
		Weight:      15,
		Check:       structureCheck,
	},
	{
		Name:        "descriptive_check",
		Description: "The message should describe what and why",
		Weight:      15,
		Check:       descriptiveCheck,
	},
	{
		Name:        "no_period_check",
		Description: "Don't end message with a period",
		Weight:      7,
		Check:       noPeriodCheck,
	},
	{
		Name:        "imperative_mood_check",
		Description: "Start with an imperative verb ('Add' instead of 'Added')",
		Weight:      20,
		Check:       imperativeMoodCheck,
	},
	{
		Name:        "capitalized_check",
		Description: "Capitalize the description",
		Weight:      10,
		Check:       capitalizedCheck,
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
			return fmt.Sprintf("Commit message is too long. %s", rule.Description)
		}
		return fmt.Sprintf("Commit message is too short. %s", rule.Description)
	default:
		return rule.Description
	}
}

func checkLength(msg string) bool {
	length := len(strings.TrimSpace(msg))
	return length >= 10 && length <= 50
}

func requiredTypeCheck(msg string) bool {
	parts := strings.SplitN(msg, ":", 2)
	if len(parts) < 2 {
		return false
	}
	return strings.TrimSpace(parts[1]) != ""
}

func conventionalTypeCheck(msg string) bool {
	commonTypes := []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "build", "ci", "chore", "revert"}
	parts := strings.SplitN(msg, ":", 2)
	if len(parts) < 2 {
		return false
	}
	typePart := strings.TrimSuffix(strings.TrimSpace(parts[0]), "!")
	if idx := strings.Index(typePart, "("); idx != -1 {
		typePart = typePart[:idx]
	}
	for _, commonType := range commonTypes {
		if strings.EqualFold(typePart, commonType) {
			return true
		}
	}
	return false
}

func structureCheck(msg string) bool {
	pattern := `^(\w+)(\([^)]+\))?!?: .+`
	matched, _ := regexp.MatchString(pattern, msg)
	return matched
}

func descriptiveCheck(msg string) bool {
	cleanMsg := msg
	if idx := strings.Index(msg, ": "); idx != -1 {
		cleanMsg = strings.TrimSpace(msg[idx+1:])
	}

	words := strings.Fields(cleanMsg)
	if len(words) < 3 {
		return false
	}

	totalChars := 0
	for _, word := range words {
		totalChars += len(word)
	}
	avgLength := float64(totalChars) / float64(len(words))
	return avgLength >= 3.5
}

func noPeriodCheck(msg string) bool {
	trimmed := strings.TrimSpace(msg)
	if len(trimmed) == 0 {
		return true
	}
	lastChar := trimmed[len(trimmed)-1]
	return lastChar != '.'
}

func imperativeMoodCheck(msg string) bool {
	imperativeVerbs := []string{"add", "fix", "update", "remove", "implement", "refactor", "document", "improve", "optimize", "simplify", "resolve", "create", "delete", "change", "move", "rename", "bump"}
	cleanMsg := msg
	if idx := strings.Index(msg, ": "); idx != -1 {
		cleanMsg = strings.TrimSpace(msg[idx+1:])
	}
	if cleanMsg == "" {
		return false
	}

	firstWord := strings.ToLower(strings.Fields(cleanMsg)[0])
	for _, verb := range imperativeVerbs {
		if firstWord == verb {
			return true
		}
	}
	return false
}

func capitalizedCheck(msg string) bool {
	cleanMsg := msg
	if idx := strings.Index(msg, ": "); idx != -1 {
		cleanMsg = strings.TrimSpace(msg[idx+1:])
	}
	if cleanMsg == "" {
		return false
	}
	firstWord := strings.Fields(cleanMsg)[0]
	return unicode.IsUpper(rune(firstWord[0]))
}
