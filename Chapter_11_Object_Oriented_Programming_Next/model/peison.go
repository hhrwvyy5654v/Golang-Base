package model

import (
	"fmt"
)

//不能随便查看一个人的年龄、工资等隐私
type person struct {
	Name   string
	age    int 
	salary float64
}

//写一个工厂模式的函数,相当于构造函数
func NewPerson(name string)*person{
	return &person{
		Name: name,
	}
}

//为了访问age和salary，编写一对SetXxx方法和GetXxx方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水范围错误")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}

func main() {
	fmt.Println(``)
}
