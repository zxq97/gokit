package kafka

import (
	"context"
	"time"

	"github.com/zxq97/gokit/pkg/constant"
)

func consumerContext(traceID string, d time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	ctx = context.WithValue(context.Background(), constant.TraceIDKey, traceID)
	if d.Seconds() != 0 {
		ctx, cancel = context.WithTimeout(ctx, d)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	return ctx, cancel
}
