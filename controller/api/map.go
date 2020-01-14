package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// TestMap TEST
func TestMap(c *gin.Context) {
	mapStr := map[string]string{"name": "pero", "age": "8"}
	mapStr2 := make(map[string]int)
	mapStr2["age"] = 8
	fmt.Println(mapStr2["age"])
	fmt.Println(mapStr["name"])
	// 刪除属性
	mapStr["add"] = "addValue"
	delete(mapStr, "name")

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	c.JSON(200, gin.H{
		"data": mapStr,
	})
}
