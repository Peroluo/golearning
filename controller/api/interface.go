package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// MoneyCreate MoneyCreate
type MoneyCreate interface {
	create() int
}

// DevFix DevFix
type DevFix struct {
	Base  int
	Other int
}

//DevBase DevBase
type DevBase struct {
	Base int
}

func (d DevFix) create() int {
	return d.Base + d.Other
}

func (d DevBase) create() int {
	return d.Base
}

// GetInter 控制器
func GetInter(c *gin.Context) {
	devFix1 := DevFix{Base: 100, Other: 10}
	devFix2 := DevFix{200, 20}
	devBase1 := DevBase{150}
	pers := [3]MoneyCreate{devFix1, devFix2, devBase1}
	total := 0
	for _, m := range pers {
		total += m.create()
	}
	fmt.Println(total)
	c.JSON(200, gin.H{
		"data": pers,
	})
}
