// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog2

// Redis adapter.
type RedisLogAdapter struct{}

// New redis adapter struct.
func NewRedisLogAdapter() *RedisLogAdapter {
	return &RedisLogAdapter{}
}

// Redis adapter handler.
func (o *RedisLogAdapter) Handler(line *Line) {}
