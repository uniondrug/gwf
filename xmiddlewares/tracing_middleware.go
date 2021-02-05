// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xmiddlewares

import (
	"regexp"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/uniondrug/gwf/xlog"
)

// 追加OpenTracing控制.
// 本方法用于在创建IRIS服务时, 在最顶层引入.
func TracingMiddleware(ctx iris.Context) {
	// 初始化请求链.
	tracing := xlog.NewTracing()
	tracing.UseIris(ctx)
	ctx.Header(xlog.Config.TraceId, tracing.TraceId())
	// 入参写入Log.
	if bs, err := context.GetBody(ctx.Request(), true); err == nil {
		str := regexp.MustCompile(`\s*:\s*`).ReplaceAllString(string(bs), ":")
		str = regexp.MustCompile(`\n\s*`).ReplaceAllString(str, "")
		xlog.Infofc(ctx, "原始入参: %s.", str)
	}
	// 中间件链路.
	ctx.Next()
}
