// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package xlog

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// 配置接口.
type Configuration interface {
	// 导出方法.
	Level() LogLevel
	LevelText(level LogLevel) string
	LoadYaml(path string) error
	SetAdapter(adapter LogAdapter)
	SetHandler(handler func(line *Line))
	SetLevel(level LogLevel)
	SetTimeFormat(timeFormat string)
	// 包内方法.
	log(ctx interface{}, level LogLevel, format string, args ...interface{})
	onInit()
	listen()
}

// 日志配置.
type configuration struct {
	AdapterName string `yaml:"adapter-name"`
	LevelName   string `yaml:"level-name"`
	TimeFormat  string `yaml:"time-format"`
	adapter     LogAdapter
	handler     func(line *Line)
	level       LogLevel
	ch          chan *Line
}

// 读取日志级别.
func (o *configuration) Level() LogLevel {
	return o.level
}

// 级别名称.
func (o *configuration) LevelText(level LogLevel) string {
	if s, ok := LevelText[level]; ok {
		return s
	}
	return ""
}

// 从YAML文件读取配置.
func (o *configuration) LoadYaml(path string) error {
	data, err := ioutil.ReadFile(path)
	// 1. 读取文件出错.
	if err != nil {
		return err
	}
	// 2. 解析JSON出错.
	if err = yaml.Unmarshal(data, o); err != nil {
		return err
	}
	// 3. 日志级别检查.
	if o.LevelName != "" {
		lf := true
		ln := strings.ToUpper(o.LevelName)
		for l, n := range LevelText {
			if n == ln {
				o.SetLevel(l)
				lf = false
			}
		}
		if lf {
			panic(fmt.Sprintf("unknown log level `%s` defined logger.yaml", o.LevelName))
		}
	} else {
		o.LevelName = LevelText[o.level]
	}
	// 4. 日志适配器.
	if o.AdapterName != "" {
		af := true
		an := strings.ToLower(o.AdapterName)
		for a, n := range AdapterText {
			if n == an {
				o.SetAdapter(a)
				af = false
				break
			}
		}
		if af {
			panic(fmt.Sprintf("unknown adapter name `%s` defined logger.yaml", o.AdapterName))
		}
	}
	// n. ended
	return nil
}

// 设置适配器.
func (o *configuration) SetAdapter(adapter LogAdapter) {
	o.adapter = adapter
	switch adapter {
	case FileAdapter:
		o.handler = NewFileLogAdapter().Handler
	case TermAdapter:
		o.handler = NewTermLogAdapter().Handler
	case RedisAdapter:
		o.handler = NewRedisLogAdapter().Handler
	}
}

// 设置回调.
func (o *configuration) SetHandler(handler func(line *Line)) {
	o.handler = handler
}

// 设置级别.
func (o *configuration) SetLevel(level LogLevel) {
	o.level = level
	enableDebug = level >= DebugLevel
	enableInfo = level >= InfoLevel
	enableWarn = level >= WarnLevel
	enableError = level >= ErrorLevel
	enableAlert = level >= AlertLevel
}

// 设置时间格式.
func (o *configuration) SetTimeFormat(timeFormat string) {
	o.TimeFormat = timeFormat
}

// 初始化Channel.
func (o *configuration) listen() {
	go func() {
		defer o.listen()
		for {
			select {
			case x := <-o.ch:
				go o.handler(x)
			}
		}
	}()
}

// 发送日志.
func (o *configuration) log(ctx interface{}, level LogLevel, format string, args ...interface{}) {
	if o.handler != nil {
		o.handler(NewLine(ctx, level, format, args...))
	}
}

// 在包的init方法中触发.
func (o *configuration) onInit() {
	o.ch = make(chan *Line)
	// default fields.
	o.SetAdapter(DefaultAdapter)
	o.SetLevel(DefaultLevel)
	o.SetTimeFormat(DefaultTimeFormat)
	// read from yaml.
	for _, path := range []string{"./config/logger.yaml", "../config/logger.yaml"} {
		err := o.LoadYaml(path)
		if err == nil {
			break
		}
	}
}
