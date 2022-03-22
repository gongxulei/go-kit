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

package nsq

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/gongxulei/go_kit/balancer"
	"github.com/gongxulei/go_kit/balancer/base"
	"github.com/gongxulei/go_kit/logger/origin"
	"github.com/gongxulei/go_kit/mq/msg"
	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
	"time"
)

type Producer struct {
	NsqProducerPool []*nsq.Producer
	AddrInfoList    []*base.AddrInfo
	Logger          origin.LoggerInterface
	current         uint64
	LoadStrategy    balancer.Strategy
	MsgBodyProtocol msg.Protocol
}

func NewProducer(addrInfoList []*base.AddrInfo, logger origin.LoggerInterface, LoadStrategy balancer.Strategy) (producer *Producer) {
	var nsqProducerPool = make([]*nsq.Producer, 0)
	producer = new(Producer)
	producer.AddrInfoList = addrInfoList
	producer.Logger = logger
	producer.LoadStrategy = LoadStrategy
	for _, addrInfo := range addrInfoList {
		nsqProducer, err := nsq.NewProducer(addrInfo.Addr, nsq.NewConfig())
		if err != nil {
			panic("connect nsq new producer error" + err.Error())
		}
		nsqProducerPool = append(nsqProducerPool, nsqProducer)
	}
	producer.NsqProducerPool = nsqProducerPool
	return producer
}

func (p *Producer) SendPb(topic, tag string, message proto.Message, delay int64) (err error) {
	var bytes []byte
	header := &msg.MessageHeader{
		Tag:          tag,
		UniqueId:     uuid.New().String(),
		BodyProtocol: msg.Protocol_PROTOCOL_PROTOBUF,
	}
	// message serialize
	bytes, err = proto.Marshal(message)
	if err != nil {
		return errors.New("proto.Message marshal  error " + err.Error())
	}
	return p.send(topic, string(bytes), header, delay)
}

func (p *Producer) SendString(topic, tag, message string, delay int64) (err error) {
	header := &msg.MessageHeader{
		Tag:          tag,
		UniqueId:     uuid.New().String(),
		BodyProtocol: msg.Protocol_PROTOCOL_STRING,
	}
	return p.send(topic, message, header, delay)
}

func (p *Producer) SendJson(topic, tag string, jsonBytes []byte, delay int64) (err error) {
	header := &msg.MessageHeader{
		Tag:          tag,
		UniqueId:     uuid.New().String(),
		BodyProtocol: msg.Protocol_PROTOCOL_JSON,
	}
	return p.send(topic, string(jsonBytes), header, delay)
}

func (p *Producer) send(topic, serializeBody string, mqMessageHeader *msg.MessageHeader, delay int64) (err error) {
	var msgData = new(msg.Message)
	var bytes []byte
	msgData = &msg.Message{
		Header: mqMessageHeader,
		Body:   serializeBody,
	}
	// message marshal
	bytes, err = msgData.Marshal()
	if err != nil {
		return errors.New("message marshal error " + err.Error())
	}
	// todo: 此处发送nsq消息需要添加重试机制
	if delay <= 0 {
		err = p.selectProducer().Publish(topic, bytes)
	} else {
		err = p.selectProducer().DeferredPublish(topic, time.Duration(delay)*time.Second, bytes)
	}
	return
}

// selectProducer by load balancing algorithm
func (p *Producer) selectProducer() (producer *nsq.Producer) {
	var ba base.Balancer
	// 默认使用【随机】算法
	switch p.LoadStrategy {
	case balancer.RANDOM:
		ba = balancer.NewBalancer(p.AddrInfoList, balancer.RANDOM)
	case balancer.ROUND:
		ba = balancer.NewBalancer(p.AddrInfoList, balancer.ROUND)
	default:
		ba = balancer.NewBalancer(p.AddrInfoList, balancer.RANDOM)
	}
	addr := ba.Balance()
	return p.NsqProducerPool[p.getProducerIndex(addr)]
}

func (p *Producer) getProducerIndex(addr string) (index int) {
	for i, producer := range p.NsqProducerPool {
		if producer.String() == addr {
			return i
		}
	}
	return
}

func (p *Producer) Close() {
	for _, producer := range p.NsqProducerPool {
		p.Logger.Log("stop producer : %s \n", producer.String())
		producer.Stop()
	}
}
