package listen

import (
	"github.com/gongxulei/go_kit/mq/msg"
)

var ListenMqMap = make(map[string]Listener, 10)

type Listener interface {
	OnMessage(messageBody []byte, messageHeader *msg.MessageHeader) error
}
