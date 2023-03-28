package main

import (
	"fmt"
)

//使用fmt.Scanln()获取
func func1() {
	var name string
	var age byte
	var salary float32
	var whether_pass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&salary)

	fmt.Println("请输入是否通过考试:")
	fmt.Scanln(&whether_pass)

	fmt.Printf("名字是 %v\n 年龄是 %v\n 薪水是 %v \n 是否通过考试 %v \n", name, age, salary, whether_pass)
}

//使用fmt.Scanf()获取
func func2() {
	var name string
	var age byte
	var salary float32
	var whether_pass bool

	fmt.Println("请输入你的姓名、年龄、薪水、是否通过考试,使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &whether_pass)
	fmt.Printf("名字是 %v\n年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, salary, whether_pass)
}

func main() {
	func2()
}
