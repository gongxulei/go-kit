package test

import (
	"github.com/gongxulei/go_kit/mq/msg"
	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func TestNsqProducer(t *testing.T) {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	header := &msg.MessageHeader{
		Tag:      "order_canal",
		UniqueId: uuid.New().String(),
	}
	body := &GetProductDetailRequest{
		ProductId: "123123123",
		NameSpace: "aaaaa",
	}
	// anyBody, _ := anypb.New(body)
	protoBody, _ := proto.Marshal(body)
	msg := &msg.Message{
		Header: header,
		Body:   string(protoBody),
	}
	bytes, err := msg.Marshal()
	if err != nil {
		t.Errorf("marshal error %#v", err)
		return
	}

	topicName := "topic_test1"

	// Synchronously publish a single msg to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, bytes)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()
}
