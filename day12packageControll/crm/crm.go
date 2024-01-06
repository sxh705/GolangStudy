package crm

import "fmt"

// 建议package后面的包声明与所在文件夹同名
// main是入口包 一般main会放在该包下

func GetConn() {
	// 首字母大写是public
	fmt.Println("执行了crm.GetConn")
}
