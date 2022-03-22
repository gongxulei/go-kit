/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/23
 * +----------------------------------------------------------------------
 * |Time: 9:28 下午
 * +----------------------------------------------------------------------
 */

package driver

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gongxulei/go_kit/cache/config"
	// "github.com/google/wire"
	"time"
)

type RedisCache struct {
	// default expired
	expired time.Duration
	handler *redis.Client
}

var RedisHandler *redis.Client

func (cache *RedisCache) Set(ctx context.Context, key, value string, expired time.Duration) (err error) {
	if expired == 0 {
		expired = cache.expired
	}
	return cache.handler.Set(ctx, key, value, expired).Err()
}

func (cache *RedisCache) Get(ctx context.Context, key string) (res string, err error) {
	return cache.handler.Get(ctx, key).Result()
}

func (cache *RedisCache) Handler() (client interface{}) {
	return cache.handler
}

func NewRedis(cfg *config.Config) (redisCache *RedisCache, cleanup func(), err error) {
	redisCache = new(RedisCache)
	RedisHandler = redis.NewClient(&redis.Options{
		Addr:         cfg.Socket,
		Password:     cfg.Password,
		DB:           cfg.DB,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConnes,
	})

	if err = RedisHandler.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("connect redis error：%s", err.Error())
		return
	}
	redisCache.expired = cfg.Expired
	redisCache.handler = RedisHandler
	cleanup = func() {
		_ = RedisHandler.Close()
	}
	return
}
