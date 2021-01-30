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

func Debug(text string) {
	if enableDebug {
		Config.log(nil, DebugLevel, text)
	}
}

func Debugf(format string, args ...interface{}) {
	if enableDebug {
		Config.log(nil, DebugLevel, format, args...)
	}
}

func Debugfc(ctx interface{}, format string, args ...interface{}) {
	if enableDebug {
		Config.log(ctx, DebugLevel, format, args...)
	}
}

func Info(text string) {
	if enableInfo {
		Config.log(nil, InfoLevel, text)
	}
}

func Infof(format string, args ...interface{}) {
	if enableInfo {
		Config.log(nil, InfoLevel, format, args...)
	}
}

func Infofc(ctx interface{}, format string, args ...interface{}) {
	if enableInfo {
		Config.log(ctx, InfoLevel, format, args...)
	}
}

func Warn(text string) {
	if enableWarn {
		Config.log(nil, WarnLevel, text)
	}
}

func Warnf(format string, args ...interface{}) {
	if enableWarn {
		Config.log(nil, WarnLevel, format, args...)
	}
}

func Warnfc(ctx interface{}, format string, args ...interface{}) {
	if enableWarn {
		Config.log(ctx, WarnLevel, format, args...)
	}
}

func Error(text string) {
	if enableError {
		Config.log(nil, ErrorLevel, text)
	}
}

func Errorf(format string, args ...interface{}) {
	if enableError {
		Config.log(nil, ErrorLevel, format, args...)
	}
}

func Errorfc(ctx interface{}, format string, args ...interface{}) {
	if enableError {
		Config.log(ctx, ErrorLevel, format, args...)
	}
}

func Alert(text string) {
	if enableAlert {
		Config.log(nil, AlertLevel, text)
	}
}

func Alertf(format string, args ...interface{}) {
	if enableAlert {
		Config.log(nil, AlertLevel, format, args...)
	}
}

func Alertfc(ctx interface{}, format string, args ...interface{}) {
	if enableAlert {
		Config.log(ctx, AlertLevel, format, args...)
	}
}
