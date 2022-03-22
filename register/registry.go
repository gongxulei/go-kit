/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/11/29
 * +----------------------------------------------------------------------
 * |Time: 10:45 下午
 * +----------------------------------------------------------------------
 */

package register

import (
	"context"
	"github.com/gongxulei/go_kit/register/plugins"
)

// Registry 服务注册插件的接口
type Registry interface {
	// PluginName 插件的名称
	PluginName() string
	// NewClient 初始化客户端
	NewClient(ctx context.Context, opts ...plugins.OptionFun) (err error)
	// Register 服务注册
	Register(ctx context.Context, service *Service) (err error)
	// Deregister 服务注销
	Deregister(ctx context.Context, serviceId string) (err error)
}
