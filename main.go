package main

import (
	"golearning/config"
	"golearning/inits/mysql"
	"golearning/inits/qmlog"
	"golearning/inits/redis"
	"golearning/inits/regist"
	"golearning/inits/router"
	"net/http"
	"time"
)

func main() {
	qmlog.InitLog()                                            // 初始化日志
	db := mysql.InitMysql(config.GinVueAdminconfig.MysqlAdmin) // 初始化数据库
	_ = redis.InitRedis()                                      // 初始化redis服务
	regist.InitTable(db)                                       // 注册库表
	defer mysql.DEFAULTDB.Close()                              // 程序结束前关闭数据库链接
	Router := router.InitRouter()                              // 路由注册
	// qmlog.QMLog.Info("服务器开启")                                  // 日志测试代码
	s := &http.Server{
		Addr:           ":5000",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
