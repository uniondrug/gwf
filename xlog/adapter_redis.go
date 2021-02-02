// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

// Redis适配器.
type AdapterRedisHandler struct{}

// 创建Redis适配器结构体.
func NewAdapterRedisHandler() *AdapterRedisHandler {
	o := &AdapterRedisHandler{}
	return o
}

// 指定Handler回调.
func (o *AdapterRedisHandler) Handler(line *Line) {}
