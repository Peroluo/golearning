package main

import (
	"golearning/inits"
	"net/http"
	"time"
)

func main() {
	inits.InitLog()
	_ = inits.InitRedis() // 初始化redis服务
	Router := inits.InitRouter()
	inits.QMLog.Info("服务器开启") // 日志测试代码
	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
