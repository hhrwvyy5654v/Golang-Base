package main

import (
	"fmt"
)

//有两个变量a和b,要求将其进行交换，但是不允许使用中间变量,最终打印结果
func interview() {
	var a int = 10
	var b int = 20

	a = a + b
	b = a - b //b=a+b-b=a
	a = a - b //a=a+b-b=b
	fmt.Printf("a=%v b=%v", a, b)
}

func demo1() {
	//逻辑运算符的特点
	//(1)运算顺序从右到左
	var a int = 3
	var c int
	c = a + 3
	fmt.Println("c =", c)
	//(2)赋值运算符的左边只能是变量,右边可以是变量、表达式、常量
	var d int
	d = a
	a = 8 + 2*8
	fmt.Println("d =", d)
}

func main() {
	interview()
}
