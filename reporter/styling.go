package reporter

func getScoreColor(score int) string {
	switch {
	case score >= 80:
		return "🟢"
	case score >= 60:
		return "🟡"
	case score >= 40:
		return "🟠"
	default:
		return "🔴"
	}
}

func GetScoreEmoji(score int) string {
	switch {
	case score >= 90:
		return "🎉"
	case score >= 80:
		return "👍"
	case score >= 70:
		return "✅"
	case score >= 60:
		return "⚠️"
	case score >= 50:
		return "📝"
	default:
		return "❌"
	}
}

func getScoreBar(score int) string {
	bar := ""
	bars := score / 5
	for i := 0; i < 20; i++ {
		if i < bars {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}
