// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xdb

import (
	"xorm.io/xorm"
)

// Service结构.
type Service struct {
	sess *xorm.Session
}

// 获取主库连结.
func (o *Service) Master() *xorm.Session {
	if o.sess == nil {
		return Config.engines.Master().NewSession()
	}
	return o.sess
}

// 获取从库连结.
func (o *Service) Slave() *xorm.Session {
	if o.sess == nil {
		return Config.engines.Slave().NewSession()
	}
	return o.sess
}

// 使用指定连结.
func (o *Service) Use(s ...*xorm.Session) {
	if s != nil && len(s) > 0 {
		o.sess = s[0]
	}
}
