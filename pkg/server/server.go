package server

import (
	"context"
	"time"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func WithStartTimeout(d time.Duration) Option {
	return func(s *ServiceMeta) {
		s.StartTimeout = d
	}
}

func WithStopTimeout(d time.Duration) Option {
	return func(s *ServiceMeta) {
		s.StopTimeout = d
	}
}

type Option func(s *ServiceMeta)

type ServiceMeta struct {
	StartTimeout time.Duration
	StopTimeout  time.Duration
}
