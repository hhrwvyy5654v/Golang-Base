package main

import (
	"Basic/Chapter_11_Object_Oriented_Programming_Next/model"
	"fmt"
)

func func1() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSalary(5000)
	fmt.Println(p)
	fmt.Println(p.Name, "age=", p.GetAge(), "salary=", p.GetSalary())
}

func func2() {
	a := model.NewAccount("hrq123456", 789689, "123456")
	if a != nil {
		fmt.Println("创建成功", a)
	} else {
		fmt.Println("创建失败")
	}
}

func main() {
	func2()
}
