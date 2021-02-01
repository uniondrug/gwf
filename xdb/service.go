// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xdb

import (
	"xorm.io/xorm"
)

// 根服务.
// 在Service中以匿名方式引入.
//
//   type ExampleService struct{
//       xdb.Service
//   }
//
//   func NewExampleService(s ...*xorm.Session) *ExampleService {
//       o := &ExampleService{}
//       o.Use(s...)
//       return o
//   }
//
type Service struct {
	sess *xorm.Session
}

// 获取主库连接.
// 在写方法(INSERT、UPDATE、DELETE)中调用此连接.
func (o *Service) Master() *xorm.Session {
	if o.sess == nil {
		return Config.engines.Master().NewSession()
	}
	return o.sess
}

// 获取从库连接.
// 在读方法(SELECT)中调用此连接.
func (o *Service) Slave() *xorm.Session {
	if o.sess == nil {
		return Config.engines.Slave().NewSession()
	}
	return o.sess
}

// 指定连接.
// 通过在事务中使用指定连接.
func (o *Service) Use(s ...*xorm.Session) {
	if s != nil && len(s) > 0 {
		o.sess = s[0]
	}
}
