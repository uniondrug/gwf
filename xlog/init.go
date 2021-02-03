// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

// Package for framework logger.
package xlog

import (
	"sync"
)

// 适配器定义.
type LogAdapter int

const (
	FileAdapter LogAdapter = iota
	RedisAdapter
	TermAdapter
)

var AdapterTexts = map[LogAdapter]string{
	FileAdapter:  "file",
	RedisAdapter: "redis",
	TermAdapter:  "term",
}

// 日志级别定义.
type LogLevel int

const (
	OffLevel LogLevel = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

var LevelTexts = map[LogLevel]string{
	OffLevel:   "OFF",
	ErrorLevel: "ERROR",
	WarnLevel:  "WARN",
	InfoLevel:  "INFO",
	DebugLevel: "DEBUG",
}

// Handler
type LogHandler func(line *Line)

// 默认配置.
const (
	DefaultAdapter      = TermAdapter
	DefaultLevel        = DebugLevel
	DefaultTimeFormat   = "2006-01-02 15:04:05.999999"
	DefaultParentSpanId = "X-B3-Parentspanid"
	DefaultSpanId       = "X-B3-Spanid"
	DefaultSpanVersion  = "Version"
	DefaultTraceId      = "X-B3-Traceid"
)

// 配置结构体.
var Config *Configuration

// 包级初始化.
func init() {
	new(sync.Once).Do(func() {
		Config = NewConfiguration()
	})
}
