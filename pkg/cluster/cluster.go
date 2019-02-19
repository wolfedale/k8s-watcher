package cluster

import (
	"github.com/wolfedale/k8s-watcher/pkg/cluster/meta"
	"k8s.io/client-go/informers"
)

type infor informers.SharedInformerFactory

type Context struct {
	Factory infor
}

func NewContext() *Context {
	return &Context{
		Factory: meta.NewClusterMeta(),
	}
}
