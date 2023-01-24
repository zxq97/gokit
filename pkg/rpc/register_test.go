package rpc

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/zxq97/gokit/pkg/config"
	"github.com/zxq97/gokit/pkg/etcd"
)

func TestNewRegister(t *testing.T) {
	conf := &etcd.Config{}
	err := config.LoadYaml("../../internal/yaml/etcd.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr, conf.TTL)
	client, err := etcd.NewEtcd(conf)
	if err != nil {
		t.Error(err)
	}
	r := newRegister(client)
	ctx, cancel := context.WithCancel(context.Background())
	svc := &Service{Name: "name", Bind: "9999"}
	err = r.Register(ctx, svc)
	if err != nil {
		t.Error(err)
	}
	<-time.After(time.Minute)
	cancel()
	<-time.After(time.Second * 2)
}
