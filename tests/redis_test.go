// author: wsfuyibing <websearch@163.com>
// date: 2021-02-06

package tests

import (
	"testing"
	"time"

	"github.com/uniondrug/gwf/xlog"
	"github.com/uniondrug/gwf/xredis"
)

func TestRedisInstance(t *testing.T) {
	ctx := xlog.NewTracing().UseDefault()
	lock, err := xredis.NewLock("test")

	if err != nil {
		t.Errorf("lock error: %v.", err)
		return
	}

	ok1 := lock.Set(ctx)
	if !ok1 {
		t.Errorf("lock failed.")
		return
	}

	t.Logf("lock and got: %s.", lock.Receipt)

	time.Sleep(time.Second * 5)

	lock.Unset(ctx)

	// time.Sleep(time.Second * 5)
}
