package msg

import (
	"github.com/golang/protobuf/proto"
)

func (msg *Message) Marshal() (bytes []byte, err error) {
	return proto.Marshal(msg)
}

func (msg *Message) Unmarshal(bytes []byte) (err error) {
	return proto.Unmarshal(bytes, msg)
}
