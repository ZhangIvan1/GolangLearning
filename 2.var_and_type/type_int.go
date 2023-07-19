package main

/*
整型
int8	有符号	1字节	-2^7~2^7-1(-128~127)
int16	有符号	2字节	-2^15~2^15-1(-32768~32767)
int32	有符号	4字节	-2^31~2^31-1(-2147483648~2147483647)
int64	有符号	8字节	-2^63~2^63-1
uint8	无符号	1字节	0~255
uint16	无符号	2字节	0~2^16-1
uint32	无符号	4字节	0~2^31-1
uint 64	无符号	8字节	0~2^63-1
**/

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个整数类型:var num1 int8 = 120fmt.println(num1)
	var num2 uint8 = 200
	fmt.Println(num2)
	var num3 = 28
	// Printf函数的作用就是:格式化的，把num3的类型填充到%T的位置上
	fmt.Printf("num3的类型是:%T \n", num3)
	//num3的类型是:intfmt Drint1nll
	fmt.Println(unsafe.Sizeof(num3))
}
