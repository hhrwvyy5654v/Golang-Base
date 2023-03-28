package main

import (
	"fmt"
)

//Go语言里没有while和do…while语法，可以通过for循环来实现其效果
func func1() {
	var i int = 1
	for {
		fmt.Println("hello world", i)
		i++	//循环变量迭代
		if i > 10 {
			break
		}
	}
}

func main() {
	func1()
}
