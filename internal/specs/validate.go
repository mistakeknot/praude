package specs

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func Validate(raw []byte) error {
	var doc struct {
		ID      string `yaml:"id"`
		Title   string `yaml:"title"`
		Summary string `yaml:"summary"`
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return err
	}
	if doc.ID == "" || doc.Title == "" || doc.Summary == "" {
		return fmt.Errorf("missing required fields")
	}
	return nil
}
