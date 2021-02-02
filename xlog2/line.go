// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package xlog2

import (
	"context"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

// 单条日志.
type Line struct {
	ctx         interface{}
	Args        []interface{}
	Format      string
	Level       LogLevel
	Time        time.Time
	SpanId      string
	SpanVersion string
}

// 创建日志记录.
func NewLine(ctx interface{}, level LogLevel, format string, args ...interface{}) *Line {
	line := &Line{Args: args, ctx: ctx, Format: format, Level: level, Time: time.Now()}
	line.parse()
	return line
}

// 日志内容.
func (o *Line) Message() string {
	if o.Args != nil && len(o.Args) > 0 {
		return fmt.Sprintf(o.Format, o.Args...)
	}
	return o.Format
}

// 解析Context.
func (o *Line) parse() {
	// 1. 未绑定Context
	if o.ctx == nil {
		return
	}
	// 2. 标准Tracing
	if t1, o1 := o.ctx.(*Tracing); o1 && t1 != nil {
		o.parseTracing(t1)
		return
	}
	// 3. 标准context.Context.
	if c2, o2 := o.ctx.(context.Context); o2 {
		x2 := c2.Value(OpenTracing)
		if x2 != nil {
			if t2, o21 := x2.(*Tracing); o21 && t2 != nil {
				o.parseTracing(t2)
			}
		}
		return
	}
	// 4. 标准iris.Context.
	if c3, o3 := o.ctx.(iris.Context); o3 {
		x3 := c3.Values().Get(OpenTracing)
		if x3 != nil {
			if t3, o31 := x3.(*Tracing); o31 && t3 != nil {
				o.parseTracing(t3)
			}
		}
		return
	}
}

// 解析Tracing.
func (o *Line) parseTracing(t *Tracing) {
	o.SpanId = t.spanId
	o.SpanVersion = fmt.Sprintf("%s.%d", t.spanVersion, t.Increment())
}
