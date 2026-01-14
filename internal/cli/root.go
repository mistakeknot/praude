package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Execute() error {
	return NewRoot().Execute()
}

func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "praude",
		Short: "PM-focused PRD CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			if _, err := os.Stat(filepath.Join(cwd, ".praude")); err != nil {
				fmt.Fprintln(cmd.OutOrStdout(), "Not initialized. Run `praude init` first.")
				return nil
			}
			fmt.Fprintln(cmd.OutOrStdout(), "TUI not wired yet.")
			return nil
		},
	}
	return root
}
