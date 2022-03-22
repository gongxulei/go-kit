/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/12
 * +----------------------------------------------------------------------
 * |Time: 5:50 下午
 * +----------------------------------------------------------------------
 */

package test

import (
	"fmt"
	"github.com/gongxulei/go_kit/mq/driver/nsq"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	c := &nsq.ConsumerConfig{
		LookupAddr:  []string{"localhost:4161"},
		Topic:       "topic_test1",
		Channel:     "channel11111",
		MaxInFlight: 10,
	}
	consumer, _, _ := nsq.NewConsumer(c, nil)

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println("signal: ", sig)
	// Gracefully stop the consumer.
	consumer.Close()
}
