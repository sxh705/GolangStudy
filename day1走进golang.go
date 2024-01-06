package main // 每个go文件都必须有归属包,主函数必须属于main包
import "fmt" // 为了使用包下的函数, 引入程序所需的包

// 每个文件夹只能有一个包
// 包名从GOPATH/src后开始计算 使用/进行分离

func main_day1() { // main 主函数 入口
	fmt.Println("hello" + "world") // 在控制台打印一句话
}

// 编译使用 go build day1 xxx.go

// 运行使用 go run day1 xxx.go
