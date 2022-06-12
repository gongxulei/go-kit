/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/23
 * +----------------------------------------------------------------------
 * |Time: 9:51 下午
 * +----------------------------------------------------------------------
 */

package config

import (
	"time"
)

type Driver string

const (
	Redis   Driver = "redis"
	GoCache Driver = "go-cache"
)

type ConfigOption func(c *Config)

type Config struct {
	Store   Driver
	Expired time.Duration
	// 192.168.0.21:6379
	Socket        string
	Password      string
	DB            int
	MaxRetries    int
	PoolSize      int
	MinIdleConnes int
}

func Store(driver Driver) ConfigOption {
	return func(c *Config) { c.Store = driver }
}

func Expired(expired time.Duration) ConfigOption {
	return func(c *Config) { c.Expired = expired }
}

func Socket(socket string) ConfigOption {
	return func(c *Config) { c.Socket = socket }
}

func Password(password string) ConfigOption {
	return func(c *Config) { c.Password = password }
}

func DB(db int) ConfigOption {
	return func(c *Config) { c.DB = db }
}

func MaxRetries(maxRetries int) ConfigOption {
	return func(c *Config) { c.MaxRetries = maxRetries }
}

func PoolSize(poolSize int) ConfigOption {
	return func(c *Config) { c.PoolSize = poolSize }
}
func MinIdleConnes(minIdleConnes int) ConfigOption {
	return func(c *Config) { c.MinIdleConnes = minIdleConnes }
}

func DefaultConfig() *Config {
	return &Config{
		Store:         GoCache,
		Expired:       time.Hour * 2,
		Socket:        "127.0.0.1:6379",
		Password:      "",
		DB:            0,
		MaxRetries:    5,
		PoolSize:      1,
		MinIdleConnes: 1,
	}
}
