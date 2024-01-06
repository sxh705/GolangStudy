package main

// 变量类型

import "fmt"

// 全局变量
var n7 = 100

var n8 = 9.7

var (
	n9  = 500
	n10 = "netty"
)

func main_day3() { //变量的使用
	//局部变量
	var num int = 19
	fmt.Println(num)

	var num2 int
	fmt.Println(num2) // 不赋值是0

	var num3 = 10 // 不写类型会自动推断
	fmt.Println(num3)

	sex := "男" // 省略var,不能写为=
	fmt.Println(sex)

	fmt.Println("----------------")

	var n1, n2, n3 int //多变量声明
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, name, n5 = 10, "jack", 7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	n6, height := 6.9, 100.6
	fmt.Println(n6)
	fmt.Println(height)

	fmt.Println(n7)
	fmt.Println(n8)

	fmt.Println(n9)
	fmt.Println(n10)
}
