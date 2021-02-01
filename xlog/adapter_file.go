// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog

// File adapter.
type FileLogAdapter struct{}

// New file adapter struct.
func NewFileLogAdapter() *FileLogAdapter {
	return &FileLogAdapter{}
}

// File adapter handler.
func (o *FileLogAdapter) Handler(line *Line) {}
