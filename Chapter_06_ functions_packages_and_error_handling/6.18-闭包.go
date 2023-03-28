package main

import (
	"fmt"
	"strings"
)

//基本介绍:闭包就是一个函数和与其相关的引用环境组成的一个整体(实体)
func AddUpper() func(int) int {
	var n int = 10
	var str = "Hello "
	return func(x int) int {
		n = n + x
		str += "world "
		fmt.Println("str=", str)
		return n
	}
}

/*
(1)AddUpper是一个函数,返回数据类型是fun(int)int
(2)闭包返回的是一个匿名函数，但是这个匿名函数引用到函数外的n,因此这个匿名函数就和n形成了一个整体，构成闭包。
(3)可以这样理解：闭包是类，函数是操作，n是字段。函数和它使用到n构成闭包。
(4)当我们反复地调用f函数时，因为n是初始化一次，因此每调用一次就进行累计
*/

/*
闭包的最佳实践:
	(1)编写一个函数makeSuffix(suffix string)可以接受一个文件名后缀名(比如.jpg),并返回一个闭包
	(2)调用闭包，可以传入一个文件名，如果该文件没有指定的后缀(比如.jpg)，则返回文件名.jpg,如果已经有.jpg后缀，则返回原文件名。
	(3)要求使用闭包的方式完成
	(4)string.HasSuffix,该函数可以判断某个字符串是否指定的后缀
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name没有指定后缀则加上，否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/*
代码说明:
	返回的匿名函数和makeSuffix(suffix string)的suffix变量和组合合成一个闭包，因为返回的函数引用到suffix这个变量
	如果使用传统方法，也可以实现这个功能，但是需要每次都传入后缀名；而闭包可以保留上次引用的某个值，所以传入一次就可以反复使用
	*/

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))   //winter.jpg
	fmt.Println("文件名处理后=", f2("bird.jpg")) //bird.jpg
}
