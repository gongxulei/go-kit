/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/12
 * +----------------------------------------------------------------------
 * |Time: 4:01 下午
 * +----------------------------------------------------------------------
 */

package test

import (
	"github.com/golang/protobuf/proto"
	"github.com/gongxulei/go_kit/mq/msg"
	"github.com/google/uuid"
	"testing"
)

func TestMqMessage_Marshal_Unmarshal(t *testing.T) {
	header := &msg.MessageHeader{
		Tag:      "order_canal",
		UniqueId: uuid.New().String(),
	}
	body := &GetProductDetailRequest{
		ProductId: "123123123",
		NameSpace: "aaaaa",
	}
	bodyProto, _ := proto.Marshal(body)
	message := &msg.Message{
		Header: header,
		Body:   string(bodyProto),
	}

	// convert json
	// jsonBytes,_ := json.Marshal(msg)
	// t.Logf("marshal success json res :%s, ;;;;;;;length:%d", string(jsonBytes), len(string(jsonBytes)))

	bytes, err := message.Marshal()
	if err != nil {
		t.Errorf("marshal error %#v", err)
		return
	}
	// t.Logf("marshal success res :%s, ;;;;;;;length:%d", string(bytes), len(string(bytes)))
	// enc_str := base64.StdEncoding.EncodeToString(bytes)
	// t.Logf("marshal success base64 res :%s;;;;;;;length:%d", enc_str, len(enc_str))
	//
	// bytes2, _ := base64.StdEncoding.DecodeString(enc_str)

	var res = new(msg.Message)
	err = res.Unmarshal(bytes)
	if err != nil {
		t.Errorf("Unmarshal error %#v", err)
		return
	}

	t.Logf("Unmarshal res :%#v", res)

	var body2 = new(GetProductDetailRequest)
	// ptypes.UnmarshalAny(res.Body, body2)
	proto.Unmarshal([]byte(res.Body), body2)
	t.Logf("test res :%#v", body2)

}
