/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/14
 * +----------------------------------------------------------------------
 * |Time: 2:52 下午
 * +----------------------------------------------------------------------
 */

package plugins

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gongxulei/go_kit/encrypt"
	myNet "github.com/gongxulei/go_kit/net"
	"github.com/gongxulei/go_kit/register"
	"net"
	"net/http"
	"testing"
	"time"
)

func tcpServer(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		fmt.Println("get tcp")
		conn.Close()
	}
}

func httpServer(localAddress string, HttpPort int, t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get http")
		fmt.Fprintf(w, "Hello golang http!")
	})
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", localAddress, HttpPort), nil)
	if err != nil {
		t.Errorf("listen http %s failed!, err: %#v", localAddress, err)
		t.Fail()
	}
}

func TestPluginManager_InitRegistry(t *testing.T) {
	const Port = 8083
	const HttpPort = 8082
	// 获取本机地址
	var localAddress = myNet.GetIntranetIP()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", localAddress, Port))
	if err != nil {
		t.Errorf("listen tcp %s failed!", localAddress)
		t.Fail()
	}
	defer func() { _ = listen.Close() }()
	go tcpServer(listen)

	go httpServer(localAddress, HttpPort, t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// 实例化consul client
	client := Default()
	err = client.NewClient(ctx, register.WithOptionAddress([]string{"127.0.0.1:8500"}), register.WithOptionTimeout(time.Second*5))
	if err != nil {
		t.Errorf("NewClient failed! %#v", err)
		t.Fail()
	}
	// 注册到插件管理器中
	//err = register.RegistryManager.BindPlugin(client)
	//if err != nil {
	//	t.Errorf("RegisterPlugin failed! %#v", err)
	//	t.Fail()
	//}

	// 注册服务信息到consul
	service := &register.Service{
		Name: "app-id-test2",
		Nodes: []*register.Node{
			{
				Id:       encrypt.Md5V1(fmt.Sprintf("%s%d", localAddress, Port)),
				Scheme:   "tcp",
				Address:  localAddress,
				Port:     Port,
				Weight:   10,
				Tags:     []string{"tcp"},
				MetaData: nil,
			},
			{
				Id:       encrypt.Md5V1(fmt.Sprintf("%s%d", localAddress, HttpPort)),
				Scheme:   "http",
				Address:  localAddress,
				Port:     HttpPort,
				Weight:   10,
				Tags:     []string{"http"},
				MetaData: nil,
			},
		},
	}
	err = client.Register(ctx, service)
	if err != nil {
		t.Errorf("Register failed! %#v", err)
		t.Fail()
	}

	time.Sleep(6 * time.Second)

	// 注销服务
	//err = client.Deregister(ctx, "IDapp-id-test")
	//if err != nil {
	//	t.Errorf("Register failed! %#v", err)
	//	t.Fail()
	//}

	// 获取service
	serviceList, index, err := client.GetService(ctx, "app-id-test2", "tcp", true, 0)
	str, err := json.Marshal(serviceList)
	t.Logf("serviceList: %s;;;;;; index:%d", string(str), index)

	time.Sleep(10 * time.Minute)
}
