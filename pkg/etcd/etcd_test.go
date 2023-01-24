package etcd

import (
	"context"
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewEtcd(t *testing.T) {
	conf := &Config{}
	err := config.LoadYaml("../../internal/yaml/etcd.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr, conf.TTL)
	client, err := NewEtcd(conf)
	if err != nil {
		t.Error(err)
	}
	res, err := client.Get(context.TODO(), "k")
	if err != nil {
		t.Error(err)
	}
	log.Println(res)
}
