package cli

import "testing"

func TestRootCommandHasInit(t *testing.T) {
	cmd := NewRoot()
	if cmd == nil || cmd.Use != "praude" {
		t.Fatalf("expected root command")
	}
}
