/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/13
 * +----------------------------------------------------------------------
 * |Time: 9:59 下午
 * +----------------------------------------------------------------------
 */

package strategy

import (
	"github.com/gongxulei/go_kit/balancer/base"
	"math/rand"
	"time"
)

type randomBalancer struct {
	addrInfoList      []*base.AddrInfo
	addrList          []string
	weightList        []uint8
	criticalValueList []int // critical 临界值
	criticalValue     int
}

func (rd *randomBalancer) Balance() (addr string) {
	// 获取一个随机值
	rand.Seed(time.Now().UnixNano())
	randValue := rand.Intn(rd.criticalValue)
	// 判断随机值所在范围
	for i, critical := range rd.criticalValueList {
		if i == 0 {
			continue
		}
		lastIndex := i - 1
		if randValue >= rd.criticalValueList[lastIndex] && randValue < critical {
			return rd.addrList[lastIndex]
		}
	}
	return rd.addrList[0]
}

func NewRandomBalancer(addrInfoList []*base.AddrInfo) (random *randomBalancer) {
	random = new(randomBalancer)
	random.criticalValueList = []int{0}
	random.weightList = make([]uint8, 0)
	random.addrList = make([]string, 0)
	// random.criticalValueList
	for _, addrInfo := range addrInfoList {
		random.weightList = append(random.weightList, addrInfo.Weight)
		random.addrList = append(random.addrList, addrInfo.Addr)
		random.criticalValue = random.criticalValue + int(addrInfo.Weight)
		random.criticalValueList = append(random.criticalValueList, random.criticalValue)
	}
	return
}
