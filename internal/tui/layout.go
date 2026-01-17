package tui

import "strings"

func renderHeader(title, focus string) string {
	return "PRAUDE | " + title + " | [" + focus + "]"
}

func renderFooter(keys, status string) string {
	if strings.TrimSpace(status) == "" {
		status = "ready"
	}
	return "KEYS: " + keys + " | " + status
}

func renderFrame(header, body, footer string) string {
	return strings.Join([]string{header, body, footer}, "\n")
}
