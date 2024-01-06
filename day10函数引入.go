package main

import "fmt"

/*
golang 内存逻辑划分
栈:基本数据类型
堆:引用数据类型 复杂数据类型
代码区:代码本身
函数执行结束后
程序会销毁函数对应的栈空间
*/
func main_day10() {
	num1, num2 := 10, 20
	println(num1)
	println(num2)
	swap(&num1, &num2)
	println(num1)
	println(num2)

	test()
	test(3)
	test(37, 68, 23)
}

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

// golang不支持函数重载
//
//	func swap(a *int) {
//		fmt.Println(a)
//	}
//
// 但是支持可变参数
func test(args ...int) {
	// 处理可变参数时会作为切品处理
	fmt.Println(args)

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	for i, arg := range args {
		fmt.Println(i, arg)
	}

}
