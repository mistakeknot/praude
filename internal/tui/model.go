package tui

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/specs"
)

type Model struct {
	summaries []specs.Summary
	selected  int
	err       string
}

func NewModel() Model {
	cwd, err := os.Getwd()
	if err != nil {
		return Model{err: err.Error()}
	}
	if _, err := os.Stat(project.RootDir(cwd)); err != nil {
		return Model{err: "Not initialized"}
	}
	list, _ := specs.LoadSummaries(project.SpecsDir(cwd))
	return Model{summaries: list}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.selected < len(m.summaries)-1 {
				m.selected++
			}
		case "k", "up":
			if m.selected > 0 {
				m.selected--
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	left := m.renderList()
	right := m.renderDetail()
	return joinColumns(left, right, 42)
}

func (m Model) renderList() []string {
	lines := []string{"PRDs"}
	if m.err != "" {
		lines = append(lines, m.err)
		return lines
	}
	if len(m.summaries) == 0 {
		lines = append(lines, "No PRDs yet.")
		return lines
	}
	for i, s := range m.summaries {
		prefix := "  "
		if i == m.selected {
			prefix = "> "
		}
		lines = append(lines, prefix+s.ID+" "+s.Title)
	}
	return lines
}

func (m Model) renderDetail() []string {
	lines := []string{"DETAILS"}
	if m.err != "" {
		lines = append(lines, "Initialize with praude init.")
		return lines
	}
	if len(m.summaries) == 0 {
		lines = append(lines, "No PRD selected.")
		return lines
	}
	sel := m.summaries[m.selected]
	lines = append(lines, "ID: "+sel.ID)
	lines = append(lines, "Title: "+sel.Title)
	lines = append(lines, "Summary: "+sel.Summary)
	return lines
}

func joinColumns(left, right []string, leftWidth int) string {
	max := len(left)
	if len(right) > max {
		max = len(right)
	}
	var b strings.Builder
	for i := 0; i < max; i++ {
		l := ""
		r := ""
		if i < len(left) {
			l = left[i]
		}
		if i < len(right) {
			r = right[i]
		}
		b.WriteString(padRight(l, leftWidth))
		b.WriteString(" | ")
		b.WriteString(r)
		if i < max-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
