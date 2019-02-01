package version

import (
	"github.com/spf13/cobra"
)

// Version is the watcher CLI version
const Version = "0.1.0-alpha"

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "prints the watcher CLI version",
		Long:  "prints the watcher CLI version",
		RunE: func(cmd *cobra.Command, args []string) error {
			println(Version)
			return nil
		},
	}
	return cmd
}
