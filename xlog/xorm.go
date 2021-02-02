// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xlog

import (
	"fmt"

	"xorm.io/xorm/log"
)

// 用于XORM日志.
type XOrmLog struct{}

// 创建用于XORM的Logger.
func NewXOrmLog() *XOrmLog {
	o := &XOrmLog{}
	return o
}

func (o *XOrmLog) Debugf(format string, v ...interface{}) {}
func (o *XOrmLog) Errorf(format string, v ...interface{}) {}
func (o *XOrmLog) Infof(format string, v ...interface{})  {}
func (o *XOrmLog) Warnf(format string, v ...interface{})  {}
func (o *XOrmLog) Level() log.LogLevel                    { return log.LOG_INFO }
func (o *XOrmLog) SetLevel(l log.LogLevel)                {}
func (o *XOrmLog) ShowSQL(show ...bool)                   {}
func (o *XOrmLog) IsShowSQL() bool                        { return true }
func (o *XOrmLog) BeforeSQL(c log.LogContext)             {}

// SQL完成后写入Log.
func (o *XOrmLog) AfterSQL(c log.LogContext) {
	// add INFO log.
	if Config.InfoOn() {
		if c.Args != nil && len(c.Args) > 0 {
			Config.log(c.Ctx, InfoLevel, fmt.Sprintf("[SQL][d=%f] %s - %v.", c.ExecuteTime.Seconds(), c.SQL, c.Args))
		} else {
			Config.log(c.Ctx, InfoLevel, fmt.Sprintf("[SQL][d=%f] %s.", c.ExecuteTime.Seconds(), c.SQL))
		}
	}
	// add ERROR log.
	if c.Err != nil && Config.ErrorOn() {
		Config.log(c.Ctx, ErrorLevel, fmt.Sprintf("[SQL] %s.", c.Err.Error()))
	}
}
