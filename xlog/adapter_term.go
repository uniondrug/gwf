// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

import (
	"fmt"
	"os"
)

// 终端适配器.
type AdapterTermHandler struct {
	colors map[LogLevel][]int
}

// 创建终端适配器结构体.
func NewAdapterTermHandler() *AdapterTermHandler {
	o := &AdapterTermHandler{
		// 1:红, 2:绿, 3:黄, 4: 蓝
		// 5:粉, 6:青, 7:灰
		colors: map[LogLevel][]int{
			DebugLevel: {30, 47},
			InfoLevel:  {37, 44},
			WarnLevel:  {31, 43},
			ErrorLevel: {33, 41},
		},
	}
	return o
}

// 指定Handler回调.
func (o *AdapterTermHandler) Handler(line *Line) {
	// 1. Level级别.
	s := "[" + line.Time.Format(Config.TimeFormat) + "]" + o.color(line)
	// 2. Tracing链.
	if line.Tracing != nil {
		s += fmt.Sprintf("[s=%s][v=%s]", line.Tracing.spanId, fmt.Sprintf("%s.%d", line.Tracing.spanVersion, line.TracingOffset))
	}
	// 3. 日志内容.
	if line.Args != nil && len(line.Args) > 0 {
		s += " " + fmt.Sprintf(line.Format, line.Args...)
	} else {
		s += " " + line.Format
	}
	// 4. 输出内容.
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", s)
}

// 日志级别着色.
func (o *AdapterTermHandler) color(line *Line) string {
	if c, ok := o.colors[line.Level]; ok {
		return fmt.Sprintf("%c[%d;%d;%dm[%5s]%c[0m",
			0x1B, 0,
			c[1], c[0],
			LevelTexts[line.Level],
			0x1B,
		)
	}
	return fmt.Sprintf("[%5s]", LevelTexts[line.Level])
}
