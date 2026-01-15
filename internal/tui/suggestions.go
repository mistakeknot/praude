package tui

import (
	"path/filepath"

	"github.com/mistakeknot/praude/internal/git"
	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/suggestions"
)

type suggestionsState struct {
	active bool
	id     string
	path   string
	sugg   suggestions.Suggestion
	err    string
}

func (m *Model) enterSuggestions() {
	if len(m.summaries) == 0 {
		return
	}
	id := m.summaries[m.selected].ID
	dir := project.SuggestionsDir(m.root)
	sugg, path, err := suggestions.LoadLatest(dir, id)
	if err != nil {
		m.suggestions = suggestionsState{active: true, id: id, err: err.Error()}
		m.mode = "suggestions"
		return
	}
	m.suggestions = suggestionsState{active: true, id: id, path: path, sugg: sugg}
	m.mode = "suggestions"
}

func (m *Model) applySuggestions() {
	if m.suggestions.path == "" {
		return
	}
	specPath := filepath.Join(project.SpecsDir(m.root), m.suggestions.id+".yaml")
	if err := suggestions.Apply(specPath, m.suggestions.sugg); err != nil {
		m.suggestions.err = err.Error()
		return
	}
	if err := git.EnsureRepo(m.root); err == nil {
		_ = git.CommitFiles(m.root, []string{specPath}, "chore(praude): apply suggestions "+m.suggestions.id)
	}
	m.reloadSummaries()
}

func (m *Model) renderSuggestions() []string {
	lines := []string{"Suggestions"}
	if m.suggestions.err != "" {
		lines = append(lines, "Error: "+m.suggestions.err)
		return lines
	}
	lines = append(lines, "Spec: "+m.suggestions.id)
	if m.suggestions.sugg.Summary != "" {
		lines = append(lines, "Summary suggestion: "+m.suggestions.sugg.Summary)
	}
	lines = append(lines, "[a] accept  [r] reject  [q] quit")
	return lines
}
