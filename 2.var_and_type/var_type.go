package main

import (
	"fmt"
	"reflect"
)

//全局变量
var (
	n1 = 4546
	n2 = 1345.5453
	n3 = "545sfd"
)

func main() {
	var age int // =18
	fmt.Println("age = ", age)
	age = 18
	fmt.Println("age = ", age)

	//var age2 int = 12.5
	//fmt.Println(age2)
	//# command-line-arguments
	//./var_type.go:10:6: constant 12.5 truncated to integer

	var age3 = 12
	fmt.Println(age3, reflect.TypeOf(age3))

	age4 := "4525"
	fmt.Println(age4, reflect.TypeOf(age4))

	var a, b, c, d float32
	fmt.Println(a, b, c, d)

	a1, b1, c1, d1 := 12, 12.251, "2131fsf", true
	fmt.Println(a1, b1, c1, d1)

	fmt.Println(n1, n2, n3)
}
