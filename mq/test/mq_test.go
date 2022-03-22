/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/13
 * +----------------------------------------------------------------------
 * |Time: 4:20 下午
 * +----------------------------------------------------------------------
 */

package test

import (
	"encoding/json"
	"fmt"
	"github.com/gongxulei/go_kit/balancer"
	"github.com/gongxulei/go_kit/logger/driver"
	"github.com/gongxulei/go_kit/mq"
	"github.com/gongxulei/go_kit/mq/msg"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

type listenTest struct{}

func (l *listenTest) OnMessage(messageBody []byte, messageHeader *msg.MessageHeader) error {
	bytes, _ := json.Marshal(messageHeader)
	fmt.Printf("body: %s;;;;;header:%s", string(messageBody), string(bytes))
	return nil
}

func TestNewProvider(t *testing.T) {
	var configSplit = make([]mq.ConfigOption, 0)
	log := driver.NewFileLogger(0, "/Volumes/work/bb_golang/src/gong/logger", "nsq")
	topic := "topic_test2"

	configSplit = append(configSplit, mq.Store(mq.NSQ))
	configSplit = append(configSplit, mq.Host("127.0.0.1:4150|1"))
	configSplit = append(configSplit, mq.Topic(topic))
	configSplit = append(configSplit, mq.Batch(5))
	configSplit = append(configSplit, mq.Logger(log))
	configSplit = append(configSplit, mq.LoadStrategy(balancer.RANDOM))
	provider := mq.NewProducer(configSplit...)

	// send message
	err := provider.SendString(topic, "test", "this is a test", 0)
	err = provider.SendPb(topic, "test_pb", &GetProductDetailRequest{
		ProductId: "1012230988",
		NameSpace: "test",
		OneField:  nil,
		MapInfo:   map[string]int64{"price": 120},
	}, 0)
	if err != nil {
		t.Logf("send message string error: %#v", err)
	}

	provider.Close()
}

func Test_NewCustomer(t *testing.T) {
	// register listen
	listen := new(listenTest)
	mq.Register(mq.NSQ, listen)

	var configSplit = make([]mq.ConfigOption, 0)
	log := driver.NewFileLogger(0, "/Volumes/work/bb_golang/src/gong/logger", "nsq")
	topic := "topic_test2"
	channel := "channel11111"

	configSplit = append(configSplit, mq.Store(mq.NSQ))
	configSplit = append(configSplit, mq.Host("127.0.0.1:4161"))
	configSplit = append(configSplit, mq.Topic(topic))
	configSplit = append(configSplit, mq.Channel(channel))
	configSplit = append(configSplit, mq.Batch(5))
	configSplit = append(configSplit, mq.Logger(log))

	_, lookup, err := mq.NewConsumer(configSplit...)
	if err != nil {
		t.Errorf("NewConsumer error: %#v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println("signal: ", sig)
	lookup()
}
