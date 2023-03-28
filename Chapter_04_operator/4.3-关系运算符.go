package main

import (
	"fmt"
)

//演示关系运算符的使用
func func1() {
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	flag := n1 > n2
	fmt.Println("flag =", flag)
}

func main() {
	func1()
}
