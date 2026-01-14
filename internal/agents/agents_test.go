package agents

import "testing"

func TestResolveProfileMissing(t *testing.T) {
	cfg := map[string]Profile{}
	if _, err := Resolve(cfg, "codex"); err == nil {
		t.Fatalf("expected error")
	}
}
