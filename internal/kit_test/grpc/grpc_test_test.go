package grpc

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
	"github.com/zxq97/gokit/pkg/etcd"
	"github.com/zxq97/gokit/pkg/rpc"
	"go.etcd.io/etcd/client/v3"
)

func getEtcdClient() (*clientv3.Client, error) {
	conf := &etcd.Config{}
	err := config.LoadYaml("../../yaml/etcd.yaml", conf)
	if err != nil {
		return nil, err
	}
	log.Println(conf.Addr, conf.TTL)
	return etcd.NewEtcd(conf)
}

type grpcTest struct {
}

func (grpcTest) Hello(ctx context.Context, res *HelloRequest) (*HelloResponse, error) {
	log.Println(ctx, res.Name)
	return &HelloResponse{Message: "reply"}, nil
}

func TestNewGrpcConn(t *testing.T) {
	client, err := getEtcdClient()
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	conn, err := rpc.NewGrpcConn(ctx, "test", client)
	if err != nil {
		t.Error(err)
	}
	cli := NewGrpcTestClient(conn)
	res, err := cli.Hello(ctx, &HelloRequest{Name: "name"})
	log.Println(res, err)
}

func TestNewGrpcServer(t *testing.T) {
	client, err := getEtcdClient()
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	svc, err := rpc.NewGrpcServer(ctx, &rpc.Config{Name: "test", Bind: ":9999"}, client)
	if err != nil {
		t.Error(err)
	}
	RegisterGrpcTestServer(svc, grpcTest{})
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		t.Error(err)
	}
	_ = svc.Serve(lis)
}
