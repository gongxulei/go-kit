/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/1
 * +----------------------------------------------------------------------
 * |Time: 10:51 下午
 * +----------------------------------------------------------------------
 */

package register

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var RegistryManager = PluginManager{
	Plugins: make(map[string]Registry),
	Lock:    sync.Mutex{},
}

var pluginNameList = make([]string, 0)

// PluginManager 插件管理器
type PluginManager struct {
	Plugins map[string]Registry
	Lock    sync.Mutex
}

// BindPlugin 注册插件到插件管理器中
func (pm *PluginManager) BindPlugin(registry Registry) (err error) {
	pm.Lock.Lock()
	defer pm.Lock.Unlock()

	_, ok := pm.Plugins[registry.PluginName()]
	if ok {
		err = errors.New("duplicate register plugin")
		return
	}
	pm.Plugins[registry.PluginName()] = registry
	return
}

// GetPlugin 获取插件
func (pm *PluginManager) GetPlugin(ctx context.Context, name string) (registry Registry, err error) {
	var ok bool
	pm.Lock.Lock()
	defer pm.Lock.Unlock()
	// 判断插件是否存在
	registry, ok = pm.Plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s is not exists", name)
		return
	}
	return
}

// GetSupportPluginName 获取支持的插件名称列表
func GetSupportPluginName() (pluginList []string) {
	return pluginNameList
}

// PushPluginName 记录插件名称
func PushPluginName(name string) (pluginList []string) {
	pluginNameList = append(pluginNameList, name)
	return
}
