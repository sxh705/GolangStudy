package main

import "fmt"

type sxh int64            //自定义数据类型
type myFunc func(int) int //函数类型

func main_day11() {
	var a = test2
	fmt.Printf("a %T %v\n", a, a)
	fmt.Printf("test2 %T %v\n", test2, test2)
	// 通过该变量调用函数
	a(1)
	test3(a)
	var b sxh = 0xfffffffffffffff //16进制
	b = 0b1001                    //二进制
	b = 0o1234                    //8进制
	fmt.Printf("b %T %v\n", b, b)
	var c int64 = int64(b) // 不能直接赋值 需要进行类型转换
	fmt.Printf("c %T %v\n", c, c)
	d := int64(1)
	fmt.Printf("d %T %v\n", d, d)
	var e myFunc = func(i int) int {
		return i
	}
	fmt.Printf("e %T %v\n", e, e)
	fmt.Printf("\"aaa\" %T %v\n", "aaa", "aaa")
}

func test2(x int) {
	fmt.Printf("x %v\n", x)
}

func test3(testFunc func(int)) {
	testFunc(3)
}

/*
函数返回值命名
*/
func test4() (a int, b string) {
	a = 1
	b = "aa"
	return
	// 命名的返回值赋值后直接return即可
}
