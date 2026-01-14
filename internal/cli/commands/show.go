package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mistakeknot/praude/internal/project"
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
			fmt.Fprintln(cmd.OutOrStdout(), string(raw))
			return nil
		},
	}
}
