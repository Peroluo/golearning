package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"golearning/model/sqlmodel"
)

// Student test
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Goods []Good `json:"goods"`
}

// Good 商品信息
type Good struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// GetHome 控制器
func GetHome(c *gin.Context) {
	todo := sqlmodel.TodoModel{
		Title:     "SS",
		Completed: 1,
	}
	err := todo.AddTodo()
	fmt.Println(err)
	// goods := []Good{Good{Name: "商品名称", Price: 1}}
	// stu := Student{Name: "wd", Age: 22, Goods: goods}
	c.JSON(200, gin.H{
		"data": todo,
	})
}
