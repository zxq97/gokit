package mq

import (
	"log"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ApiLogger *log.Logger
	ExcLogger *log.Logger

	MetricEventHandlerMsgDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "event_handler_msg_duration",
		Help:    "消息处理延时",
		Buckets: prometheus.DefBuckets,
	}, []string{"name", "topic", "status", "reason"})
)

func init() {
	ApiLogger = log.New(os.Stderr, "API\t", log.Llongfile|log.LstdFlags)
	ExcLogger = log.New(os.Stderr, "EXC\t", log.Llongfile|log.LstdFlags)
}

func ObserveMsg(topic, name, status, reason string, start time.Time) {
	MetricEventHandlerMsgDuration.WithLabelValues(name, topic, status, reason).Observe(time.Since(start).Seconds())
}
