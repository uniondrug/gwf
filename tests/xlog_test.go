// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package tests

import (
	"encoding/json"
	"testing"
	"time"

	"gwf/xlog"
)

func TestLogger(t *testing.T) {
	// xlog.Config.SetHandler(func(line *xlog.Line, done chan bool) {
	// 	println("line: ", line.Message())
	// })

	t.Logf("---- logger ----")
	bs, _ := json.Marshal(xlog.Config)
	t.Logf("     Adapter: %s", bs)
	t.Logf("---- finish ----")

	xlog.Debug("Debug message!")
	xlog.Info("Info message!")
	xlog.Warn("Warn message!")
	xlog.Error("Error message!")
	xlog.Alert("Alert message!")

	time.Sleep(time.Second)
}

func TestLoggerForTracing(t *testing.T) {
	trace := xlog.NewTracing().FromRoot()
	data,_ := json.Marshal(trace)
	println("trace: ", string(data))

}
