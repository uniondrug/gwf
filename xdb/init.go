// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

// Package for framework service with xorm.
package xdb

import (
	"sync"

	"gwf/xlog"
)

var (
	Config *configuration
)

func init() {
	new(sync.Once).Do(func() {
		xlog.Info("[GWF][SERVICE] initialize.")
		Config = new(configuration)
		Config.onInit()
	})
}
