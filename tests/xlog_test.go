// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package tests

import (
	"testing"

	"xorm.io/xorm"

	"gwf/xdb"
	"gwf/xlog"
)

func TestXLog(t *testing.T) {
	tracing := xlog.NewTracing().UseDefault()
	xlog.Debugfc(tracing, "debug format context")
	xlog.Infofc(tracing, "info format context")
	xlog.Warnfc(tracing, "warn format context")
	xlog.Errorfc(tracing, "error format context")
}

func TestXLogXDb(t *testing.T) {
	tracing := xlog.NewTracing().UseDefault()

	sess := xdb.MasterContext(tracing)

	_, _ = NewExampleService(sess).GetById(5)

	_ = xdb.TransactionWithSession(tracing, sess, func(ctx interface{}, sess *xorm.Session) error {
		service := NewExampleService(sess)
		_, err := service.GetById(1)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx interface{}, sess *xorm.Session) error {
		service := NewExampleService(sess)
		_, err := service.GetById(2)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx interface{}, sess *xorm.Session) error {
		service := NewExampleService(sess)
		_, err := service.GetById(3)
		if err != nil {
			return err
		}
		return nil
	})

}

type Example struct {
	SubscriptionId int64 `xorm:"SubscriptionId"`
}

func (*Example) TableName() string {
	return "mbs3_subscription"
}

type ExampleService struct {
	xdb.Service
}

func NewExampleService(sess ...*xorm.Session) *ExampleService {
	o := &ExampleService{}
	o.Use(sess...)
	return o
}

func (o *ExampleService) GetById(id int64) (*Example, error) {
	m := &Example{}
	if _, err := o.Slave().Where("SubscriptionId = ?", id).Get(m); err != nil {
		return nil, err
	}
	if m.SubscriptionId > 0 {
		return m, nil
	}
	return nil, nil
}
