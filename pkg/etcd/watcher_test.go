package etcd

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewWatcher(t *testing.T) {
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
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	w, err := NewWatcher(ctx, "k", client)
	if err != nil {
		t.Error(err)
	}
	ch := w.Run(ctx, func(key string, val []byte) {
		log.Println(key, string(val))
	}, func(key string) {
		log.Println(key)
	})
	go func() {
		<-time.After(time.Minute)
		if err = w.Stop(); err != nil {
			t.Error(err)
		}
	}()
	select {
	case err = <-ch:
		if err != nil {
			t.Error(err)
		}
	case <-ctx.Done():
		if err != nil {
			t.Error(ctx.Err())
		}
	}
}
