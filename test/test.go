/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/24
 * +----------------------------------------------------------------------
 * |Time: 11:08 上午
 * +----------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const mutexLocked = 1 << iota

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked)
}

func main() {
	var m Mutex
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("TryLock: %t\n", m.TryLock()) //false
		}()
	}
	defer m.Unlock()
	time.Sleep(time.Minute * 10)
}
