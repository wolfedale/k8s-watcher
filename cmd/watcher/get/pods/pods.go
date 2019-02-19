package pods

import (


	"github.com/spf13/cobra"
	"github.com/wolfedale/k8s-watcher/pkg/cluster"
	"k8s.io/client-go/tools/cache"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/api/core/v1"
	"fmt"
	"log"
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

	informer := ctx.Factory.Core().V1().Events().Informer()
    stopper := make(chan struct{})
    defer close(stopper)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(obj interface{}) {
            mObj := obj.(metav1.Object)
            m := obj.(runtime.Object)
            event := watch.Event{Object: m}
            podEvent := event.Object.(*v1.Event)
            fmt.Println(podEvent.Message)
            fmt.Println(podEvent.InvolvedObject.Namespace)
            log.Printf("[Events]: %v", mObj.GetNamespace())
        },
	})

	informer.Run(stopper)

	return nil
}
