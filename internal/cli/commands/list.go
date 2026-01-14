package commands

import (
	"fmt"
	"os"

	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/specs"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List PRD specs",
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := os.Getwd()
			if err != nil {
				return err
			}
			summaries, _ := specs.LoadSummaries(project.SpecsDir(root))
			for _, s := range summaries {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", s.ID, s.Title)
			}
			return nil
		},
	}
}
