package rpc

import (
	"context"
	"log"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/watcher"
	"google.golang.org/grpc/resolver"
)

type discoverResolver struct {
	addr    []resolver.Address
	cc      resolver.ClientConn
	watcher watcher.Watcher
}

func (r *discoverResolver) Close() {
	_ = r.watcher.Stop()
}

func (r *discoverResolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (r *discoverResolver) watch() {
	f := func(key string, val []byte) {
		svc := &Service{}
		err := proto.Unmarshal(val, svc)
		if err != nil {
			log.Println("resolver put unmarshal err", err)
			return
		}
		r.addr = append(r.addr, resolver.Address{Addr: svc.Bind, ServerName: svc.Name})
		_ = r.cc.UpdateState(resolver.State{Addresses: r.addr})
	}
	g := func(key string) {
		for k := range r.addr {
			s := strings.Split(key, "/")
			if len(s) < 3 {
				log.Println("resolver delete not addr")
				continue
			}
			if r.addr[k].Addr == s[2] {
				r.addr = append(r.addr[:k], r.addr[k+1:]...)
				break
			}
		}
		_ = r.cc.UpdateState(resolver.State{Addresses: r.addr})
	}
	errCh := r.watcher.Run(context.Background(), f, g)
	select {
	case err := <-errCh:
		if err != nil {
			log.Println("resolver stop err", errCh)
		}
	}
}
