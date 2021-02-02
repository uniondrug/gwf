// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package tests

import (
	"errors"
	"testing"

	"github.com/kataras/iris/v12"

	"github.com/uniondrug/go-iris/xapp"
	"github.com/uniondrug/go-iris/xlog"
)

func TestMainApp(t *testing.T) {
	xapp.New().
		RegisterController("/", &ExampleController{}).
		Run()
}

type ExampleController struct{}

func (o *ExampleController) GetIndex(ctx iris.Context) interface{} {
	xlog.Infofc(ctx, "action:ExampleController:Index()")
	return xapp.WithError(ctx, 1, errors.New("unknown"))
}

func (o *ExampleController) GetPanic(ctx iris.Context) interface{} {
	panic("panic")
	return xapp.WithError(ctx, 1, errors.New("unknown"))
}
