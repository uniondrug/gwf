// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xlog

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

// 请求链.
type Tracing struct {
	offset       int32
	parentSpanId string
	spanId       string
	spanVersion  string
	traceId      string
}

// 创建请求链结构体.
func NewTracing() *Tracing {
	tracing := &Tracing{offset: 0, spanVersion: "0"}
	tracing.spanId = tracing.UUID()
	tracing.spanVersion = "0"
	return tracing
}

// Offset偏移量+1.
func (o *Tracing) Increment() int32 {
	i := o.offset
	atomic.AddInt32(&o.offset, 1)
	return i
}

// 上级链路标识.
func (o *Tracing) ParentSpanId() string {
	return o.parentSpanId
}

// 本级链路标识.
func (o *Tracing) SpanId() string {
	return o.spanId
}

// 本级链路版本.
func (o *Tracing) SpanVersion() string {
	return o.spanVersion
}

// 获取UUID.
func (o *Tracing) UUID() string {
	// 1. 通过UUID包获取.
	if u, e := uuid.NewUUID(); e == nil {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	// 2. 获取失败时创建随机值.
	t := time.Now()
	return fmt.Sprintf("a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
}

// 向Request追加.
func (o *Tracing) AssignRequest(req *http.Request, offset int) {
	req.Header.Set(Config.ParentSpanId, o.parentSpanId)
	req.Header.Set(Config.TraceId, o.traceId)
	req.Header.Set(Config.SpanVersion, fmt.Sprintf("%s.%d", o.spanVersion, offset))
}

// 向Response追加.
func (o *Tracing) AssignResponse(req iris.Context) {
	req.ResponseWriter().Header().Set(Config.TraceId, o.traceId)
}

// 基于IRIS.
func (o *Tracing) FromIris(req iris.Context) *Tracing {
	return o.FromRequest(req.Request())
}

// 基于HTTP请求.
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

// 生成根链路.
func (o *Tracing) FromRoot() *Tracing {
	o.traceId = o.spanId
	return o
}
