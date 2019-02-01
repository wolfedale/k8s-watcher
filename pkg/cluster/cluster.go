package cluster

import (
	"github.com/wolfedale/k8s-watcher/pkg/cluster/meta"
)

type Context struct {
	factory string
}

func NewContext() *Context {
	return &Context{
		factory: meta.NewClusterMeta(),
		//factory: "foo",
	}
}
