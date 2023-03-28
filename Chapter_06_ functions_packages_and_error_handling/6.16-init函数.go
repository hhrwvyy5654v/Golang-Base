package main

import (
	"fmt"
)

/*基本介绍:每一个源文件都可以包含一个init函数，该函数会在main函数前执行，被Go运行框架调用，也就是说init会在main函数前调用
func init() {
	fmt.Println("init()...")
}
*/

//如果一个文件同时包含全局变量定义，init函数和main函数，则执行流程全局变量定义->int函数->main函数
var age = test()

//为了看到全局变量是先被初始化的，这里先写函数
func test() int {
	fmt.Println("test()")
	return 90
}

//init函数,通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println(`main()...age=`, age)
}

/*
如果main.go和utils.go都有变量定义,init函数执行流程是怎样的呢?

main.go							utils.go
import("utils")					变量定义[执行1]
变量定义[执行3]				 	 init函数[执行2]
init函数[执行4]
main函数[执行5]

*/
