// author: wsfuyibing <websearch@163.com>
// date: 2021-02-06

package xredis

import (
	"sync"
	"time"
)

const (
	LockExpiration      = 10                             // 被锁资源, 默认10秒后自动删除.
	LockExpirationReset = time.Second * time.Duration(3) // 资源被锁后, 每隔3秒自动续期, 防止因业务未执行完成而自动被释放.
)

var (
	Config *configuration
)

func init() {
	new(sync.Once).Do(func() {
		Config = new(configuration)
		Config.onInit()
	})
}
