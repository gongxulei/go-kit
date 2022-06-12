/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: GongXuLei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/7
 * +----------------------------------------------------------------------
 * |Time: 11:07 下午
 * +----------------------------------------------------------------------
 */

package plugins

import (
	"context"
	"fmt"
	"github.com/gongxulei/go_kit/register"
	consulApi "github.com/hashicorp/consul/api"
	"sync"
	"time"
)

const (
	Name = "consul"
)

type ConsulRegistry struct {
	options   *register.Options
	client    *consulApi.Client
	serviceCh chan *register.Service
	lockMutex sync.Mutex
}

func init() {
	// 自动注册
	register.PushPluginName(Name)
}

func (consul *ConsulRegistry) PluginName() string {
	return Name
}

func Default() *ConsulRegistry {
	return &ConsulRegistry{}
}

func (consul *ConsulRegistry) NewClient(_ context.Context, opts ...register.OptionFun) (err error) {
	consul.options = &register.Options{}
	for _, opt := range opts {
		opt(consul.options)
	}
	// 连接consul
	consul.client, err = consulApi.NewClient(&consulApi.Config{
		Address: consul.options.Address[0],
	})
	consul.serviceCh = make(chan *register.Service, 1)
	return
}

func (consul *ConsulRegistry) Register(_ context.Context, service *register.Service) (err error) {
	//select {
	//case consul.serviceCh <- service:
	//default:
	//	err = errors.New("register chan is full")
	//	return
	//}
	//return
	var addresses = make(map[string]consulApi.ServiceAddress)
	var registration *consulApi.AgentServiceRegistration
	for _, node := range service.Nodes {
		registration = &consulApi.AgentServiceRegistration{
			Kind:              "Golang",
			ID:                node.Id,
			Name:              service.Name,
			Tags:              node.Tags,
			Port:              node.Port,
			Address:           node.Address,
			TaggedAddresses:   addresses,
			EnableTagOverride: true,
			Meta:              node.MetaData,
		}
		// 设置权重
		registration.Weights = &consulApi.AgentWeights{
			Passing: node.Weight,
			Warning: 0,
		}
		// 设置TaggedAddresses
		registration.TaggedAddresses[node.Scheme] = consulApi.ServiceAddress{
			Address: node.Address,
			Port:    node.Port,
		}

		// 设置健康检查socket
		registration.Check = &consulApi.AgentServiceCheck{
			Name:                           service.Name,
			Interval:                       "5s",
			Timeout:                        "10s",
			DockerContainerID:              "",
			DeregisterCriticalServiceAfter: "30s",
		}

		if node.Scheme == "http" {
			registration.Check.HTTP = fmt.Sprintf("http://%s:%d", node.Address, node.Port)
		}
		if node.Scheme == "tcp" {
			registration.Check.TCP = fmt.Sprintf("%s:%d", node.Address, node.Port)
		}
		if node.Scheme == "grpc" {
			registration.Check.GRPC = fmt.Sprintf("%s:%d", node.Address, node.Port)
		}

		err = consul.client.Agent().ServiceRegister(registration)
		if err != nil {
			return err
		}
	}
	return
}

func (consul *ConsulRegistry) Deregister(_ context.Context, serviceId string) (err error) {
	return consul.client.Agent().ServiceDeregister(serviceId)
}

// GetService 查询服务
func (consul *ConsulRegistry) GetService(ctx context.Context, serviceName, tag string, passingOnly bool, index uint64) (services *register.Service, lastIndex uint64, err error) {
	var entries []*consulApi.ServiceEntry
	var meta *consulApi.QueryMeta
	services = new(register.Service)

	// todo: 查询缓存中是否有服务信息

	consul.lockMutex.Lock()
	defer consul.lockMutex.Unlock()
	queryOption := &consulApi.QueryOptions{
		WaitIndex: index,
		WaitTime:  time.Second * 55,
	}
	queryOption = queryOption.WithContext(ctx)
	entries, meta, err = consul.client.Health().Service(serviceName, tag, passingOnly, queryOption)
	if err != nil {
		return
	}
	services = defaultResolver(ctx, entries, serviceName)

	// todo: 此处可以对service添加缓存

	return defaultResolver(ctx, entries, serviceName), meta.LastIndex, nil
}

// ServiceResolver 解析默认服务信息
type ServiceResolver func(ctx context.Context, entries []*consulApi.ServiceEntry, serviceName string) (services *register.Service)

func defaultResolver(_ context.Context, entries []*consulApi.ServiceEntry, serviceName string) (services *register.Service) {
	services = new(register.Service)
	for _, entry := range entries {
		var Nodes = make([]*register.Node, 0)
		if serviceName != entry.Service.Service {
			continue
		}
		for scheme, serviceAddress := range entry.Service.TaggedAddresses {
			if scheme == "lan_ipv4" || scheme == "wan_ipv4" || scheme == "lan_ipv6" || scheme == "wan_ipv6" {
				continue
			}
			Nodes = append(Nodes, &register.Node{
				Id:       entry.Service.ID,
				Tags:     entry.Service.Tags,
				MetaData: entry.Service.Meta,
				Scheme:   scheme,
				Weight:   entry.Service.Weights.Passing,
				Address:  serviceAddress.Address,
				Port:     serviceAddress.Port,
			})
		}
		services = &register.Service{
			Name:  entry.Service.Service,
			Nodes: Nodes,
		}
		return
	}
	return
}
