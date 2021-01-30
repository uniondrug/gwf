// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog

// 文件模式.
type FileLogAdapter struct{}

// 创建文件模式结构体.
func NewFileLogAdapter() *FileLogAdapter {
	return &FileLogAdapter{}
}

// 文件模式回调.
func (o *FileLogAdapter) Handler(line *Line) {}
