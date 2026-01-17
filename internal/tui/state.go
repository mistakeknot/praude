package tui

import "github.com/mistakeknot/praude/internal/specs"

type SharedState struct {
	Summaries []specs.Summary
	Selected  int
	Focus     string
	Filter    string
}

func NewSharedState() *SharedState {
	return &SharedState{Focus: "LIST"}
}
