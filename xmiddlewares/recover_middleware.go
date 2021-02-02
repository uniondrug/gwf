// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xmiddlewares

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/kataras/iris/v12"

	"gwf/xlog2"
)

// Catch panic.
func RecoverMiddleware(ctx iris.Context) {
	defer func() {
		err := recover()
		// return if stopped or not panic.
		if err == nil || ctx.IsStopped() {
			return
		}
		// stack
		var msg = fmt.Sprintf("%v at %s\n", err, ctx.HandlerName())
		var sta string
		for i := 1; ; i++ {
			_, f, l, got := runtime.Caller(i)
			if !got {
				break
			}
			sta += fmt.Sprintf("%s:%d\n", f, l)
		}
		xlog2.Errorfc(ctx, msg+"\n"+sta)
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.StopExecution()
	}()
	ctx.Next()
}
