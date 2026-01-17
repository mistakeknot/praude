package tui

import "testing"

func TestRouterSwitchesScreens(t *testing.T) {
	m := NewModel()
	if m.router.active != "list" {
		t.Fatalf("expected list screen")
	}
	m.router.Switch("help")
	if m.router.active != "help" {
		t.Fatalf("expected help screen")
	}
}
