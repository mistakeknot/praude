package specs

import "testing"

func TestValidateMissingTitle(t *testing.T) {
	raw := []byte("id: \"PRD-001\"\n")
	if err := Validate(raw); err == nil {
		t.Fatalf("expected error")
	}
}
