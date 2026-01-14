package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/research"
	"github.com/spf13/cobra"
)

func ResearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "research <id>",
		Short: "Create a research artifact for a PRD",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := os.Getwd()
			if err != nil {
				return err
			}
			id := args[0]
			path, err := research.Create(project.ResearchDir(root), id, time.Now())
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), filepath.Base(path))
			return nil
		},
	}
}
