// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

// Package for framework logger.
package xlog

import (
	"sync"
)

type LogAdapter int
type LogLevel int

// 适配器类型.
const (
	TermAdapter LogAdapter = iota
	FileAdapter
	RedisAdapter
)

// 日志级别类型.
const (
	UnknownLevel LogLevel = iota
	OffLevel
	AlertLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// 转让配置.
const (
	DefaultAdapter    = TermAdapter
	DefaultLevel      = InfoLevel
	DefaultTimeFormat = "2006-01-02 15:04:05.999999"
)

var (
	Config Configuration
	// 适配器定义.
	AdapterText = map[LogAdapter]string{
		TermAdapter:  "term",
		FileAdapter:  "file",
		RedisAdapter: "redis",
	}
	// 日志级别定义.
	LevelText = map[LogLevel]string{
		OffLevel:     "OFF",
		AlertLevel:   "ALERT",
		ErrorLevel:   "ERROR",
		WarnLevel:    "WARN",
		InfoLevel:    "INFO",
		DebugLevel:   "DEBUG",
		UnknownLevel: "UNKNOWN",
	}
)

// 初始化X-Log.
func init() {
	new(sync.Once).Do(func() {
		Config = new(configuration)
		Config.onInit()
		Config.listen()
	})
}