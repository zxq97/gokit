package etcd

import (
	"context"

	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
)

type Watcher struct {
	key       string
	watchChan clientv3.WatchChan
	watcher   clientv3.Watcher
	kv        clientv3.KV
}

func NewWatcher(ctx context.Context, key string, client *clientv3.Client) (*Watcher, error) {
	w := &Watcher{
		key:     key,
		watcher: clientv3.NewWatcher(client),
		kv:      clientv3.NewKV(client),
	}
	w.watchChan = w.watcher.Watch(ctx, key, clientv3.WithPrefix())
	err := w.watcher.RequestProgress(context.Background())
	return w, err
}

func (w *Watcher) Start(ctx context.Context) ([]*mvccpb.KeyValue, error) {
	res, err := w.kv.Get(ctx, w.key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	return res.Kvs, nil
}

// Run
// f func put  g func delete
func (w *Watcher) Run(ctx context.Context, f func(key string, val []byte), g func(key string)) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		for {
			select {
			case res, ok := <-w.watchChan:
				if !ok {
					close(errCh)
					return
				}
				for _, v := range res.Events {
					k := string(v.Kv.Key)
					switch v.Type {
					case mvccpb.PUT:
						f(k, v.Kv.Value)
					case mvccpb.DELETE:
						g(k)
					}
				}
			case <-ctx.Done():
				errCh <- w.Stop()
				return
			}
		}
	}()
	return errCh
}

func (w *Watcher) Stop() error {
	return w.watcher.Close()
}
