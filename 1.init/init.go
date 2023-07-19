package main //声明文件所在包，每个go文件必须有归属的包

import "fmt" //引入程序中需要用到的包

func main() { //主函数，程序入口
	fmt.Println("Hello World!!!")
}

//go build ./1.go 进行编译

//go run ./1.go  -编译并运行
