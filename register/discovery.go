/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/20
 * +----------------------------------------------------------------------
 * |Time: 5:44 下午
 * +----------------------------------------------------------------------
 */

package register

import (
	"context"
)

type Discovery interface {
	// GetService 根据serviceName直接拉取实例列表
	GetService(ctx context.Context, serviceName, tag string, passingOnly bool, index uint64) (services *Service, lastIndex uint64, err error)
	// Watch 根据serviceName阻塞式订阅一个服务的实例列表信息
	//Watch(ctx context.Context, serviceName string) (Watcher, error)
}
