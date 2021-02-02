// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// 日志配置.
type Configuration struct {
	AdapterName string `yaml:"adapter-name"`
	LevelName   string `yaml:"level-name"`
	TimeFormat  string `yaml:"time-format"`
	adapter     LogAdapter
	handler     LogHandler
	level       LogLevel
	_debug      bool
	_info       bool
	_warn       bool
	_error      bool
}

// 创建日志配置结构体.
func NewConfiguration() *Configuration {
	o := &Configuration{}
	o.defaults()
	return o
}

// ///////////////////////////////////////////////////////////
// Status 													//
// ///////////////////////////////////////////////////////////

func (o *Configuration) DebugOn() bool { return o._debug }
func (o *Configuration) ErrorOn() bool { return o._error }
func (o *Configuration) InfoOn() bool  { return o._info }
func (o *Configuration) WarnOn() bool  { return o._warn }

// 从YAML中加载配置.
func (o *Configuration) LoadYaml(path string) error {
	bs, e1 := ioutil.ReadFile(path)
	// 1. 读文件出错.
	if e1 != nil {
		return e1
	}
	// 2. 解析YAML出错.
	if e2 := yaml.Unmarshal(bs, o); e2 != nil {
		return e2
	}
	// 3. 更新默认值.
	if o.LevelName != "" {
		pe := true
		ln := strings.ToUpper(o.LevelName)
		for l, n := range LevelTexts {
			if n == ln {
				o.level = l
				pe = false
				break
			}
		}
		if pe {
			panic("unknown logger level name - " + o.LevelName)
		}
	}
	o._debug = o.level >= DebugLevel
	o._info = o.level >= InfoLevel
	o._warn = o.level >= WarnLevel
	o._error = o.level >= ErrorLevel
	// 4. 更新适配器.
	if o.AdapterName != "" {
		pe := true
		an := strings.ToLower(o.AdapterName)
		for a, n := range AdapterTexts {
			if n == an {
				o.adapter = a
				pe = false
				break
			}
		}
		if pe {
			panic("unknown logger adapter name - " + o.AdapterName)
		}
	}
	switch o.adapter {
	case FileAdapter:
		o.handler = NewAdapterFileHandler().Handler
	case RedisAdapter:
		o.handler = NewAdapterRedisHandler().Handler
	case TermAdapter:
		o.handler = NewAdapterTermHandler().Handler
	default:
		panic("unknown adapter")
	}
	// 5. 时间戳格式.
	if o.TimeFormat == "" {
		o.TimeFormat = DefaultTimeFormat
	}
	// 6. 完成.
	return nil
}

// 设置Handler回调.
func (o *Configuration) SetHandler(handler LogHandler) {
	o.handler = handler
}

// 默认值.
func (o *Configuration) defaults() {
	// 预置字段.
	o.adapter = DefaultAdapter
	o.level = DefaultLevel
	o.TimeFormat = DefaultTimeFormat
	// 解析YAML文件.
	for _, path := range []string{"./config/logger.yaml", "../config/logger.yaml"} {
		if o.LoadYaml(path) == nil {
			break
		}
	}
}

// 发送日志.
// ctx 支持 context.Context, iris.Context, *xlog.Tracing.
func (o *Configuration) log(ctx interface{}, level LogLevel, text string, args ...interface{}) {
	Config.handler(NewLine(ctx, level, text, args...))
}
