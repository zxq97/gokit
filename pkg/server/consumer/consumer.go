package consumer

import (
	"context"

	"github.com/zxq97/gokit/pkg/mq"
	"github.com/zxq97/gokit/pkg/server"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	*server.ServiceMeta
	consumers []mq.Consumer
}

func NewServer(cs []mq.Consumer, opts ...server.Option) (*Server, error) {
	meta := &server.ServiceMeta{}
	for _, o := range opts {
		o(meta)
	}
	s := &Server{
		ServiceMeta: meta,
		consumers:   make([]mq.Consumer, len(cs)),
	}
	for k := range cs {
		s.consumers[k] = cs[k]
	}
	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	var cancel context.CancelFunc
	if s.StartTimeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, s.StartTimeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	eg, _ := errgroup.WithContext(ctx)
	for k := range s.consumers {
		c := s.consumers[k]
		eg.Go(c.Start)
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- eg.Wait()
	}()
	select {
	case <-ctx.Done():
		cancel()
		return ctx.Err()
	case err := <-errCh:
		cancel()
		return err
	}
}

func (s *Server) Stop(ctx context.Context) error {
	var cancel context.CancelFunc
	if s.StopTimeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, s.StopTimeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	eg, _ := errgroup.WithContext(ctx)
	for k := range s.consumers {
		c := s.consumers[k]
		eg.Go(c.Stop)
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- eg.Wait()
	}()
	select {
	case <-ctx.Done():
		cancel()
		return ctx.Err()
	case err := <-errCh:
		cancel()
		return err
	}
}
