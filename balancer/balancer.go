/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/13
 * +----------------------------------------------------------------------
 * |Time: 10:00 下午
 * +----------------------------------------------------------------------
 */

package balancer

import (
	"github.com/gongxulei/go_kit/balancer/base"
	"github.com/gongxulei/go_kit/balancer/strategy"
)


type Strategy string

const (
	RANDOM           Strategy = "random"
	ROUND            Strategy = "round"
	CONSISTENCY_HASH Strategy = "consistency_hash"
)

func NewBalancer(addrInfoList []*base.AddrInfo, mode Strategy) (balancer base.Balancer) {
	switch mode {
	case RANDOM:
		return strategy.NewRandomBalancer(addrInfoList)
	case ROUND:
		return strategy.NewRoundRobinBalancer(addrInfoList)
	default:
		return strategy.NewRandomBalancer(addrInfoList)
	}
}
