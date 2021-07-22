package router

import (
	"github.com/biningo/eagle/app/api"
	"github.com/gin-gonic/gin"
)

/**
*@Author icepan
*@Date 7/20/21 10:48
*@Describe
**/

func InitRegistry(r *gin.Engine) {
	router := r.Group("/registry/:namespace")
	router.GET("/services", api.ListService)
	router.GET("/services/:serviceName", api.GetServiceByName)
	router.GET("/services/:serviceName/:serviceID", api.GetServiceInstance)
}
