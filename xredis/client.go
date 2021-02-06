// author: wsfuyibing <websearch@163.com>
// date: 2021-02-06

package xredis

import (
	"github.com/gomodule/redigo/redis"
)

type commands struct {
}

// 打开连接.
func Client() redis.Conn {
	return Config.pools.Get()
}
