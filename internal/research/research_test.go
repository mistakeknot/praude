package research

import (
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCreateResearchFile(t *testing.T) {
	dir := t.TempDir()
	path, err := Create(dir, "PRD-001", time.Date(2026, 1, 14, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(filepath.Base(path), "PRD-001") {
		t.Fatalf("expected prd in filename")
	}
}
