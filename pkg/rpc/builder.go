package rpc

import (
	"context"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/constant"
	"github.com/zxq97/gokit/pkg/watcher"
	"google.golang.org/grpc/resolver"
)

type builder struct {
	watcher watcher.Watcher
}

func newBuilder(watcher watcher.Watcher) resolver.Builder {
	return &builder{
		watcher: watcher,
	}
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.DefaultEtcdTimeout)
	defer cancel()
	kvs, err := b.watcher.Start(ctx)
	if err != nil {
		log.Println("resolver start err", err)
		return nil, err
	}
	addr := make([]resolver.Address, 0, len(kvs))
	for _, v := range kvs {
		svc := &Service{}
		err = proto.Unmarshal(v.Value, svc)
		if err != nil {
			continue
		}
		addr = append(addr, resolver.Address{
			Addr:       svc.Bind,
			ServerName: svc.Name,
		})
	}
	r := &discoverResolver{
		addr:    addr,
		cc:      cc,
		watcher: b.watcher,
	}
	err = cc.UpdateState(resolver.State{Addresses: r.addr})
	go r.watch()
	return r, nil
}

func (b *builder) Scheme() string {
	return "resolver"
}
