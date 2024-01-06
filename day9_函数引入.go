package main

import (
	"fmt"
)

func main_day9() {
	a := add(10, 20)
	fmt.Println(a)
	add1(a, 10)
	sum1, neg1 := cal2(10, 20)
	println(sum1)
	println(neg1)
	//忽略值
	sum2, _ := cal2(20, 30)
	println(sum2)

}

/*
*
形参列表: 一个或者n个
实际参数列表: 实际传入的数据
func 函数名 (形参列表)(返回列表) {
执行语句
return + 返回值列表
}
*/
func add(a int, b int) int {
	// 自定义函数, 功能: 相加
	var sum = 0
	sum += a
	sum += b
	return sum
}

/*
返回值只有一个时返回值旁边的括号可以不写
*/
func add1(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

/*
计算和差
*/
func cal2(a int, b int) (int, int) {
	sum := a + b
	neg := a - b
	return sum, neg
}
