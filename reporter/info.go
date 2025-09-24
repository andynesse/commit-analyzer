package reporter

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"text/tabwriter"
)

func PrintVersion() {
	fmt.Printf("commit-analyzer %s\n", getModuleVersion())
}

func getModuleVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err == nil {
		return strings.TrimSpace(string(output))
	}

	return "dev"
}

func PrintUsage() {
	usageOutput := "commit-analyzer - Analyze the quality of your git commit messages\n\n"
	usageOutput += "📚 Usage:\n"
	usageOutput += "   commit-analyzer [flags]\n\n"
	usageOutput += "🚩 Flags:\n"

	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "   -repo string\tPath to git repository (default \".\")")
	fmt.Fprintln(w, "   -since string\tAnalyze commits since")
	fmt.Fprintln(w, "   -until string\tAnalyze commits until")
	fmt.Fprintln(w, "   -limit int\tLimit number of commits to analyze (0 = all)")
	fmt.Fprintln(w, "   -version\tShow module version")
	fmt.Fprintln(w, "   -help\tShow this help message")
	fmt.Fprintln(w, "   -log\tLists the score for every commit")
	w.Flush()
	usageOutput += b.String()

	usageOutput += "\n🛠  Examples:\n"
	usageOutput += "   commit-analyzer -log\n"
	usageOutput += "   commit-analyzer -repo ./path/to/my-repo\n"
	usageOutput += "   commit-analyzer -since \"2025-01-01\" -until \"2025-02-01\"\n"
	usageOutput += "   commit-analyzer -limit 5\n\n"
	usageOutput += "For more information, visit: https://github.com/andynesse/commit-analyzer\n"
	fmt.Println(usageOutput)
}
