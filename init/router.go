package initrouter

import (
	"github.com/gin-gonic/gin"
	"golearning/router"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	Router := gin.Default()
	APIRouters := Router.Group("")
	router.APIRouters(APIRouters)
	return Router
}
