/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/23
 * +----------------------------------------------------------------------
 * |Time: 10:42 下午
 * +----------------------------------------------------------------------
 */

package cache

import (
	"context"
	"github.com/gongxulei/go_kit/cache/config"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cfgOps := make([]config.ConfigOption, 0)
	cfgOps = append(cfgOps, config.Store("redis"))
	cfgOps = append(cfgOps, config.Socket("127.0.0.1:6391"))
	cfgOps = append(cfgOps, config.Password("123456"))

	cache, _, err := NewCache(cfgOps...)
	if err != nil {
		t.Errorf("NewCache error:%#v", err)
		return
	}
	err = cache.Set(context.TODO(), "name", "zhanga", time.Second)
	if err != nil {
		t.Errorf("set cache error:%#v", err)
	}
	cache.Handler()

}
