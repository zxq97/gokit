package rpc

import (
	"context"
	"time"

	"github.com/zxq97/gokit/pkg/constant"
	"github.com/zxq97/gokit/pkg/trace"
	"google.golang.org/grpc/metadata"
)

func getInComing(ctx context.Context) (traceID string, parentID string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		rs := md.Get(constant.TraceIDKey)
		if len(rs) > 0 {
			traceID = rs[0]
		}
		rs = md.Get(constant.SpanIDKey)
		if len(rs) > 0 {
			parentID = rs[0]
		}
	}
	return traceID, parentID
}

func withOutgoing(ctx context.Context) context.Context {
	traceID := trace.GetTraceID(ctx)
	spanID, ok := ctx.Value(constant.SpanIDKey).(string)
	if !ok {
		spanID = constant.DefaultSpanID
	}
	return metadata.AppendToOutgoingContext(ctx, constant.TraceIDKey, traceID, constant.SpanIDKey, spanID)
}

func defaultTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, d)
	}
	return ctx, cancel
}
