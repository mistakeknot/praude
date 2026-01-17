package tui

import (
	"testing"

	"github.com/mistakeknot/praude/internal/specs"
)

func TestSearchFiltersList(t *testing.T) {
	state := NewSharedState()
	state.Summaries = []specs.Summary{
		{ID: "PRD-001", Title: "Alpha"},
		{ID: "PRD-002", Title: "Beta"},
	}
	state.Filter = "Alpha"
	items := filterSummaries(state.Summaries, state.Filter)
	if len(items) != 1 {
		t.Fatalf("expected filtered list")
	}
}
