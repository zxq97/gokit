package trace

import (
	"context"

	"github.com/zxq97/gokit/pkg/constant"
	"github.com/zxq97/gokit/pkg/generate"
)

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(constant.TraceIDKey).(string)
	if !ok {
		traceID = generate.UUIDStr()
	}
	return traceID
}
