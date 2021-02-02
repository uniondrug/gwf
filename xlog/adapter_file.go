// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

// 文件适配器.
type AdapterFileHandler struct{}

// 创建文件适配器结构体.
func NewAdapterFileHandler() *AdapterFileHandler {
	o := &AdapterFileHandler{}
	return o
}

// 指定Handler回调.
func (o *AdapterFileHandler) Handler(line *Line) {}
