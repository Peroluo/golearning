package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
// TestMap TEST
func TestMap(c *gin.Context){
	mapStr:=map[string]string{"name":"pero","age":"8"}
	fmt.Println(mapStr["name"])
	// 刪除属性
	mapStr["add"]="addValue"
	delete(mapStr,"name")
	c.JSON(200, gin.H{
		"data": mapStr,
	})
}