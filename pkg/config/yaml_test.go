package config

import (
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/etcd"
)

func TestLoadYaml(t *testing.T) {
	path := "../../internal/yaml/etcd.yaml"
	conf := etcd.Config{}
	err := LoadYaml(path, &conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr, conf.TTL)
}
