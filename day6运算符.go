package main

import "fmt"

/*
运算符
1.算法运算符	+ - * / % ++ --
2.赋值运算符	= += -= *= /= %=
3.关系运算符	== != > < >= <=
4.逻辑运算符	&& || !
5.位运算符	& | ^ << >>
6.其他运算符	& *
*/
func main_day6() {
	var a = 0
	a++ // a++不返回值,没有++a
	var f = a
	fmt.Println(f)
	var b = &a // 取地址
	_, err := fmt.Scan(b)
	if err != nil {
		fmt.Println(err)
	}
	a = a << 4 // 移位
	fmt.Println(a)
	fmt.Println("地址", b)
	fmt.Println(*b) //取值
	fmt.Printf("%T\n", b)
}
