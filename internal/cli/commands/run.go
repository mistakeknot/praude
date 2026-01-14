package commands

import (
	"fmt"
	"os"

	"github.com/mistakeknot/praude/internal/agents"
	"github.com/spf13/cobra"
)

func RunCmd() *cobra.Command {
	var agent string
	cmd := &cobra.Command{
		Use:   "run <brief>",
		Short: "Run agent with brief",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			briefPath := args[0]
			profile, err := agents.Resolve(map[string]agents.Profile{}, agent)
			if err != nil {
				return err
			}
			if err := agents.Launch(profile, briefPath); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "agent not found; brief at %s\n", briefPath)
				return nil
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&agent, "agent", "codex", "Agent profile to use")
	cmd.SetOut(os.Stdout)
	return cmd
}
