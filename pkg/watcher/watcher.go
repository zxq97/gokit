package watcher

import (
	"context"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type Watcher interface {
	Start(ctx context.Context) ([]*mvccpb.KeyValue, error)
	Run(ctx context.Context, f func(key string, val []byte), g func(key string)) <-chan error
	Stop() error
}
