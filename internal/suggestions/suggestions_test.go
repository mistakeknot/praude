package suggestions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mistakeknot/praude/internal/specs"
)

func TestCreateSuggestionTemplateIncludesSections(t *testing.T) {
	dir := t.TempDir()
	path, err := Create(dir, "PRD-001", time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	content := string(raw)
	if !strings.Contains(content, "## Summary") {
		t.Fatalf("expected summary section")
	}
	if !strings.Contains(content, "## Requirements") {
		t.Fatalf("expected requirements section")
	}
	if !strings.Contains(content, "## Critical User Journeys") {
		t.Fatalf("expected cuj section")
	}
}

func TestApplySuggestionUpdatesSummary(t *testing.T) {
	root := t.TempDir()
	specPath := filepath.Join(root, "PRD-001.yaml")
	specRaw := []byte("id: \"PRD-001\"\ntitle: \"Title\"\nsummary: \"Old\"\n")
	if err := os.WriteFile(specPath, specRaw, 0o644); err != nil {
		t.Fatal(err)
	}
	sugg := Suggestion{Summary: "New summary"}
	if err := Apply(specPath, sugg); err != nil {
		t.Fatal(err)
	}
	updated, err := os.ReadFile(specPath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(updated), "summary: New summary") {
		t.Fatalf("expected updated summary")
	}
	res, err := specs.Validate(updated, specs.ValidationOptions{Mode: specs.ValidationSoft, Root: root})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("expected no errors")
	}
}
