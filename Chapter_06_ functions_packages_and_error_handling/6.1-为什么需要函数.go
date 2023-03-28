package main

import (
	"fmt"
)

//使用传统的方法解决：输入两个数，再运算一个运算符(+,-,*,/)得出结果
func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}
	fmt.Println("res=",res)
}
