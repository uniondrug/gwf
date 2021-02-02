// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

import (
	"context"
	"time"

	"github.com/kataras/iris/v12"
)

// 日志行记录.
type Line struct {
	Args          []interface{}
	Ctx           interface{}
	Format        string
	Time          time.Time
	Level         LogLevel
	Tracing       *Tracing
	TracingOffset int32
}

// 创建日志行记录.
func NewLine(ctx interface{}, level LogLevel, format string, args ...interface{}) *Line {
	line := &Line{Args: args, Ctx: ctx, Format: format, Level: level, Time: time.Now()}
	line.parse()
	return line
}

// 解析Ctx.
func (o *Line) parse() {
	// 空Ctx.
	if o.Ctx == nil {
		return
	}
	// 校验 iris.Context.
	if o.parseWithIrisContext() {
		return
	}
	// 校验 context.Context.
	if o.parseWithContext() {
		return
	}
	// 校验 *Tracing
	if t, ok := o.Ctx.(*Tracing); ok && t != nil {
		o.parseTracing(t)
		return
	}
}

// 解析*Tracing.
func (o *Line) parseTracing(tracing *Tracing) {
	o.Tracing = tracing
	o.TracingOffset, _ = tracing.incrOffset()
}

// 校验是否为context.Context上下文.
func (o *Line) parseWithContext() bool {
	c, ok1 := o.Ctx.(context.Context)
	// interface.
	if !ok1 || c == nil {
		return false
	}
	// with value
	if x := c.Value(OpenTracing); x != nil {
		if t, ok2 := x.(*Tracing); ok2 && t != nil {
			o.parseTracing(t)
			return true
		}
	}
	return false
}

// 校验是否为iris.Context上下文.
func (o *Line) parseWithIrisContext() bool {
	c, ok1 := o.Ctx.(iris.Context)
	// interface.
	if !ok1 || c == nil {
		return false
	}
	// with value
	if x := c.Values().Get(OpenTracing); x != nil {
		if t, ok2 := x.(*Tracing); ok2 && t != nil {
			o.parseTracing(t)
			return true
		}
	}
	return false
}
