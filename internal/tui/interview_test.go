package tui

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInterviewCreatesSpecWithWarnings(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, ".praude", "specs"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, ".praude", "research"), 0o755); err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(cwd) }()
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	m := NewModel()
	m = pressKey(m, "g")
	m = pressKey(m, "n")
	m = pressKey(m, "y")
	m = typeAndEnter(m, "Vision statement")
	m = typeAndEnter(m, "Primary users")
	m = typeAndEnter(m, "Problem to solve")
	m = typeAndEnter(m, "First requirement")
	m = pressKey(m, "n")
	entries, err := os.ReadDir(filepath.Join(root, ".praude", "specs"))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected one spec file, got %d", len(entries))
	}
	path := filepath.Join(root, ".praude", "specs", entries[0].Name())
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "critical_user_journeys") {
		t.Fatalf("expected cuj section")
	}
	if !strings.Contains(string(raw), "validation_warnings") {
		t.Fatalf("expected validation warnings metadata")
	}
}

func TestInterviewMentionsPMFocusedAgent(t *testing.T) {
	m := NewModel()
	m.mode = "interview"
	m.interview = startInterview(m.root)
	out := m.View()
	if !strings.Contains(out, "PM-focused") {
		t.Fatalf("expected PM-focused agent hint")
	}
}

func pressKey(m Model, key string) Model {
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)}
	if key == "enter" {
		msg = tea.KeyMsg{Type: tea.KeyEnter}
	}
	updated, _ := m.Update(msg)
	return updated.(Model)
}

func typeAndEnter(m Model, input string) Model {
	for _, r := range input {
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		updated, _ := m.Update(msg)
		m = updated.(Model)
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	return updated.(Model)
}
