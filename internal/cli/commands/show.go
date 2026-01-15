package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/specs"
	"github.com/spf13/cobra"
)

func ShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show <id>",
		Short: "Show a PRD spec",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := os.Getwd()
			if err != nil {
				return err
			}
			id := args[0]
			path := filepath.Join(project.SpecsDir(root), id+".yaml")
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			spec, err := specs.LoadSpec(path)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), string(raw))
			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintln(cmd.OutOrStdout(), summarizeSpec(spec))
			return nil
		},
	}
}

func summarizeSpec(spec specs.Spec) string {
	lines := []string{
		"Summary:",
		"CUJ: " + itoa(len(spec.CriticalUserJourneys)),
		"Market: " + itoa(len(spec.MarketResearch)),
		"Competitive: " + itoa(len(spec.CompetitiveLandscape)),
	}
	return strings.Join(lines, "\n")
}

func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}
