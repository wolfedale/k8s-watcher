package pods

import (
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for getting the list of pods
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "pods",
		Short: "pods will show pods events",
		Long:  "pods will show pods events",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(cmd, args)
		},
	}
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	/*
		pods, err := pod.Run()
		if err != nil {
			return errors.Wrap(err, "error show pods")
		}

		return nil
	*/
	return nil
}
