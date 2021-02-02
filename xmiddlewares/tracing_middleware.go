// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xmiddlewares

import (
	"github.com/kataras/iris/v12"

	"github.com/uniondrug/gwf/xlog"
)

// 追加OpenTracing控制.
// 本方法用于在创建IRIS服务时, 在最顶层引入.
func TracingMiddleware(ctx iris.Context) {
	xlog.NewTracing().UseIris(ctx)
	xlog.Debugfc(ctx, "middleware:tracing")
	ctx.Next()
}
