/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/11/29
 * +----------------------------------------------------------------------
 * |Time: 10:39 下午
 * +----------------------------------------------------------------------
 */

package plugins

import (
	"time"
)

type Options struct {
	Address []string
	Timeout time.Duration
}

type OptionFun func(opts *Options)

func WithOptionAddress(address []string) OptionFun {
	return func(opts *Options) {
		opts.Address = address
	}
}

func WithOptionTimeout(timeout time.Duration) OptionFun {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}
