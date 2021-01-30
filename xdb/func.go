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

// Return context for xdb, cross tracing.
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

// Return master connection session.
func Master() *xorm.Session {
	return Config.engines.Master().NewSession()
}

// Return master connection session and set context.
func MasterContext(ctx interface{}) *xorm.Session {
	sess := Master()
	Context(sess, ctx)
	return sess
}

// Return slave connection session.
func Slave() *xorm.Session {
	return Config.engines.Slave().NewSession()
}

// Return slave connection session and set context.
func SlaveContext(ctx interface{}) *xorm.Session {
	sess := Slave()
	Context(sess, ctx)
	return sess
}

// Transaction.
//
//   if err, done := xdb.TransactionWithSession(func(sess *xorm.Session) error {
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
func Transaction(handlers ...func(sess *xorm.Session) error) (err error, done bool) {
	return TransactionWithSession(nil, handlers...)
}

// Transaction with session.
//
// 1st returned param - rollback or not executed if error returned.
// 2nd returned param - commit or rollback status.
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
