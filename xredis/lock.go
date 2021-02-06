// author: wsfuyibing <websearch@163.com>
// date: 2021-02-06

package xredis

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/uniondrug/gwf/xlog"
)

// Redis锁
//
//   ctx := xlog.NewTracing().UseDefault()
//
//   lock, err := xredis.NewLock("key")
//   if err != nil {
//       xlog.Errorfc(ctx, "create lock error - %v.", err)
//       return
//   }
//
//   if ok := lock.Set(ctx); !ok {
//       xlog.Errorfc(ctx, "lock")
//       return
//   }
//
//   defer lock.Unset(ctx)
//
type Lock struct {
	Key     string        // Redis锁/键名
	Receipt string        // Redis锁/回执
	conn    redis.Conn    // Redis连接.
	mu      *sync.RWMutex // 锁
	quit    chan bool     // 退出Ticker.
	retry   int           // 加锁重试次数
	ticked  bool          // Ticker状态.
}

// 创建锁实例.
func NewLock(key string) (*Lock, error) {
	o := &Lock{
		Key:   key,
		conn:  Config.pools.Get(),
		mu:    new(sync.RWMutex),
		quit:  make(chan bool),
		retry: 1,
	}
	return o, o.conn.Err()
}

// 设置重试次数.
func (o *Lock) Retry(retry int) {
	o.retry = retry
}

// 续期.
// 重置过期时间, 防止被锁资源因执行用时过长而提前释放, 导致锁失败.
func (o *Lock) Expiration(ctx interface{}) {
	_, err := o.conn.Do("EXPIRE", o.Key, LockExpiration)
	if err != nil {
		xlog.Warnfc(ctx, "[redis:lock] reset %s key expiration error - %v.", o.Key, err)
		return
	}
	xlog.Debugfc(ctx, "[redis:lock] reset %s key expiration succeed.", o.Key)
}

// 加锁.
func (o *Lock) Set(ctx interface{}) (succeed bool) {
	// 1. prepare.
	ct := time.Now()
	receipt := fmt.Sprintf("%d.%d.%d", ct.Unix(), ct.Nanosecond(), rand.Int31())
	// 2. send command and responsed.
	for i := 0; i < o.retry; i++ {
		// 2.1. send command
		r1, e1 := o.conn.Do("SET", o.Key, receipt, "NX", "EX", LockExpiration)
		// 2. send command error
		if e1 != nil {
			xlog.Warnfc(ctx, "[redis:lock][retry=%d.%d] set %s key lock error - %v.", i, o.retry, o.Key, e1)
			continue
		}
		// 2.2. parse response
		_, e2 := redis.String(r1, e1)
		if e2 != nil {
			xlog.Warnfc(ctx, "[redis:lock][retry=%d.%d] set %s key lock error - %v.", i, o.retry, o.Key, e2)
			continue
		}
		// 2.3. succeed.
		o.Receipt = receipt
		succeed = true
		xlog.Debugfc(ctx, "[redis:lock][retry=%d.%d] set %s key succeed.", i, o.retry, o.Key)
		break
	}
	// 3. reset lifetime
	if succeed {
		o.ticker(ctx)
	}
	return
}

// 解锁.
func (o *Lock) Unset(ctx interface{}) {
	// 发送信号以退出定时器.
	o.quit <- true
	// 删除锁资源.
	_, err := o.conn.Do("DEL", o.Key)
	if err != nil {
		xlog.Warnfc(ctx, "[redis:lock] delete %s key error - %v.", o.Key, err)
	} else {
		xlog.Debugfc(ctx, "[redis:lock] delete %s key succeed.", o.Key)
	}
	// 关闭连接.
	if err2 := o.conn.Close(); err2 != nil {
		xlog.Warnfc(ctx, "[redis:lock] release connection error - %v.", err2)
	}
}

// 定时器.
func (o *Lock) ticker(ctx interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.ticked {
		return
	}
	// ticker
	o.ticked = true
	go func() {
		xlog.Debugfc(ctx, "[redis:lock] start %s key ticker.", o.Key)
		defer xlog.Debugfc(ctx, "[redis:lock] quit %s key ticker.", o.Key)
		ticker := time.NewTicker(LockExpirationReset)
		for {
			select {
			case <-ticker.C:
				go o.Expiration(ctx)
			case <-o.quit:
				return
			}
		}
	}()
}
