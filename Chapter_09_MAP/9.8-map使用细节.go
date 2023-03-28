package main

import (
	"fmt"
)

//(1)map是引用类型，遵守引用类型传递的机制，在一个函数接受map,修改后就会直接修改原来的map
func modify(map1 map[int]int) {
	map1[10] = 900
}

//(2)map的容量达到后，再想增加map元素会自动扩容，并不会发生panic,也就是map能动态的增长键值对(key-value)
func func1_2() {
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[10] = 98
	map1[20] = 1
	map1[30] = 578978
	modify(map1)
	fmt.Println(map1)
}

/*
(3)map的value也经常使用struct类型，更适合管理复杂的数据(前面value是一个map更好),比如value为Student结构体
	map的key为学生的学号，是唯一的
	map的value为结构体，包含学生的姓名、年龄、地址
*/

//定义学生结构体
type stu struct {
	Name    string
	Age     int
	Address string
}

func func2() {
	students := make(map[string]stu, 10)
	stu1 := stu{"tom", 18, "北京"}
	stu2 := stu{"marry", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	fmt.Println(students)

	//遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.Name)
		fmt.Printf("学生的年龄是%v \n", v.Age)
		fmt.Printf("学生的地址是%v \n", v.Address)
		fmt.Println()
	}
}

func main() {
	func2()
}
