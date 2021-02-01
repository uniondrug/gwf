// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog

import (
	"fmt"
)

// Terminal adapter.
type TermLogAdapter struct {
	colors map[LogLevel][]int
}

// New terminal adapter.
func NewTermLogAdapter() *TermLogAdapter {
	return &TermLogAdapter{
		colors: map[LogLevel][]int{
			DebugLevel: {37, 0},
			InfoLevel:  {34, 0},
			WarnLevel:  {35, 0},
			ErrorLevel: {31, 43},
			AlertLevel: {31, 43},
		},
	}
}

// Terminal adapter handler.
func (o *TermLogAdapter) Handler(line *Line) {
	println(o.format(line))
}

// Render text color.
func (o *TermLogAdapter) color(line *Line) string {
	if c, ok := o.colors[line.Level]; ok {
		return fmt.Sprintf("%c[%d;%d;%dm[%5s]%c[0m",
			0x1B, 0,
			c[1], c[0],
			Config.LevelText(line.Level),
			0x1B,
		)
	}
	return fmt.Sprintf("[%5s]", Config.LevelText(line.Level))
}

// Format terminal string.
func (o *TermLogAdapter) format(line *Line) string {
	s := fmt.Sprintf("%s", o.color(line))
	if line.SpanId != "" {
		s += fmt.Sprintf("[s=%s][v=%s]", line.SpanId, line.SpanVersion)
	}
	s += " " + line.Message()
	return s
}
