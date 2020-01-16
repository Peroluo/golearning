package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 引入mysql
	"golearning/config"
)

// DEFAULTDB DB全局变量
var DEFAULTDB *gorm.DB

// InitMysql 初始化数据库并产生数据库全局变量
func InitMysql(admin config.MysqlAdmin) *gorm.DB {
	db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config)
	if err != nil {
		fmt.Println("DEFAULTDB数据库启动异常%S", err)
	} else {
		DEFAULTDB = db
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		fmt.Println("DEFAULTDB数据库已启动！")
	}
	return db
}
