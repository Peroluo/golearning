package router

import (
	"github.com/gin-gonic/gin"
	"golearning/controller/api"
)

// APIRouters 是
func APIRouters(Router *gin.RouterGroup) {
	v1 := Router.Group("")
	{
		v1.GET("/", api.GetHome)
		v1.GET("/getGoodsList", api.GetGoodsList)
		v1.GET("/getMap", api.TestMap)
		v1.GET("/getInter", api.GetInter)
	}
}
