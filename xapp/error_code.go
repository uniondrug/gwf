// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package xapp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"

	"github.com/uniondrug/gwf/xlog"
)

func ErrorCode(ctx iris.Context) {
	str := fmt.Sprintf("HTTP %d %s", ctx.GetStatusCode(), http.StatusText(ctx.GetStatusCode()))
	err := errors.New(str)
	data := WithError(ctx, ctx.GetStatusCode(), err)
	xlog.Warnfc(ctx, str)
	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(data)
}
