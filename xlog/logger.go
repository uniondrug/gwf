// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package xlog

var (
	enableDebug = false
	enableInfo  = false
	enableWarn  = false
	enableError = false
	enableAlert = false
)

// 日志结构.
type Logger struct{}

func (o *Logger) Debug(text string) {
	if enableDebug {
		Config.log(nil, DebugLevel, text)
	}
}

func (o *Logger) Debugf(format string, args ...interface{}) {
	if enableDebug {
		Config.log(nil, DebugLevel, format, args...)
	}
}

func (o *Logger) Debugfc(ctx interface{}, format string, args ...interface{}) {
	if enableDebug {
		Config.log(ctx, DebugLevel, format, args...)
	}
}

func (o *Logger) Info(text string) {
	if enableInfo {
		Config.log(nil, InfoLevel, text)
	}
}

func (o *Logger) Infof(format string, args ...interface{}) {
	if enableInfo {
		Config.log(nil, InfoLevel, format, args...)
	}
}

func (o *Logger) Infofc(ctx interface{}, format string, args ...interface{}) {
	if enableInfo {
		Config.log(ctx, InfoLevel, format, args...)
	}
}

func (o *Logger) Warn(text string) {
	if enableWarn {
		Config.log(nil, WarnLevel, text)
	}
}

func (o *Logger) Warnf(format string, args ...interface{}) {
	if enableWarn {
		Config.log(nil, WarnLevel, format, args...)
	}
}

func (o *Logger) Warnfc(ctx interface{}, format string, args ...interface{}) {
	if enableWarn {
		Config.log(ctx, WarnLevel, format, args...)
	}
}

func (o *Logger) Error(text string) {
	if enableError {
		Config.log(nil, ErrorLevel, text)
	}
}

func (o *Logger) Errorf(format string, args ...interface{}) {
	if enableError {
		Config.log(nil, ErrorLevel, format, args...)
	}
}

func (o *Logger) Errorfc(ctx interface{}, format string, args ...interface{}) {
	if enableError {
		Config.log(ctx, ErrorLevel, format, args...)
	}
}

func (o *Logger) Alert(text string) {
	if enableAlert {
		Config.log(nil, AlertLevel, text)
	}
}

func (o *Logger) Alertf(format string, args ...interface{}) {
	if enableAlert {
		Config.log(nil, AlertLevel, format, args...)
	}
}

func (o *Logger) Alertfc(ctx interface{}, format string, args ...interface{}) {
	if enableAlert {
		Config.log(ctx, AlertLevel, format, args...)
	}
}

func (o *Logger) Level() LogLevel {
	return Config.Level()
}

func (o *Logger) SetLevel(level LogLevel) {
	Config.SetLevel(level)
}
