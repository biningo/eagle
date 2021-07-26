package api

import (
	"fmt"
	"github.com/biningo/eagle/app/service"
	"github.com/biningo/eagle/internal/config"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/**
*@Author icepan
*@Date 7/19/21 17:17
*@Describe
**/

func ListService(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		key := fmt.Sprintf("%s/registry/%s", config.Conf.Prefix, namespace)
		results, err := service.GetServiceResultsFromEtcd(etcdCli, key)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, results)
	}
}

func GetServiceByName(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		name := ctx.Param("serviceName")
		key := fmt.Sprintf("%s/registry/%s/%s", config.Conf.Prefix, namespace, name)
		results, err := service.GetServiceResultsFromEtcd(etcdCli, key)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, results)
	}
}

func GetServiceInstance(etcdCli *clientv3.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		name := ctx.Param("serviceName")
		id := ctx.Param("serviceID")
		key := fmt.Sprintf("%s/registry/%s/%s/%s", config.Conf.Prefix, namespace, name, id)
		results, err := service.GetServiceResultsFromEtcd(etcdCli, key)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(200, results)
	}
}
