package main

import (
	"fmt"
)

/*
&:返回变量地址,比如&a就是给出变量的实际地址
* :指针变量,比如*a是一个指针变量
*/

func demo1() {
	a := 100
	fmt.Println("a的地址=", &a)
	
	var ptr *int=&a
	fmt.Println("ptr指向的值是=",*ptr)
}

func main() {
	demo1()
}
