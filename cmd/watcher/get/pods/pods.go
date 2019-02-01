package pods

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wolfedale/k8s-watcher/pkg/cluster"
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
	ctx := cluster.NewContext()

	fmt.Println(ctx)
	/*
		if err = ctx.Create(cfg, flags.Retain, flags.Wait); err != nil {
			return errors.Wrap(err, "failed to create cluster")
		}
	*/

	return nil
}
