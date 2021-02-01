// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package tests

import (
	"testing"

	"xorm.io/xorm"

	"gwf/xdb"
	"gwf/xlog"
)

// Framework model.
type Example struct {
	SubscriptionId int `xorm:"pk autoincr"`
}

func (o *Example) TableName() string {
	return "mbs3_subscription"
}

// Framework service.
type ExampleService struct {
	xdb.Service
}

func NewExampleService(s ...*xorm.Session) *ExampleService {
	o := &ExampleService{}
	o.Use(s...)
	return o
}

func (o *ExampleService) GetById(id int) (*Example, error) {
	m := &Example{}
	if _, err := o.Slave().Where("SubscriptionId = ?", id).Get(m); err != nil {
		return nil, err
	}
	if m.SubscriptionId > 0 {
		return m, nil
	}
	return nil, nil
}

func TestXDb(t *testing.T) {

	ss := xdb.MasterContext(nil)

	if se, _ := xdb.TransactionWithSession(ss, func(sess *xorm.Session) error {
		se := NewExampleService(ss)
		_, err := se.GetById(4)
		return err
	}, func(sess *xorm.Session) error {
		panic("panic transaction")
	}, func(sess *xorm.Session) error {
		NewExampleService(sess).GetById(5)
		return nil
	}, func(sess *xorm.Session) error {
		NewExampleService(sess).GetById(6)
		return nil
	}); se != nil {
		xlog.Errorf("transaction error - %s", se)
	}


}
