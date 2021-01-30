// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package tests

import (
	"encoding/json"
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
	se := NewExampleService()
	m, err := se.GetById(5)
	if err != nil {
		xlog.Errorf("service error - %s", err)
		return
	}

	s, _ := json.Marshal(m)
	xlog.Infof("service response - %s", s)

}
