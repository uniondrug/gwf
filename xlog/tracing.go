// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

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

const (
	OpenTracing = "OpenTracingHandler"
)

type Tracing struct {
	parentSpanId string
	spanId       string
	spanOffset   int32
	spanVersion  string
	traceId      string
	method       string
	uri          string
}

func NewTracing() *Tracing {
	o := &Tracing{spanOffset: 0, spanVersion: "0"}
	o.spanId = o.UUID()
	return o
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

// 使用默认值初始化.
func (o *Tracing) UseDefault() *Tracing {
	o.traceId = o.spanId
	return o
}

// 使用IRIS初始化.
func (o *Tracing) UseIris(ctx iris.Context) {
	o.UseRequest(ctx.Request())
	ctx.Values().Set(OpenTracing, o)
}

// 使用HTTP请求初始化.
func (o *Tracing) UseRequest(req *http.Request) {
	// 1. assign parent span id
	if s := req.Header.Get(o.spanId); s != "" {
		o.parentSpanId = s
	}
	// 2. span version
	if s := req.Header.Get(o.spanVersion); s != "" {
		o.spanVersion = s
	}
	// 3. trace id
	if s := req.Header.Get(o.traceId); s != "" {
		o.traceId = s
	} else {
		o.traceId = o.spanId
	}
	// 4. http info
	o.method = req.Method
	o.uri = req.RequestURI
}

// 返回Offset值.
func (o *Tracing) incrOffset() (ci int32, ni int32) {
	i := atomic.LoadInt32(&o.spanOffset)
	return i, atomic.AddInt32(&o.spanOffset, 1)
}
