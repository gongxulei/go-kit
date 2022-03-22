/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/17
 * +----------------------------------------------------------------------
 * |Time: 12:27 上午
 * +----------------------------------------------------------------------
 */

package strategy

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestNewCounterLimiter(t *testing.T) {
	limiter := NewCounterLimiter(1*time.Second, 200)
	cpu := runtime.NumCPU()
	na := time.Now().UnixNano()
	fmt.Println("time:", na)
	runtime.GOMAXPROCS(cpu)
	wg := &sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(i))
			limiter.Allow()
			// fmt.Printf("i:%d, res: %v\n", i, )
		}(i, wg)
	}
	wg.Wait()
	fmt.Println(limiter)
	fmt.Println(limiter.unixNano - na)
}

// 10.006332
