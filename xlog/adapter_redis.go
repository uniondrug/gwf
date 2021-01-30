// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog

// Redis模式.
type RedisLogAdapter struct{}

// 创建Redis结构体.
func NewRedisLogAdapter() *RedisLogAdapter {
	return &RedisLogAdapter{}
}

// Redis结构体回调.
func (o *RedisLogAdapter) Handler(line *Line) {}
