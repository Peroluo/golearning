package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetGoodsList 控制器
func GetGoodsList(c *gin.Context) {
	// 定义一个长度为5的数组
	// arr := [5]int{1, 2, 3, 4, 5}
	// 数组循环
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("p[%d] == %d\n", i, arr[i])
	// }
	// 定义一个长度为5的slice
	slice1 :=[]int{2, 3, 5, 7, 11, 13}
	fmt.Println("p[:4] ==", slice1[:4])
	fmt.Println("p[3:] ==", slice1[3:])
	// 构造slice
	slice2:=make([]int, 10,20)
	// slice append在slice后添加元素
	slice3:=append(slice2,11) //
	fmt.Println("slice1=",slice1)
	slice1_2:=slice1[:3]
	slice1_1 :=slice1[4:]
	fmt.Println(slice1_1)
	fmt.Println("slice1=",slice1)
	slice4:=append(slice1_2,4)
	fmt.Println("slice1=",slice1)
	fmt.Println(slice2,slice3,slice4)
	// range循环
	// for i,v:=range s3{
	// 	fmt.Println(i,v)
	// }
	c.JSON(200, gin.H{
		"data": slice1,
	})
}
