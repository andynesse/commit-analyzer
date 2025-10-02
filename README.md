# Commit Analyzer

A Go tool that analyzes Git commit messages and provides scores with improvement suggestions.

## 🚀 Quick Install

```bash
go install github.com/andynesse/commit-analyzer@latest
```
## 📖 Usage
To run a simple analysis do:
```bash
commit-analyzer
```
For additional functionality do:
```bash
commit-analyzer -help
```
## 🎯 What It Checks

    ✅ Conventional Commits format (feat:, fix:, etc.)

    ✅ Imperative mood ("Add" not "Added")

    ✅ Appropriate length (10-100 characters)

    ✅ No trailing period

    ✅ Descriptive content
    ...
    and more

## 📊 Example Output
```text
$ commit-analyzer

📝 Summary:
   Total Commits: 15
   Average Score: 89.9%  █████████████████░░░
   Best Score:    100%   🎉
   Worst Score:   25%    ❌
```
```text
$ commit-analyzer -log

Commit Message Analysis Report
================================

🎉 feat: add user authentication system
🟢 Score: 100%
    Hash: 12abcd1213abcd232abcdefgh212321
    Author: johndoe | Date: 2025-02-01
   
   ✅ All checks passed!

❌ fix stuff  
   🔴 Score: 25%
   💡 5 improvement(s):
      • Recommended improvement here ...
      • And another suggestion here ...
        ...

...

================================

📝 Summary:
...

```

## 📦 Install From Source
```bash
git clone https://github.com/andynesse/commit-analyzer
cd commit-analyzer
go build -o commit-analyzer
```
