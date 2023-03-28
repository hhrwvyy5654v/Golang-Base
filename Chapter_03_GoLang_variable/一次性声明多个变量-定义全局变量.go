package main

import "fmt"

//定义全局变量
var n1=100
var n2=200
var name="jack"

//一次性声明全局变量
var(
	n3=300
	n4=900
	name1="marry"
)


func main() {
	//一次性声明多个变量方式1
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)

	//一次性声明多个变量方式2
	var n4, name, n5 = 1000, "Tom", 888
	fmt.Println("n4=",n4, "name=",name, "n5=",n5)

	//一次性使用多个变量方式3:使用类型推导
	n6,time,n7:=100,"2022/05/30",990
	fmt.Println("n1=",n6,"time=",time,"n7=",n7)
}
