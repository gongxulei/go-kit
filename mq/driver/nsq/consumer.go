/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/11
 * +----------------------------------------------------------------------
 * |Time: 4:43 下午
 * +----------------------------------------------------------------------
 */

package nsq

import (
	"github.com/gongxulei/go_kit/logger/origin"
	"github.com/gongxulei/go_kit/mq/listen"
	"github.com/gongxulei/go_kit/mq/msg"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

const (
	LookupdPollInterval = 3 * time.Second
	Concurrency         = 10
)

type ConsumerConfig struct {
	LookupAddr  []string
	Topic       string
	Channel     string
	MaxInFlight int
	Driver      string
}

type Consumer struct {
	config   *ConsumerConfig
	topic    string
	consumer *nsq.Consumer
	log      origin.LoggerInterface
	lock     sync.Mutex
}

// HandleMessage implements the Handler interface.
func (consumer *Consumer) HandleMessage(m *nsq.Message) (err error) {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the msg as processed.
		// In this case, a msg with an empty body is simply ignored/discarded.
		return
	}

	// do whatever actual msg processing is desired
	var mqMessage = new(msg.Message)
	err = mqMessage.Unmarshal(m.Body)
	if err != nil {
		consumer.log.LogError("Unmarshal error %#v", err)
		return
	}
	lis, ok := listen.ListenMqMap[consumer.config.Driver]
	if !ok {
		consumer.log.LogError("ListenMqMap not bind listener error")
		return
	}
	err = lis.OnMessage([]byte(mqMessage.Body), mqMessage.Header)

	return
}

// NewConsumer ...
func NewConsumer(c *ConsumerConfig, log origin.LoggerInterface) (consumerNsq *Consumer, lookup func(), err error) {
	consumerNsq = new(Consumer)
	var consumer *nsq.Consumer
	consumerNsq.topic = c.Topic
	consumerNsq.log = log
	consumerNsq.config = c
	config := nsq.NewConfig()
	config.LookupdPollInterval = LookupdPollInterval
	config.MaxInFlight = c.MaxInFlight
	consumer, err = nsq.NewConsumer(c.Topic, c.Channel, config)
	if err != nil {
		return
	}
	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	consumer.AddConcurrentHandlers(consumerNsq, Concurrency)
	// Use nsqlookupd to discover nsqd instances.  See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupds(c.LookupAddr)

	lookup = func() {
		consumer.Stop()
	}
	consumerNsq.consumer = consumer
	return

}

func (consumer *Consumer) Close() {
	consumer.consumer.Stop()
}
