package router

import (
	"github.com/biningo/eagle/app/api"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/**
*@Author icepan
*@Date 7/20/21 10:48
*@Describe
**/

func InitRegistry(r *gin.Engine, etcdCli *clientv3.Client) {
	router := r.Group("/registry/:namespace")
	router.GET("/services", api.ListService(etcdCli))
	router.GET("/services/:serviceName", api.GetServiceByName(etcdCli))
	router.GET("/services/:serviceName/:serviceID", api.GetServiceInstance(etcdCli))
}
