// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog2

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// Tracing struct.
type Tracing struct {
	offset       int32
	parentSpanId string
	spanId       string
	spanVersion  string
	traceId      string
}

// Create tracing struct.
func NewTracing() *Tracing {
	tracing := &Tracing{offset: 0, spanVersion: "0"}
	tracing.spanId = tracing.UUID()
	tracing.spanVersion = "0"
	return tracing
}

// Increment offset.
func (o *Tracing) Increment() int32 {
	i := o.offset
	atomic.AddInt32(&o.offset, 1)
	return i
}

// Parent span id.
func (o *Tracing) ParentSpanId() string {
	return o.parentSpanId
}

// Current span id.
func (o *Tracing) SpanId() string {
	return o.spanId
}

// Current span version.
func (o *Tracing) SpanVersion() string {
	return o.spanVersion
}

// Return unique identify.
func (o *Tracing) UUID() string {
	// 1. 通过UUID包获取.
	if u, e := uuid.NewUUID(); e == nil {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	// 2. 获取失败时创建随机值.
	t := time.Now()
	return fmt.Sprintf("a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
}

// Assign to request.
func (o *Tracing) AssignRequest(req *http.Request, offset int) {
	req.Header.Set(Config.ParentSpanId, o.parentSpanId)
	req.Header.Set(Config.TraceId, o.traceId)
	req.Header.Set(Config.SpanVersion, fmt.Sprintf("%s.%d", o.spanVersion, offset))
}

// Assign to response.
func (o *Tracing) AssignResponse(req iris.Context) {
	req.ResponseWriter().Header().Set(Config.TraceId, o.traceId)
}

// Trace from iris.
func (o *Tracing) FromIris(ctx iris.Context) *Tracing {
	return o.FromRequest(ctx.Request())
}

// Trace form http request.
func (o *Tracing) FromRequest(req *http.Request) *Tracing {
	// 读取TraceID, 如为空则使用SpanId.
	if s := req.Header.Get(Config.TraceId); s != "" {
		o.traceId = s
	} else {
		o.traceId = o.spanId
	}
	// 上级SpanId.
	if s := req.Header.Get(Config.ParentSpanId); s != "" {
		o.parentSpanId = s
	}
	// 本级Version.
	if s := req.Header.Get(Config.SpanVersion); s != "" {
		o.spanVersion = s
	}
	return o
}

// Trace from root.
func (o *Tracing) FromRoot() *Tracing {
	o.traceId = o.spanId
	return o
}
