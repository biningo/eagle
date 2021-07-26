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

func InitConfig(r *gin.Engine, etcdCli *clientv3.Client) {
	router := r.Group("/config/:namespace")
	router.GET("/configurations", api.ListConfiguration(etcdCli))
	router.GET("/configurations/:filename", api.GetConfiguration(etcdCli))
	router.PUT("/configurations/:filename", api.UploadConfiguration(etcdCli))
	router.DELETE("/configurations/:filename", api.DelConfiguration(etcdCli))
}
