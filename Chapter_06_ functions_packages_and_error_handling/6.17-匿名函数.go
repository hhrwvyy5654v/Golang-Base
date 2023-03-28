package main

import (
	"fmt"
)

//Go支持匿名函数，匿名函数就是没有名字的函数。如果我们希望某个函数只是使用希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用

//匿名函数使用方式1:在定义匿名函数时就直接调用，这种匿名函数只能调用一次
func func1() {
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)
}

//匿名函数使用方式2:将匿名函数赋给一个变量(函数变量),再通过该变量来调用匿名函数
func func2() {
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res1 := a(10, 30)
	fmt.Println("res1=", res1)
	res2 := a(90, 30)
	fmt.Println("res2=", res2)
}

//匿名函数使用方式2:如果将匿名函数赋给一个全局变量，那么这个匿名函数就成为一个全局匿名函数，可以在程序有效
var (
	Func3 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//全局匿名函数的使用
	res := Func3(4, 9)
	fmt.Println("res=", res)
}
