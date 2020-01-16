package regist

import (
	"github.com/jinzhu/gorm"
	"golearning/model/sqlmodel"
)

//InitTable 库表初始化
func InitTable(db *gorm.DB) {
	db.AutoMigrate(sqlmodel.TodoModel{})
}
