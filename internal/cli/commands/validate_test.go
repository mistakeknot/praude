package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateCmdRejectsInvalidMode(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, ".praude", "specs"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".praude", "config.toml"), []byte("validation_mode = \"weird\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	spec := "id: \"PRD-001\"\ntitle: \"T\"\nsummary: \"S\"\n"
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
	cmd := ValidateCmd()
	cmd.SetArgs([]string{"PRD-001"})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected invalid mode error")
	}
}
