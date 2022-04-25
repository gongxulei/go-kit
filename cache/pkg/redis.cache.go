package pkg

import (
	"context"
	"fmt"
	"time"
)

/*redis缓存*/

const (
	// front_category:%s:*
	FrontCategoryPrefixPattern = "%s*"
	ScanKeyCount               = 1000
	ExpiredTime                = time.Minute * 30
	DistributedLockTime        = time.Second * 15
	DistributedLockTryTime     = time.Second * 5
)

/*
Set 设置缓存
@param ctx context.Context
@param key string
@param content []byte
*/
func Set(ctx context.Context, key string, content []byte) error {
	return Redis.Set(ctx, key, string(content), ExpiredTime).Err()
}

/*
DeleteCacheBatch 根据key前缀批量删除缓存
@param ctx context.Context
@param key string 要匹配的key
*/
func DeleteCacheBatch(ctx context.Context, key string) (err error) {
	cacheKeyPrefix := fmt.Sprintf(FrontCategoryPrefixPattern, key)
	redis := Redis
	var (
		cursor uint64
		keys   []string
	)
	for {
		keys, cursor, err = redis.Scan(ctx, cursor, cacheKeyPrefix, ScanKeyCount).Result()
		if err != nil {
			return
		}
		if len(keys) > 0 {
			// 删除key
			_, err = redis.Del(ctx, keys...).Result()
			if err != nil {
				return
			}
		}
		// 遍历完成
		if cursor == 0 {
			return
		}
	}
}

/*
TryGetDistributedLock 添加分布式锁
@param ctx context.Context
@param lockKey string 加锁的key
@param requestId string 加锁key存储的唯一标识
*/
func TryGetDistributedLock(ctx context.Context, lockKey string, requestId string) bool {
	// 获取时间
	start := time.Now().UnixNano()
	ok, err := Redis.SetNX(ctx, lockKey, requestId, DistributedLockTime).Result()
	if err != nil {
		return false
	}
	end := time.Now().UnixNano()
	if time.Duration(end-start) > DistributedLockTryTime {
		// 加锁返回如果超过指定时间就认为加锁失败
		return false
	}
	return ok
}

/*
ReleaseDistributedLock 释放分布式锁
@param ctx context.Context
@param lockKey string 被加锁的key
@param requestId string 加锁key存储的唯一标识
*/
func ReleaseDistributedLock(ctx context.Context, lockKey string, requestId string) bool {
	script := `if redis.call("GET", KEYS[1]) == ARGV[1] then
                return redis.call("DEL", KEYS[1])
            else
                return 0
            end`
	ok, err := Redis.Eval(ctx, script, []string{lockKey}, requestId).Int()
	if err != nil {
		return false
	}
	return ok != 0
}
