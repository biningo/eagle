package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/biningo/eagle/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

/**
*@Author icepan
*@Date 7/19/21 14:01
*@Describe
**/

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ttl    time.Duration
	prefix string
}

// Registry is etcd registry.
type Registry struct {
	opts   *options
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

// NewRegistry creates etcd registry
func NewRegistry(client *clientv3.Client, opts ...Option) *Registry {
	options := &options{
		ttl:    time.Second * 15,
		prefix: "",
	}
	for _, o := range opts {
		o(options)
	}
	return &Registry{
		opts:   options,
		client: client,
		kv:     clientv3.NewKV(client),
	}
}

// RegisterTTL with register ttl.
func RegisterTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.ttl = ttl
	}
}

// Prefix with register prefix
func Prefix(prefix string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

// Register the registration.
func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	key := r.ServiceKey(service)
	value, err := json.Marshal(service)
	if err != nil {
		return err
	}
	if r.lease != nil {
		r.lease.Close()
	}
	r.lease = clientv3.NewLease(r.client)
	grant, err := r.lease.Grant(ctx, int64(r.opts.ttl))
	if err != nil {
		return err
	}
	_, err = r.kv.Put(ctx, key, string(value), clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}

	hb, err := r.client.KeepAlive(ctx, grant.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case _, ok := <-hb:
				if !ok {
					return
				}
			}
		}
	}()
	return nil
}

// Deregister the registration.
func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	defer func() {
		if r.lease != nil {
			r.lease.Close()
		}
	}()
	key := r.ServiceKey(service)
	_, err := r.client.Delete(ctx, key)
	return err
}

// GetService return the service instances according to the service name.
func (r *Registry) GetService(ctx context.Context, opts ...registry.ServiceOption) ([]*registry.ServiceInstance, error) {
	svc := &registry.ServiceInstance{}
	svc.Service.Namespace = "default"
	for _, o := range opts {
		o(svc)
	}
	key := fmt.Sprintf("%s/%s/%s", r.opts.prefix, svc.Service.Namespace, svc.Service.Name)
	resp, err := r.kv.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var instances []*registry.ServiceInstance
	for _, kv := range resp.Kvs {
		if err := json.Unmarshal(kv.Value, svc); err != nil {
			return nil, err
		}
		instances = append(instances, svc)
	}
	return instances, nil
}

func (r *Registry) HealthCheck(ctx context.Context, service *registry.ServiceInstance) bool {
	if ok := service.Service.Check.Checker.Ping(); !ok {
		return false
	}
	return true
}

func (r *Registry) ServiceKey(service *registry.ServiceInstance) string {
	return fmt.Sprintf("%s/%s/%s/%s", r.opts.prefix, service.Service.Namespace, service.Service.Name, service.ID)
}
