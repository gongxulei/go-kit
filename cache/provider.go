/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/23
 * +----------------------------------------------------------------------
 * |Time: 9:22 下午
 * +----------------------------------------------------------------------
 */

package cache

import (
	"context"
	"time"
)

type Provider interface {
	Set(ctx context.Context, key, value string, expired time.Duration) (err error)
	Get(ctx context.Context, key string) (res string, err error)
	Handler() (client interface{})
}
