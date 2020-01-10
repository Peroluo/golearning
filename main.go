package main

import (
	s "golearning/src/testruct"

	"github.com/gin-gonic/gin"
)

func main() {
	goodT := s.Goot{IsGood: true, Teacher: s.Teacher{Name: "pero", Age: 30}}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": goodT,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
