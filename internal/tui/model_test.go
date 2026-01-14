package tui

import (
	"strings"
	"testing"
)

func TestViewIncludesHeaders(t *testing.T) {
	m := NewModel()
	out := m.View()
	if !strings.Contains(out, "PRDs") || !strings.Contains(out, "DETAILS") {
		t.Fatalf("expected headers")
	}
}
