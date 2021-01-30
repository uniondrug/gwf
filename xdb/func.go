// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xdb

import (
	"errors"
	"fmt"

	"xorm.io/xorm"
)

// 获取主库连结.
func Master() *xorm.Session {
	return Config.engines.Master().NewSession()
}

// 获取从库连结.
func Slave() *xorm.Session {
	return Config.engines.Slave().NewSession()
}

// 指定连续执行事务.
// 第1返回值: 为error类型时, 表示事务未执行或最终回滚.
// 第2返回值: 表示事务提交、回滚状态.
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

// 指定连结执行事务.
// 第1返回值: 为error类型时, 表示事务未执行或最终回滚.
// 第2返回值: 表示事务提交、回滚状态.
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
