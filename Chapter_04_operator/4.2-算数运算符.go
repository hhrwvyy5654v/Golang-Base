package main

import (
	"fmt"
)



func func1() {
	//如果运算的数都是整数，那么相除之后去掉小数部分，保留整数部分
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	//如果我们希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	//++和--的使用
	var i int = 10
	i++ //等价于i=i+1
	fmt.Println("i=", i)
	i-- //等价于i=i-1
	fmt.Println("i=", i)

	//取模的使用:a%b=a-(a/b*b)
	fmt.Println("\n10%3=", 10%3)
	fmt.Println("110%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)

}


//注意事项：
func func2(){
	/*
	*Golang的自增自减只能当作一个独立语言来使用,以下是错误用法
	var i int=8
	var a int
	a=i++
	a=i--

	if i++>0{
		fmt.Println("Ok")
	}
	*/

	/*
	*Golang的++和--只能写在变量的后面，不能写在变量的前面，即只有a++、a--,没有++a、--a
	*/

}


func main() {
	func1()
}
