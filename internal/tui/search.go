package tui

import (
	"strings"

	"github.com/mistakeknot/praude/internal/specs"
)

func filterSummaries(items []specs.Summary, filter string) []specs.Summary {
	trim := strings.TrimSpace(filter)
	if trim == "" {
		return items
	}
	needle := strings.ToLower(trim)
	var out []specs.Summary
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.ID), needle) || strings.Contains(strings.ToLower(item.Title), needle) {
			out = append(out, item)
		}
	}
	return out
}
