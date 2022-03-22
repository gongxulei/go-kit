/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/15
 * +----------------------------------------------------------------------
 * |Time: 2:08 下午
 * +----------------------------------------------------------------------
 */

package test

import (
	"fmt"
	"github.com/gongxulei/go_kit/balancer/base"
	"github.com/gongxulei/go_kit/balancer/strategy"
	"testing"
)

func TestRandom_Balance(t *testing.T) {
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
	random := strategy.NewRandomBalancer(addrInfoList)

	for i := 0; i < 10; i++ {
		addr := random.Balance()
		fmt.Println("random address: ", addr)
	}
}
