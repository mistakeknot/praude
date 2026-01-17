package tui

import (
	"strings"
	"testing"
)

func TestLayoutIncludesHeaderFooter(t *testing.T) {
	m := NewModel()
	out := m.View()
	if !strings.Contains(out, "PRAUDE") {
		t.Fatalf("expected header")
	}
	if !strings.Contains(out, "KEYS:") {
		t.Fatalf("expected footer")
	}
}
