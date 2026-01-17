package tui

import "strings"

func renderHelpOverlay() string {
	lines := []string{
		"Help",
		"j/k: move  /: search",
		"g: interview  r: research  p: suggestions  s: review",
		"?: help  `: tutorial  q: quit",
		"Esc: close",
	}
	return strings.Join(lines, "\n")
}

func renderTutorialOverlay() string {
	lines := []string{
		"Tutorial",
		"1) Press g to create a PRD via interview",
		"2) Press r to launch research",
		"3) Press p to generate suggestions",
		"4) Press s to review/apply suggestions",
		"Esc: close",
	}
	return strings.Join(lines, "\n")
}
