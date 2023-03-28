package main

import (
	"fmt"
)

func func1() {
	//基本数据类型在内存布局
	var i int = 100
	//i的地址是什么,&i
	fmt.Println("i的地址=", &i)

	/*
	ptr是一个指针变量
	ptr的类型是*int
	ptr本身的值是&i
	 */
	var ptr *int = &i
	fmt.Printf("\n指针变量ptr=%v\n", ptr)
	fmt.Printf("指针变量ptr的地址=%v\n", &ptr)
	fmt.Printf("指针变量ptr指向的值=%v\n", *ptr)
}

func main(){
	func1()
}