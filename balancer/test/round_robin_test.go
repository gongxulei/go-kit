/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/15
 * +----------------------------------------------------------------------
 * |Time: 3:20 下午
 * +----------------------------------------------------------------------
 */

package test

import (
	"fmt"
	"github.com/gongxulei/go_kit/balancer/base"
	"github.com/gongxulei/go_kit/balancer/strategy"
	"testing"
	"time"
)

func TestRoundRobin_Balance(t *testing.T) {
	addrInfoList := make([]*base.AddrInfo, 0)
	addrInfoList = append(addrInfoList, &base.AddrInfo{
		Addr:   "127.0.0.1:9000",
		Weight: 1,
	})
	addrInfoList = append(addrInfoList, &base.AddrInfo{
		Addr:   "127.0.0.1:8000",
		Weight: 2,
	})
	addrInfoList = append(addrInfoList, &base.AddrInfo{
		Addr:   "127.0.0.1:7000",
		Weight: 1,
	})

	// 权重值不能超过10
	// 如果不使用权重可将每个权重设置为相等的数字
	round := strategy.NewRoundRobinBalancer(addrInfoList)

	for i := 0; i < 100; i++ {
		go func() {
			addr := round.Balance()
			fmt.Println("round address: ", addr)
		}()
	}
	time.Sleep(10 *time.Second)
}
