package rocketmq

import (
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewRocketmq(t *testing.T) {
	conf := &Config{}
	err := config.LoadYaml("../../../internal/yaml/rocketmq.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr, conf.NsAddr)
}
