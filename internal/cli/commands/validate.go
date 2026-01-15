package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mistakeknot/praude/internal/config"
	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/specs"
	"github.com/spf13/cobra"
)

func ValidateCmd() *cobra.Command {
	var mode string
	cmd := &cobra.Command{
		Use:   "validate <id>",
		Short: "Validate a PRD spec",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := os.Getwd()
			if err != nil {
				return err
			}
			cfg, err := config.LoadFromRoot(root)
			if err != nil {
				return err
			}
			selected := mode
			if selected == "" {
				selected = cfg.ValidationMode
			}
			if err := validateMode(selected); err != nil {
				return err
			}
			id := args[0]
			path, err := resolveSpecPath(project.SpecsDir(root), id)
			if err != nil {
				return err
			}
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			res, err := specs.Validate(raw, specs.ValidationOptions{
				Mode: specs.ValidationMode(selected),
				Root: root,
			})
			if err != nil {
				return err
			}
			if len(res.Errors) > 0 {
				return fmt.Errorf("validation failed: %s", strings.Join(res.Errors, "; "))
			}
			if len(res.Warnings) > 0 {
				if specs.ValidationMode(selected) == specs.ValidationSoft {
					if err := specs.StoreValidationWarnings(path, res.Warnings); err != nil {
						return err
					}
				}
				for _, warning := range res.Warnings {
					fmt.Fprintln(cmd.OutOrStdout(), "WARN:", warning)
				}
				return nil
			}
			fmt.Fprintln(cmd.OutOrStdout(), "OK")
			return nil
		},
	}
	cmd.Flags().StringVar(&mode, "mode", "", "Validation mode (hard|soft)")
	return cmd
}

func validateMode(mode string) error {
	if mode != "hard" && mode != "soft" {
		return fmt.Errorf("invalid validation mode %q", mode)
	}
	return nil
}

func resolveSpecPath(dir, id string) (string, error) {
	path := filepath.Join(dir, id+".yaml")
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	alt := filepath.Join(dir, id+".yml")
	if _, err := os.Stat(alt); err == nil {
		return alt, nil
	}
	return "", fmt.Errorf("spec not found: %s", id)
}
