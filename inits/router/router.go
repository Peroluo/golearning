package router

import (
	"github.com/gin-gonic/gin"
	"golearning/middleware"
	"golearning/router"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Logger()) // 如果不需要日志 请关闭这里
	APIRouters := Router.Group("")
	router.APIRouters(APIRouters)
	return Router
}
