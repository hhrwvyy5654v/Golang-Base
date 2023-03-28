package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	1.反射是程序在运行时检查其变量和值并找到它们类型的能力
	2.Go语言自带的reflect包实现了在运行时进行反射的功能，这个包可以帮助识别一个interface{}类型变量其底层的具体类型和值。我们的createQuery函数接收到一个interface{}类型的实参后，需要根据这个实参的底层类型和值去创建并返回INSERT语句，这正是反射包的作用所在。
	`)
}
