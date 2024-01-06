package main

import (
	"GolangStudy/day12packageControll/crm"
	"fmt"
)

// 一个包下不能定义重名函数
func main_day12() {
	fmt.Println("包测试")
	crm.GetConn()
}
