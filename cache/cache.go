/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/23
 * +----------------------------------------------------------------------
 * |Time: 9:15 下午
 * +----------------------------------------------------------------------
 */

package cache

import (
	"github.com/gongxulei/go_kit/cache/config"
	"github.com/gongxulei/go_kit/cache/driver"
)

func NewCache(opts ...config.ConfigOption) (cache Provider, cleanup func(), err error) {
	cfg := config.DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.Store == config.Redis {
		cache, cleanup, err = driver.NewRedis(cfg)
	}
	if cfg.Store == config.GoCache {

	}
	return
}
