// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package xlog

import (
	"fmt"

	"xorm.io/xorm/log"
)

type XDBLogger struct{}

func (o *XDBLogger) Debugf(format string, v ...interface{}) {}
func (o *XDBLogger) Errorf(format string, v ...interface{}) {}
func (o *XDBLogger) Infof(format string, v ...interface{})  {}
func (o *XDBLogger) Warnf(format string, v ...interface{})  {}
func (o *XDBLogger) Level() log.LogLevel                    { return log.LOG_INFO }
func (o *XDBLogger) SetLevel(l log.LogLevel)                {}
func (o *XDBLogger) ShowSQL(show ...bool)                   {}
func (o *XDBLogger) IsShowSQL() bool                        { return true }
func (o *XDBLogger) BeforeSQL(c log.LogContext)             {}
func (o *XDBLogger) AfterSQL(c log.LogContext) {
	var ctx *Tracing
	if c.Ctx != nil {
		if x := c.Ctx.Value(OpenTracing); x != nil {
			ctx, _ = x.(*Tracing)
		}
	}
	// add INFO log.
	if enableInfo {
		if c.Args != nil && len(c.Args) > 0 {
			Config.log(ctx, InfoLevel, fmt.Sprintf("[SQL][d=%f] %s - %v.", c.ExecuteTime.Seconds(), c.SQL, c.Args))
		} else {
			Config.log(ctx, InfoLevel, fmt.Sprintf("[SQL][d=%f] %s.", c.ExecuteTime.Seconds(), c.SQL))
		}
	}
	// add ERROR log.
	if c.Err != nil && enableError {
		Config.log(ctx, ErrorLevel, fmt.Sprintf("[SQL] %s.", c.Err.Error()))
	}
}
