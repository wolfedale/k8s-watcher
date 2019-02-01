package get

import (
	"github.com/spf13/cobra"
	"github.com/wolfedale/k8s-watcher/cmd/watcher/get/pods"
)

// NewCommand returns a new cobra.Command for get
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "get",
		Short: "get shows info about specific object",
		Long:  "get shows info about specific objects",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	// add subcommands
	cmd.AddCommand(pods.NewCommand())
	return cmd
}
