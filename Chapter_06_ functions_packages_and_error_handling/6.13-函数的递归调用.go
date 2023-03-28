package main

import (
	"fmt"
)

/*
一个函数在函数体内又调用了本身，称为递归调用
函数递归需要遵守的重要原则:
(1)执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
(2)函数的局部变量是独立的，不会相互影响
(3)递归必须向退出递归的条件逼近，否则就是无限递归
(4)当一个函数执行完毕，或者遇到return就会返回，遵守谁调用就将结果返回给谁
(5)当函数执行完毕或者返回时，该函数本身也会被系统销毁
*/

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

func main() {
	res := Fibonacci(5)
	fmt.Println(res)
}
