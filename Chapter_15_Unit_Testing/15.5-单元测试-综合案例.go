package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	1.编写一个Monster结构体,字段Name,Age,Skill
	2.给Monster绑定方法Store,可以将一个Monster变量(对象),序列化后保存到文件中
	3.给Monster绑定方法ReStore,可以将一个反序列化的Monster,从文件中读取并反序列化为Monster对象,检查Monster对象,检查反序列化,名字正确
	4.编程测试用例文件store_test.go,编写测试用例函数TestStore和TestRestore进行测试。
	`)
}
