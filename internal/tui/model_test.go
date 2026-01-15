package tui

import (
	"os"
	"path/filepath"
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

func TestViewShowsFirstSpec(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, ".praude", "specs"), 0o755); err != nil {
		t.Fatal(err)
	}
	spec := "id: \"PRD-001\"\ntitle: \"Alpha\"\nsummary: \"First\"\n"
	if err := os.WriteFile(filepath.Join(root, ".praude", "specs", "PRD-001.yaml"), []byte(spec), 0o644); err != nil {
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
	out := m.View()
	if !strings.Contains(out, "Alpha") || !strings.Contains(out, "First") {
		t.Fatalf("expected spec content in view")
	}
}
