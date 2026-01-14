package specs

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreateTemplateSpec(t *testing.T) {
	root := t.TempDir()
	specsDir := filepath.Join(root, ".praude", "specs")
	if err := os.MkdirAll(specsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	path, err := CreateTemplate(specsDir, time.Date(2026, 1, 14, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(path)
	if !bytes.Contains(raw, []byte("id: \"PRD-001\"")) {
		t.Fatalf("expected PRD-001 id")
	}
	if !bytes.Contains(raw, []byte("strategic_context:")) {
		t.Fatalf("expected full schema")
	}
}
