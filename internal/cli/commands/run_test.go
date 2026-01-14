package commands

import (
	"testing"
)

func TestRunCommandMissingAgent(t *testing.T) {
	cmd := RunCmd()
	cmd.SetArgs([]string{"missing.md"})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error for unknown agent")
	}
}
