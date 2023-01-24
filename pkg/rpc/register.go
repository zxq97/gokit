package rpc

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/constant"
	"go.etcd.io/etcd/client/v3"
)

const (
	defaultTTL      = 10 * time.Second
	defaultMaxRetry = 5

	serviceKey = "service_key"
)

type Register struct {
	client *clientv3.Client
	lease  clientv3.Lease
}

func newRegister(client *clientv3.Client) *Register {
	return &Register{
		client: client,
	}
}

// Register
// ctx Can only be cancelled and cannot carry the timeout
func (r *Register) Register(ctx context.Context, svc *Service) error {
	key := GetKey(svc)
	val, err := proto.Marshal(svc)
	if err != nil {
		return err
	}
	if r.lease != nil {
		_ = r.lease.Close()
	}
	ctx = context.WithValue(ctx, serviceKey, key)
	r.lease = clientv3.NewLease(r.client)
	leaseID, err := r.store(ctx, key, string(val))
	if err != nil {
		return err
	}
	go r.keepalive(ctx, leaseID, key, string(val))
	return nil
}

func (r *Register) store(ctx context.Context, key, val string) (clientv3.LeaseID, error) {
	grant, err := r.client.Grant(ctx, int64(defaultTTL.Seconds()))
	if err != nil {
		return 0, err
	}
	_, err = r.client.Put(ctx, key, val, clientv3.WithLease(grant.ID))
	return grant.ID, err
}

func (r *Register) keepalive(ctx context.Context, leaseID clientv3.LeaseID, key, val string) {
	curID := leaseID
	kac, err := r.client.KeepAlive(ctx, leaseID)
	if err != nil {
		curID = 0
	}
	for {
		if curID == 0 {
			for cnt := 0; cnt < defaultMaxRetry; cnt++ {
				ctx = context.WithValue(context.Background(), serviceKey, key)
				idCh := make(chan clientv3.LeaseID)
				errCh := make(chan error)
				go func() {
					nCtx, cancel := context.WithTimeout(ctx, constant.DefaultEtcdTimeout)
					defer cancel()
					nid, nErr := r.store(nCtx, key, val)
					if nErr != nil {
						errCh <- nErr
					} else {
						idCh <- nid
					}
				}()
				select {
				case err = <-errCh:
					log.Println("retry store err", err)
				case curID = <-idCh:
					kac, err = r.client.KeepAlive(ctx, curID)
					if err == nil {
						break
					}
				}
				<-time.After(time.Duration(1<<cnt) * time.Second)
			}
			if _, ok := <-kac; !ok {
				log.Println("retry failed")
				return
			}
		}

		select {
		case _, ok := <-kac:
			if !ok {
				if errors.Is(ctx.Err(), context.Canceled) {
					ctx = context.WithValue(context.Background(), serviceKey, key)
					_ = r.Deregister(ctx)
					return
				}
				curID = 0
				continue
			}
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				curID = 0
				continue
			}
			ctx = context.WithValue(context.Background(), serviceKey, key)
			_ = r.Deregister(ctx)
			return
		}
	}
}

func (r *Register) Deregister(ctx context.Context) error {
	defer func() {
		if r.lease != nil {
			_ = r.lease.Close()
		}
	}()
	key, ok := ctx.Value(serviceKey).(string)
	if !ok {
		return nil
	}
	_, err := r.client.Delete(ctx, key)
	return err
}
