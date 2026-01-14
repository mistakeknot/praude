package specs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSummaries(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "PRD-001.yaml"), []byte("id: \"PRD-001\"\ntitle: \"A\"\nsummary: \"S\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	list, _ := LoadSummaries(dir)
	if len(list) != 1 || list[0].ID != "PRD-001" {
		t.Fatalf("expected summary")
	}
}
