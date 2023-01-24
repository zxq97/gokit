package rpc

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"github.com/zxq97/gokit/internal/grpcerr"
	"github.com/zxq97/gokit/pkg/constant"
	"github.com/zxq97/gokit/pkg/errcode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(info.FullMethod, err, string(debug.Stack()))
		}
	}()
	return handler(ctx, req)
}

func access(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()
	defer func() {
		log.Println(info.FullMethod, time.Since(now))
	}()
	return handler(ctx, req)
}

func bizErr(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if e, ok := err.(*errcode.BizErr); ok {
		s, err := status.New(codes.Internal, e.Error()).WithDetails(&grpcerr.GrpcErr{Code: e.Code, Message: e.Message, Detail: e.Detail})
		if err != nil {
			return res, err
		}
		return res, s.Err()
	}
	return res, err
}

func inComingTrace(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	traceID, parentID := getInComing(ctx)
	if traceID != "" {
		ctx = context.WithValue(ctx, constant.TraceIDKey, traceID)
	}
	if parentID != "" {
		ctx = context.WithValue(ctx, constant.ParentIDKey, parentID)
	}
	return handler(ctx, req)
}

func outgoingTrace() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = withOutgoing(ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func timeout(d time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var cancel context.CancelFunc
		ctx, cancel = defaultTimeout(ctx, d)
		if cancel != nil {
			defer cancel()
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func unBizErr() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		s := status.Convert(err)
		for _, v := range s.Details() {
			if e, ok := v.(*grpcerr.GrpcErr); ok {
				return errcode.NewBizErr(e.Code, e.Message, e.Detail)
			}
		}
		return err
	}
}
