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
	tracing := xlog.NewTracing()
	tracing.UseIris(ctx)
	ctx.Header(xlog.Config.TraceId, tracing.TraceId())
	ctx.Next()
}
