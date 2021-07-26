package router

/**
*@Author icepan
*@Date 7/19/21 17:11
*@Describe
**/

import (
	"github.com/biningo/eagle/internal/config"
	"github.com/biningo/eagle/utils"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func initEtcd() *clientv3.Client {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: config.Conf.EtcdConfig.Endpoints,
	})
	utils.CheckError(err)
	return etcdCli
}

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})
	etcdCli := initEtcd()
	InitRegistry(r, etcdCli)
	InitConfig(r, etcdCli)
	return r
}
