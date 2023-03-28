package main

import (
	"fmt"
)

//编写一个函数swap(n1 *int ,n2 *int)可以交换n1和n2的值
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%v,b=%v", a, b)
}
