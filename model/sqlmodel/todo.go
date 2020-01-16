package sqlmodel

import (
	"github.com/jinzhu/gorm"
	"golearning/inits/mysql"
)

// TodoModel 定义原始的数据库字段
type TodoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// AddTodo AddTodo
func (t *TodoModel) AddTodo() (err error) {
	err = mysql.DEFAULTDB.Create(t).Error
	return err
}
