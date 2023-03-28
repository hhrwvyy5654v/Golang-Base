package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1的类型:%T\nnum1的值:%v\nnum1的地址:%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Printf("\nnum2的类型%T\nnum2的值=%v\nnum2的地址%v\nnum2这个指针,指向的值=%v", num2, num2, &num2, *&num2)
}
