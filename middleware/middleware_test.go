/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/4
 * +----------------------------------------------------------------------
 * |Time: 10:37 上午
 * +----------------------------------------------------------------------
 */

package middleware

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
)

func TestChain(t *testing.T) {
	middleware1 := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Printf("middleware 1 start\n")
			num := rand.Intn(2)
			if num == 2 {
				err = fmt.Errorf("this is request is not allow")
				return
			}
			resp, err = next(ctx, req)
			if err != nil {
				return
			}
			fmt.Printf("middleware1 end\n")
			return
		}
	}

	middleware2 := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Printf("middleware 2 start\n")

			resp, err = next(ctx, req)
			if err != nil {
				return
			}
			fmt.Printf("middleware2 end\n")
			return
		}
	}

	outer := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Printf("outer  start\n")
			resp, err = next(ctx, req)
			if err != nil {
				return
			}
			fmt.Printf("outer end\n")
			return
		}
	}

	proc := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		fmt.Printf("req process start\n")
		fmt.Printf("req process end\n")
		return
	}

	chain := Chain(outer, middleware1, middleware2)

	// chain := func(next HandlerFunc) HandlerFunc {
	// 	next = middleware2(next)
	// 	next = middleware1(next)
	// 	return outer(next)
	// }

	// outer.next(middleware1) --> middleware1.next((middleware2)) --> middleware2.next((proc))-->

	chainFunc := chain(proc)

	resp, err := chainFunc(context.Background(), "test")
	fmt.Printf("resp:%#v, err:%v\n", resp, err)
}
