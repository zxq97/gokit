package rocketmq

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/config"
	"github.com/zxq97/gokit/pkg/mq"
	"github.com/zxq97/gokit/pkg/mq/rocketmq"
)

func TestNewProducer(t *testing.T) {
	conf := &rocketmq.Config{}
	err := config.LoadYaml("../../yaml/rocketmq.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}
	p, err := rocketmq.NewProducer(conf)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		msg := &TestMsg{A: "a", B: int64(i + 100)}
		if err = p.SendMessage(context.TODO(), "test_tt", "key", mq.TagCreate, msg, 0); err != nil {
			t.Error(err)
		}
	}
	if err = p.Close(); err != nil {
		t.Error(err)
	}
}

func TestNewConsumer(t *testing.T) {
	conf := &rocketmq.Config{}
	err := config.LoadYaml("../../yaml/rocketmq.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}
	fn := func(ctx context.Context, msg *mq.MqMessage) error {
		rm := &TestMsg{}
		if err = proto.Unmarshal(msg.Message, rm); err != nil {
			t.Error(err)
		}
		log.Println(rm)
		return nil
	}
	c, err := rocketmq.NewConsumer(conf, "test_tt", "g", fn, mq.WithProcTimeout(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if err = c.Start(); err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	time.AfterFunc(time.Second*5, func() {
		if err = c.Stop(); err != nil {
			t.Error(err)
		}
		close(done)
	})
	<-done
}
