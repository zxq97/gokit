package kafka

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/config"
	"github.com/zxq97/gokit/pkg/mq/kafka"
)

func TestNewProducer(t *testing.T) {
	conf := &kafka.Config{}
	err := config.LoadYaml("../../yaml/kafka.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr)
	p, err := kafka.NewProducer(conf)
	if err != nil {
		t.Error(err)
	}
	ticker := time.NewTicker(time.Millisecond * 10)
	done := make(chan struct{})
	go func() {
		time.AfterFunc(time.Second*10, func() {
			close(done)
		})
	}()
	var x int64 = 10000
	for {
		select {
		case <-ticker.C:
			msg := &TestMsg{A: "a", B: x}
			log.Println(x)
			x++
			err = p.SendMessage(context.TODO(), "test_tt", "key", msg, kafka.EventTypeCreate)
			if err != nil {
				t.Error(err)
			}
		case <-done:
			return
		}
	}
}

func TestNewConsumer(t *testing.T) {
	conf := &kafka.Config{}
	err := config.LoadYaml("../../yaml/kafka.yaml", conf)
	if err != nil {
		t.Error(err)
	}
	log.Println(conf.Addr)
	c, done, err := kafka.NewConsumer(conf, []string{"test_tt"}, "g", "test", func(ctx context.Context, msg *kafka.KafkaMessage) error {
		km := &TestMsg{}
		err = proto.Unmarshal(msg.Message, km)
		if err != nil {
			t.Error(err)
		}
		log.Println(km)
		<-time.After(time.Second)
		return nil
	}, 1, 1, time.Second)
	if err != nil {
		t.Error(err)
	}
	c.Start()
	time.AfterFunc(time.Second*3, func() {
		if err = c.Stop(); err != nil {
			t.Error(err)
		}
	})
	log.Println("wait done chan")
	<-done
}