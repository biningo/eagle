package api

import (
	"context"
	"fmt"
	"github.com/biningo/eagle/app/service"
	"github.com/biningo/eagle/etcd"
	"github.com/biningo/eagle/internal/config"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/**
*@Author icepan
*@Date 7/19/21 17:17
*@Describe
**/

func ListConfiguration(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		key := fmt.Sprintf("%s/config/%s", config.Conf.Prefix, namespace)
		results, err := service.GetConfigResultsFromEtcd(etcdCli, key)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, results)
	}
}

func GetConfiguration(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		filename := ctx.Param("filename")
		key := fmt.Sprintf("%s/config/%s/%s", config.Conf.Prefix, namespace, filename)
		results, err := service.GetConfigResultsFromEtcd(etcdCli, key)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, results)
	}
}

func DelConfiguration(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		filename := ctx.Param("filename")
		key := fmt.Sprintf("%s/cofnig/%s/%s", config.Conf.Prefix, namespace, filename)
		if _, err := etcdCli.Delete(context.Background(), key, clientv3.WithPrefix()); err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, "OK")
	}
}

func UploadConfiguration(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		filename := ctx.Param("filename")
		key := fmt.Sprintf("%s/config/%s/%s", config.Conf.Prefix, namespace, filename)
		var c etcd.Config
		if err := ctx.BindJSON(&c); err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		if _, err := etcdCli.Put(context.Background(), key, c.Content); err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, "OK")
	}
}
