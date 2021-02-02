// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package xapp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/uniondrug/gwf/xlog"
	"github.com/uniondrug/gwf/xmiddlewares"
)

type Application struct {
	app *iris.Application
}

func New() *Application {
	xlog.Debugf("[application] %s created", Config.Name)
	o := &Application{}
	o.app = iris.New()
	o.app.Use(xmiddlewares.TracingMiddleware)
	o.app.Use(xmiddlewares.RecoverMiddleware)
	o.app.OnAnyErrorCode(ErrorCode)
	o.app.Logger().SetLevel("disable")
	return o
}

func (o *Application) RegisterController(path string, handler interface{}) *Application {
	mvc.Configure(o.app.Party(path), func(m *mvc.Application) {
		m.Handle(handler)
	})
	return o
}

func (o *Application) Run() {
	if err := o.app.Run(iris.Addr(Config.Addr)); err != nil {
		xlog.Errorf("%v", err)
	}
}
