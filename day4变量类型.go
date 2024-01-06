package main

/*
golang数据类型
基本数据类型:

	数值型
		整数int8-int64 uint8-uint64
	浮点型
		float32 float64
	字符型 byte(=uint8)
	布尔型 bool
	字符串 string

派生数据类型/复杂数据类型

	指针
	数组
	结构体
	管道
	函数
	切片
	接口
	map

有符号整数类型:
int8	-128-127(-2^7-2^7-1
int16	-32768-32767(-2^15-2^15-1
int32	-2147483648-2147483647(-2^31~2^31-1
int64	-2^63-2^63-1

无符号整数类型
uint8	0-255
uint16	0-2^16-1
uint32	0-2^31-1
uint64	0-2^63-1


*/

import "fmt"
import (
	"unsafe"
)

func main_day4() {
	var num1 int8 = 23
	//num1 = 230 //超范围会编译出错
	fmt.Println(num1)

	var num2 uint8
	//num2 = -20
	fmt.Println(num2)

	var num3 = 28
	fmt.Printf("num3的类型:%T\n", num3)
	// https://golang.org
	// https://studygolang.com/pkgdoc
	fmt.Printf("num3占用的比特数:%v\n", unsafe.Sizeof(num3))

}
