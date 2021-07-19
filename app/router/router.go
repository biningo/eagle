package router

/**
*@Author icepan
*@Date 7/19/21 17:11
*@Describe
**/

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})
	return r
}
