package main

import "fmt"

/*
流程控制
顺序 分支 循环
*/
func main_day8() {
	var sum int
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 没有while循环
	sum1 := 0
	var j = 100
	for j > 0 {
		sum1 += j
		j--
	}
	fmt.Println(sum1)
}
