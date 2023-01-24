package mq

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MetricEventHandlerMsgDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "event_handler_msg_duration",
		Help:    "消息处理延时",
		Buckets: prometheus.DefBuckets,
	}, []string{"name", "topic", "status", "reason"})
)
