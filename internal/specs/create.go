package specs

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateTemplate(dir string, now time.Time) (string, error) {
	id, err := NextID(dir)
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, id+".yaml")
	doc := fmt.Sprintf(`id: "%s"
title: "Example PRD Title"
created_at: "%s"
strategic_context:
  cuj_id: "CUJ-1"
  cuj_name: "Example Journey"
  feature_id: "example-feature"
  mvp_included: true
user_story:
  text: "As a user, I want X so that Y."
  hash: "pending"
summary: |
  One paragraph describing what to build and why.
requirements:
  - "Requirement one"
acceptance_criteria:
  - id: "ac-1"
    description: "Acceptance criterion one"
files_to_modify:
  - action: "create"
    path: "path/to/file"
    description: "Why this file"
research:
  - ".praude/research/PRD-001-YYYYMMDD-HHMMSS.md"
complexity: "medium"
estimated_minutes: 25
priority: 1
`, id, now.UTC().Format(time.RFC3339))
	return path, os.WriteFile(path, []byte(doc), 0o644)
}
