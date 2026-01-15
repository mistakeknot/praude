package suggestions

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mistakeknot/praude/internal/specs"
	"gopkg.in/yaml.v3"
)

type Suggestion struct {
	Summary              string
	Requirements         []string
	CriticalUserJourneys []specs.CriticalUserJourney
	MarketResearch       []specs.MarketResearchItem
	CompetitiveLandscape []specs.CompetitiveLandscapeItem
}

func Create(dir, id string, now time.Time) (string, error) {
	name := fmt.Sprintf("%s-%s.md", id, now.UTC().Format("20060102-150405"))
	path := filepath.Join(dir, name)
	body := fmt.Sprintf(`# Suggestions for %s

## Summary
- status: pending
- suggestion: ""

## Requirements
- status: pending
- suggestion:
  - "REQ-001: Add requirement"

## Critical User Journeys
- status: pending
- suggestion:
  - id: "CUJ-001"
    title: "Primary Journey"
    priority: "high"
    steps:
      - "Step"
    success_criteria:
      - "Outcome"
    linked_requirements:
      - "REQ-001"

## Market Research
- status: pending
- suggestion:
  - id: "MR-001"
    claim: "Market claim"
    evidence_refs:
      - path: ".praude/research/%s-YYYYMMDD-HHMMSS.md"
        anchor: "section-1"
        note: "Source quote"
    confidence: "medium"
    date: "2026-01-15"

## Competitive Landscape
- status: pending
- suggestion:
  - id: "COMP-001"
    name: "Competitor"
    positioning: "Positioning"
    strengths:
      - "Strength"
    weaknesses:
      - "Weakness"
    risk: "Medium"
    evidence_refs:
      - path: ".praude/research/%s-YYYYMMDD-HHMMSS.md"
        anchor: "section-2"
        note: "Source quote"
`, id, id, id)
	return path, os.WriteFile(path, []byte(body), 0o644)
}

func Apply(path string, suggestion Suggestion) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var doc specs.Spec
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return err
	}
	if suggestion.Summary != "" {
		doc.Summary = suggestion.Summary
	}
	if len(suggestion.Requirements) > 0 {
		doc.Requirements = suggestion.Requirements
	}
	if len(suggestion.CriticalUserJourneys) > 0 {
		doc.CriticalUserJourneys = suggestion.CriticalUserJourneys
	}
	if len(suggestion.MarketResearch) > 0 {
		doc.MarketResearch = suggestion.MarketResearch
	}
	if len(suggestion.CompetitiveLandscape) > 0 {
		doc.CompetitiveLandscape = suggestion.CompetitiveLandscape
	}
	updated, err := yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	return os.WriteFile(path, updated, 0o644)
}
