package cli

import "github.com/spf13/cobra"

func Execute() error {
	return NewRoot().Execute()
}

func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "praude",
		Short: "PM-focused PRD CLI",
	}
	return root
}
