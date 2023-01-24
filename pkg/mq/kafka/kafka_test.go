package kafka

import (
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewKafka(t *testing.T) {
	conf := &Config{}
	err := config.LoadYaml("../../../internal/yaml/kafka.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr)
}
