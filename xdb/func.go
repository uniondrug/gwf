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

// 向Session注入Context上下文.
// param sess is XORM connection session.
// param x accept xdb.Tracing, iris.Context, context.Context.
func Context(sess *xorm.Session, x interface{}) {
	// use *Tracing for cross.
	if t1, o1 := x.(*xlog.Tracing); o1 && t1 != nil {
		sess.Context(context.WithValue(rootContext, xlog.OpenTracing, t1))
		return
	}
	// use iris.Context.
	if i2, o2 := x.(iris.Context); o2 && i2 != nil {
		if g2 := i2.Values().Get(xlog.OpenTracing); g2 != nil {
			if t2, ok := g2.(*xlog.Tracing); ok && t2 != nil {
				sess.Context(context.WithValue(rootContext, xlog.OpenTracing, t2))
				return
			}
		}
	}
	// use context.Context.
	if i3, o3 := x.(context.Context); o3 && i3 != nil {
		if g3 := i3.Value(xlog.OpenTracing); g3 != nil {
			if t3, ok := g3.(*xlog.Tracing); ok && t3 != nil {
				sess.Context(context.WithValue(rootContext, xlog.OpenTracing, t3))
				return
			}
		}
	}
	// use DEFAULT.
	sess.Context(context.WithValue(rootContext, xlog.OpenTracing, xlog.NewTracing().FromRoot()))
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
func Transaction(handlers ...func(sess *xorm.Session) error) (err error, done bool) {
	return TransactionWithSession(nil, handlers...)
}

// 执行事务.
//
// 在事务中必须保证使用同一个连接, 且各回调以串行方式执行.
//
// err - 执行执行不成功时, 返回error类型结构, 反之正常执行.
//
// done - 事务commit/rollback状态.
//
//   sess := xdb.Master()
//
//   if err, done := xdb.TransactionWithSession(sess, func(sess *xorm.Session) error {
//
//       // logic
//
//   }, func(sess *xorm.Session) error {
//
//       // logic
//
//   }, func(sess *xorm.Session) error {
//
//       // logic
//
//   }); err != nil {
//
//       println("Transaction error - ", err.Error())
//       println("Rollback status - ", done)
//
//   }
//
func TransactionWithSession(sess *xorm.Session, handlers ...func(sess *xorm.Session) error) (err error, done bool) {
	// open master session if not specified.
	if sess == nil {
		sess = Master()
		if sess == nil {
			err = errors.New("can not get master connection session")
			return
		}
	}
	// ping connection.
	if err = sess.Ping(); err != nil {
		return
	}
	// begin transaction.
	if err = sess.Begin(); err != nil {
		return
	}
	// invoke when handlers executed.
	defer func() {
		// reset error if panic fired.
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		// rollback.
		if err != nil {
			done = sess.Rollback() == nil
		} else {
			done = sess.Commit() == nil
		}
	}()
	// loop handlers.
	// break if error returned when handler execute.
	for _, handler := range handlers {
		if err = handler(sess); err != nil {
			break
		}
	}
	// end loop
	return
}
