package kafka

import (
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/mq"
)

type Config struct {
	Addr []string `yaml:"addr"`
}

var (
	apiLogger *log.Logger
	excLogger *log.Logger
)

func init() {
	apiLogger = log.New(os.Stdout, "API\t", log.Llongfile|log.LstdFlags)
	excLogger = log.New(os.Stderr, "EXC\t", log.Llongfile|log.LstdFlags)
}

func unmarshal(message []byte) (*KafkaMessage, error) {
	msg := &KafkaMessage{}
	err := proto.Unmarshal(message, msg)
	return msg, err
}

func observeMsg(msg *sarama.ConsumerMessage, name, status, reason string, start time.Time) {
	mq.MetricEventHandlerMsgDuration.WithLabelValues(name, msg.Topic, status, reason).Observe(time.Since(start).Seconds())
}
