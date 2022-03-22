/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/23
 * +----------------------------------------------------------------------
 * |Time: 11:17 上午
 * +----------------------------------------------------------------------
 */

package pkg

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/appengine/log"
	"time"
)

type RedisConfig struct {
	NetWork        string
	Address        string
	UserName       string
	Password       string
	ConnectTimeout time.Duration
	Database       int
}

var Config = RedisConfig{
	NetWork:        "tcp",
	Address:        "127.0.0.1:6390",
	UserName:       "", // redis>6.0可填写用户名
	Password:       "123456",
	ConnectTimeout: time.Second * 3,
	Database:       0,
}

var Redis *redis.Client

func InitRedis() {
	var err error
	Redis = redis.NewClient(&redis.Options{
		Addr:       Config.Address,
		Password:   Config.Password,
		DB:         Config.Database,
		MaxRetries: 3,
		PoolSize:   10,
	})

	if err = Redis.Ping(context.Background()).Err(); err != nil {
		log.Errorf(context.TODO(), "未成功链接上redis,err:%s", err.Error())
	}
}
