/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/16
 * +----------------------------------------------------------------------
 * |Time: 11:39 下午
 * +----------------------------------------------------------------------
 */

package strategy

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CounterLimiter struct {
	counter      int64 // 计数器
	limit        int64 // 指定时间窗口内的允许通过的最大请求数
	intervalNano int64 // 指定的时间窗口
	unixNano     int64 // unix时间戳，单位为纳秒
}

func NewCounterLimiter(interval time.Duration, limit int64) *CounterLimiter {
	return &CounterLimiter{
		counter:      0,
		limit:        limit,
		intervalNano: int64(interval),
		unixNano:     time.Now().UnixNano(),
	}
}

func (counter *CounterLimiter) Allow() (isAllow bool) {
	nowUnixNano := time.Now().UnixNano()
	if nowUnixNano-counter.unixNano > counter.intervalNano {
		fmt.Println("counter:", counter.counter)
		atomic.StoreInt64(&counter.counter, 1)
		atomic.StoreInt64(&counter.unixNano, nowUnixNano)
		return true
	}
	atomic.AddInt64(&counter.counter, 1)
	return counter.counter <= counter.limit
}
