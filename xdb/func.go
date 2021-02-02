// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"

	"gwf/xlog"
)

type TransactionHandler func(ctx interface{}, sess *xorm.Session) error

// 向Session注入Context上下文.
// param sess is XORM connection session.
// param x accept xdb.Tracing, iris.Context, context.Context.
func Context(sess *xorm.Session, x interface{}) {
	// create if nil.
	if x == nil {
		sess.Context(context.WithValue(rootContext, xlog.OpenTracing, xlog.NewTracing().UseDefault()))
		return
	}
	// xlog.Tracing.
	if t, ok := x.(*xlog.Tracing); ok && t != nil {
		sess.Context(context.WithValue(rootContext, xlog.OpenTracing, t))
		return
	}
	// iris.Context.
	if c, ok := x.(iris.Context); ok && c != nil {
		sess.Context(context.WithValue(rootContext, xlog.OpenTracing, c.Values().Get(xlog.OpenTracing)))
		return
	}
	// context.Context.
	if c, ok := x.(context.Context); ok && c != nil {
		sess.Context(c)
		return
	}
}

// 读取主库连接.
func Master() *xorm.Session {
	return Config.engines.Master().NewSession()
}

// 读取主库连接.
func MasterContext(ctx interface{}) *xorm.Session {
	sess := Master()
	Context(sess, ctx)
	return sess
}

// 读取从库连接.
func Slave() *xorm.Session {
	return Config.engines.Slave().NewSession()
}

// 读取从库连接.
func SlaveContext(ctx interface{}) *xorm.Session {
	sess := Slave()
	Context(sess, ctx)
	return sess
}

// 执行事务.
func Transaction(ctx interface{}, handlers ...TransactionHandler) (err error) {
	return TransactionWithSession(ctx, nil, handlers...)
}

// 执行事务.
//
// 在事务中必须保证使用同一个连接, 且各回调以串行方式执行.
//
// err - 执行不成功, 返回error类型结构, 反之正常执行.
//
//   tracing := xdb.Master()
//   sess := xdb.MasterContext(tracing)
//   if err := xdb.TransactionWithSession(tracing, sess, func(ctx interface{}, sess *xorm.Session) error {
//       // logic
//   }, func(sess *xorm.Session) error {
//       // logic
//   }, func(sess *xorm.Session) error {
//       // logic
//   }); err != nil {
//       println("Transaction error - ", err.Error())
//   }
//
func TransactionWithSession(ctx interface{}, sess *xorm.Session, handlers ...TransactionHandler) (err error) {
	// 校验连接.
	// 若未指定连接, 则自动选择主库连接.
	if sess == nil {
		sess = Master()
	}
	// 校验状态.
	if err = sess.Ping(); err != nil {
		return
	}
	// 打开事务.
	if err = sess.Begin(); err != nil {
		return
	}
	// 完成事务.
	defer func() {
		// 捕获异常.
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		// 结束事务.
		if err != nil {
			// 回滚事务.
			_ = sess.Rollback()
		} else {
			// 提交事务.
			_ = sess.Commit()
		}
	}()
	// 遍历.
	for _, handler := range handlers {
		if err = handler(ctx, sess); err != nil {
			break
		}
	}
	// 结束.
	return
}
