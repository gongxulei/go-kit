package mq

import (
	selfNsq "github.com/gongxulei/go_kit/mq/driver/nsq"
	"github.com/gongxulei/go_kit/mq/listen"
)

func NewProducer(opts ...ConfigOption) (provider Producer) {
	config := DefaultConfig()
	for _, opt := range opts {
		opt(config)
	}
	if config.Store == NSQ {
		return selfNsq.NewProducer(config.AddrInfoList, config.Logger, config.LoadStrategy)
	}
	return nil
}

func NewConsumer(opts ...ConfigOption) (consumer Consumer, lookup func(), err error) {
	config := DefaultConfig()
	for _, opt := range opts {
		opt(config)
	}
	if config.Store == NSQ {
		inFlight := config.Batch
		if config.Batch <= 0 {
			inFlight = 10
		}
		c := &selfNsq.ConsumerConfig{
			LookupAddr:  config.Hosts,
			Topic:       config.Topic,
			Channel:     config.Channel,
			MaxInFlight: inFlight,
			Driver:      string(NSQ),
		}
		consumer, lookup, err = selfNsq.NewConsumer(c, config.Logger)
	}
	return
}

func Register(driver Driver, lis listen.Listener) {
	listen.ListenMqMap[string(driver)] = lis
}
