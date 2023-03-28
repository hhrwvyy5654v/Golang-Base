package main

import (
	"fmt"
	"unsafe"
)

func main(){

	//该区域的数据值可以在同一个类型范围内不断变化
	var i int =10
	i=30
	i=50
	fmt.Println("i=",i)
	//不能改变i的数据类型,比如i=1.2

	//变量在同一个作用域(在一个函数或者代码块)内不能重名,比如
	/*var i int =59
	i:=99
	*/

	//变量=变量名+值+数据类型
	var n1=100
	fmt.Printf("n1的类型 %T\n",n1)

	var n2 int64 =10
	fmt.Printf("n2的类型 %T\n  n2占用的字节数 %d \n",n1,unsafe.Sizeof(n2))
}