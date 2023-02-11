package mq

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/generate"
	"github.com/zxq97/gokit/pkg/trace"
)

func WarpMessage(ctx context.Context, tag string, msg proto.Message) ([]byte, error) {
	traceID := trace.GetTraceID(ctx)
	bs, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	rocketMsg := &MqMessage{
		TxId:    generate.UUIDStr(),
		TraceId: traceID,
		Tag:     tag,
		Message: bs,
	}
	return proto.Marshal(rocketMsg)
}
