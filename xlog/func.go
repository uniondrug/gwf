// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

func Debug(text string) {
	if Config.DebugOn() {
		Config.log(nil, DebugLevel, text)
	}
}

func Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		Config.log(nil, DebugLevel, text, args...)
	}
}

func Debugfc(ctx interface{}, text string, args ...interface{}) {
	if Config.DebugOn() {
		Config.log(ctx, DebugLevel, text, args...)
	}
}

func Info(text string) {
	if Config.InfoOn() {
		Config.log(nil, InfoLevel, text)
	}
}

func Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		Config.log(nil, InfoLevel, text, args...)
	}
}

func Infofc(ctx interface{}, text string, args ...interface{}) {
	if Config.InfoOn() {
		Config.log(ctx, InfoLevel, text, args...)
	}
}

func Warn(text string) {
	if Config.WarnOn() {
		Config.log(nil, WarnLevel, text)
	}
}

func Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		Config.log(nil, WarnLevel, text, args...)
	}
}

func Warnfc(ctx interface{}, text string, args ...interface{}) {
	if Config.WarnOn() {
		Config.log(ctx, WarnLevel, text, args...)
	}
}

func Error(text string) {
	if Config.ErrorOn() {
		Config.log(nil, ErrorLevel, text)
	}
}

func Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		Config.log(nil, ErrorLevel, text, args...)
	}
}

func Errorfc(ctx interface{}, text string, args ...interface{}) {
	if Config.ErrorOn() {
		Config.log(ctx, ErrorLevel, text, args...)
	}
}
