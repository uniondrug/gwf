// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package xlog

import (
	"fmt"
	"time"
)

// 单条日志.
type Line struct {
	args   []interface{}
	ctx    interface{}
	format string
	level  LogLevel
	time   time.Time
	traced struct{}
}

// 创建日志记录.
func NewLine(ctx interface{}, level LogLevel, format string, args ...interface{}) *Line {
	line := &Line{args: args, ctx: ctx, format: format, level: level, time: time.Now()}
	return line
}

// 日志内容.
func (o *Line) Message() string {
	if o.args != nil && len(o.args) > 0 {
		return fmt.Sprintf(o.format, o.args...)
	}
	return o.format
}
