/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/15
 * +----------------------------------------------------------------------
 * |Time: 2:51 下午
 * +----------------------------------------------------------------------
 */

package strategy

import (
	"github.com/gongxulei/go_kit/balancer/base"
	"sync/atomic"
)

const (
	RoundRobinMaxValue = 1 << 30
)

type RoundRobin struct {
	AddrInfoList []*base.AddrInfo
	BalancerAddr []string
	RoundValue   *int64
}

func (round *RoundRobin) Balance() (addr string) {
	// 使用原子操作读取轮训次数，并将更新轮训次数
	for {
		roundValue := atomic.LoadInt64(round.RoundValue)
		// fmt.Println("roundValue: ", roundValue, " RoundRobinMaxValue: ", RoundRobinMaxValue)
		if !atomic.CompareAndSwapInt64(round.RoundValue, roundValue, roundValue+1) {
			// 设置失败重新设置
			// fmt.Println("设置CompareAndSwapInt64 false")
			continue
		}
		index := int(roundValue+1) % len(round.BalancerAddr)
		// 超过最大轮询次数进行重置
		if roundValue+1 >= RoundRobinMaxValue {
			atomic.StoreInt64(round.RoundValue, 0)
		}
		return round.BalancerAddr[index]
	}
}

func NewRoundRobinBalancer(addrInfoList []*base.AddrInfo) (roundRobin *RoundRobin) {
	var start int64
	roundRobin = new(RoundRobin)
	roundRobin.AddrInfoList = addrInfoList
	for _, addrInfo := range addrInfoList {
		if addrInfo.Weight == 0 {
			roundRobin.BalancerAddr = append(roundRobin.BalancerAddr, addrInfo.Addr)
			continue
		}
		for i := 0; i < int(addrInfo.Weight); i++ {
			roundRobin.BalancerAddr = append(roundRobin.BalancerAddr, addrInfo.Addr)
		}
	}
	roundRobin.RoundValue = &start
	return
}
