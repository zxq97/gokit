package xmemcache

import (
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewMemcache(t *testing.T) {
	conf := &Config{}
	err := config.LoadYaml("../../../internal/yaml/memcache.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr)
	client := NewMemcache(conf)
	res, err := client.Get("k")
	log.Println(res, err)
}
