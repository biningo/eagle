package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/biningo/eagle/etcd"
	"github.com/biningo/eagle/registry"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/**
*@Author icepan
*@Date 7/26/21 14:26
*@Describe
**/

func GetConfigResultsFromEtcd(etcdCli *clientv3.Client, key string) (results []etcd.Config, err error) {
	resp, err := etcdCli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return
	}
	var r etcd.Config
	for _, kv := range resp.Kvs {
		r.Filename = string(kv.Key)
		r.Content = string(kv.Value)
		results = append(results, r)
	}
	return results, nil
}

func GetServiceResultsFromEtcd(etcdCli *clientv3.Client, key string) ([]gin.H, error) {
	resp, err := etcdCli.Get(context.TODO(), key, clientv3.WithPrefix())
	if err != nil {
		return []gin.H{}, err
	}
	results := make([]gin.H, len(resp.Kvs))
	for i, kv := range resp.Kvs {
		svc := registry.ServiceInstance{}
		if err := json.Unmarshal(kv.Value, &svc); err != nil {
			fmt.Println(err)
			continue
		}
		body := gin.H{
			"key":     string(kv.Key),
			"service": svc,
		}
		results[i] = body
	}
	return results, nil
}
