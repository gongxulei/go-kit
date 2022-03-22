/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/12
 * +----------------------------------------------------------------------
 * |Time: 7:47 下午
 * +----------------------------------------------------------------------
 */

package mq

import (
	"github.com/gongxulei/go_kit/balancer"
	"github.com/gongxulei/go_kit/balancer/base"
	"github.com/gongxulei/go_kit/logger/origin"
	"strconv"
	"strings"
)

type Driver string

const (
	NSQ Driver = "nsq"
)

type ConfigOption func(c *Config)

type Config struct {
	Store        Driver
	AddrInfoList []*base.AddrInfo
	Hosts        []string
	Topic        string
	Channel      string
	Batch        int
	Logger       origin.LoggerInterface
	LoadStrategy balancer.Strategy
}

func Store(driver Driver) ConfigOption {
	return func(c *Config) { c.Store = driver }
}

// Host @param hosts = "127.0.0.1:4500|2,127.0.0.1:4600|1"
func Host(hosts string) ConfigOption {
	return func(c *Config) {
		var addrInfoList = make([]*base.AddrInfo, 0)
		addrList := strings.Split(hosts, ",")
		for _, addr := range addrList {
			addrInfo := new(base.AddrInfo)
			aw := strings.Split(addr, "|")
			addrInfo.Addr = aw[0]
			c.Hosts = append(c.Hosts, aw[0])
			if len(aw) >= 2 {
				weight, _ := strconv.Atoi(aw[1])
				if weight <= 0 {
					weight = 1
				}
				addrInfo.Weight = uint8(weight)
			}
			if len(aw) == 1 {
				addrInfo.Weight = 1
			}
			addrInfoList = append(addrInfoList, addrInfo)
		}
		c.AddrInfoList = addrInfoList
	}
}

func Logger(logger origin.LoggerInterface) ConfigOption {
	return func(c *Config) {
		c.Logger = logger
	}
}

func Topic(topic string) ConfigOption {
	return func(c *Config) {
		c.Topic = topic
	}
}

func Channel(channel string) ConfigOption {
	return func(c *Config) {
		c.Channel = channel
	}
}

func Batch(count int) ConfigOption {
	return func(c *Config) {
		c.Batch = count
	}
}

func LoadStrategy(loadStrategy balancer.Strategy) ConfigOption {
	return func(c *Config) {
		c.LoadStrategy = loadStrategy
	}
}

func DefaultConfig() *Config {
	return &Config{
		Store: NSQ,
		AddrInfoList: []*base.AddrInfo{
			{
				Addr:   "127.0.0.1:4500",
				Weight: 1,
			},
		},
		LoadStrategy: balancer.RANDOM, // 默认使用随机负载均衡算法
	}
}
