package consumer

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	kittest "github.com/zxq97/gokit/internal/kit_test/rocketmq"
	"github.com/zxq97/gokit/pkg/config"
	"github.com/zxq97/gokit/pkg/mq"
	"github.com/zxq97/gokit/pkg/mq/rocketmq"
	"github.com/zxq97/gokit/pkg/server"
)

func TestNewServer(t *testing.T) {
	conf := &rocketmq.Config{}
	err := config.LoadYaml("../../../internal/yaml/rocketmq.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}
	fn := func(ctx context.Context, msg *mq.MqMessage) error {
		rm := &kittest.TestMsg{}
		if err = proto.Unmarshal(msg.Message, rm); err != nil {
			t.Error(err)
		}
		log.Println(rm)
		return nil
	}
	c1, err := rocketmq.NewConsumer(conf, "test_tt", "g1", fn, mq.WithProcTimeout(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	c2, err := rocketmq.NewConsumer(conf, "test_tt", "g2", fn, mq.WithProcTimeout(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	s, err := NewServer([]mq.Consumer{c1, c2}, server.WithStartTimeout(time.Second), server.WithStopTimeout(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if err = s.Start(context.TODO()); err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	time.AfterFunc(time.Second*3, func() {
		if err = s.Stop(context.TODO()); err != nil {
			t.Fatal(err)
		}
		close(done)
	})
	<-done
}
