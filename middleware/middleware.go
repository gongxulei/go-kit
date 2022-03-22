/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/4
 * +----------------------------------------------------------------------
 * |Time: 10:28 上午
 * +----------------------------------------------------------------------
 */

package middleware

import (
	"context"
)

type HandlerFunc func(ctx context.Context, req interface{}) (resp interface{}, err error)

type Middleware func(HandlerFunc) HandlerFunc

func Chain(outer Middleware, middlewares ...Middleware) Middleware {
	return func(next HandlerFunc) HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return outer(next)
	}
}
