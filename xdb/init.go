// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

// Package for framework service with xorm.
package xdb

import (
	"context"
	"sync"

	"github.com/uniondrug/gwf/xlog"
)

var (
	Config      *configuration
	rootContext context.Context
)

func init() {
	new(sync.Once).Do(func() {
		rootContext = context.Background()
		xlog.Info("initialize golang framework service.")
		Config = new(configuration)
		Config.onInit()
	})
}
