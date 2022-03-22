/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/12
 * +----------------------------------------------------------------------
 * |Time: 8:19 下午
 * +----------------------------------------------------------------------
 */

package mq

import (
	"github.com/golang/protobuf/proto"
)

type Producer interface {
	SendPb(topic, tag string, message proto.Message, delay int64) (err error)
	SendString(topic, tag, message string, delay int64) (err error)
	SendJson(topic, tag string, jsonBytes []byte, delay int64) (err error)
	Close()
}

type Consumer interface {
	Close()
}
