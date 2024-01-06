package main

import "fmt"

/*
流程控制
顺序 分支 循环
*/
func main_day7() {
	var a = true
	//_, _ = fmt.Scan(&a)
	// 1 和 true 是 true 其他是false
	if a {
		//if不加括号 必须写花括号
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if 可以加入变量定义(只能写一个定义)
	if i, j := 1, 2; i == j {
		fmt.Println("=")
	} else if i > j { // else 不能折行写
		fmt.Println(">")
	} else {
		fmt.Println("<")
	}

	var score = 100
	// switch可以用表达式, default可以在任意位置
	// 无需break
	// switch变量类型必须一致
	// switch穿透:fallthrough
	switch score / 10 {
	case 10:
		fmt.Println("A")
		fallthrough
	case 9:
		fmt.Println("B")
	case 8, 7, 6:
		//case可带多个值
		fmt.Println("C")
	default:
		fmt.Println("E")
	}

	//switch不带条件,可以当if使用 (switch true)
	var d = 12
	switch {
	case d == 11:
		fmt.Println("aaa")
	case d == 12:
		fmt.Println("bbb")
	case d == 13:
		fmt.Println("ccc")
	}
	fmt.Scan(&a)
}
