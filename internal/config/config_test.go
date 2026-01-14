package config

import "strings"

import "testing"

func TestDefaultConfigHasAgents(t *testing.T) {
	if !strings.Contains(DefaultConfigToml, "[agents.codex]") {
		t.Fatalf("expected codex agent profile")
	}
}
