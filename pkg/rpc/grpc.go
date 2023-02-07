package rpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zxq97/gokit/pkg/etcd"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	defaultMaxRecvSize = 1024 * 1024 * 16
	defaultMetricsURI  = "/metrics"
)

func NewGrpcConn(ctx context.Context, svcName string, client *clientv3.Client) (*grpc.ClientConn, error) {
	w, err := etcd.NewWatcher(ctx, GetWatcherKey(svcName), client)
	if err != nil {
		return nil, err
	}
	rs := newBuilder(w)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithResolvers(rs),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				timeout(time.Second),
				outgoingTrace(),
				unBizErr(),
				grpc_prometheus.UnaryClientInterceptor,
			)),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	}
	return grpc.DialContext(ctx, rs.Scheme()+":///", opts...)
}

func NewGrpcServer(ctx context.Context, conf *Config, client *clientv3.Client) (*grpc.Server, error) {
	r := newRegister(client)
	err := r.Register(ctx, GetSvc(conf))
	if err != nil {
		return nil, err
	}
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(defaultMaxRecvSize),
		grpc_middleware.WithUnaryServerChain(
			recovery,
			access,
			bizErr,
			inComingTrace,
			grpc_prometheus.UnaryServerInterceptor,
		),
	}
	svc := grpc.NewServer(opts...)
	reflection.Register(svc)
	grpc_health_v1.RegisterHealthServer(svc, health.NewServer())
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(svc)
	http.Handle(defaultMetricsURI, promhttp.Handler())
	return svc, nil
}
