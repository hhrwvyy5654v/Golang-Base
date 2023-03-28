package main

import (
	"fmt"
)

func cal(n1 float64, n2 float64, operator byte) float64 {

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
	return res
}

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.6
	var operator byte = '/'
	result := cal(n1, n2, operator)
	fmt.Println(result)
}
