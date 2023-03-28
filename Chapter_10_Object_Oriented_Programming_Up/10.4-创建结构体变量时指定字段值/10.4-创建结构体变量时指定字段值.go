package main

import (
	"fmt"
)

/*
Golang在创建结构体实例(变量)时可以直接指定字段的值
*/

type Student struct{
	Name string
	Age int
}

//方式1:在创建结构体变量时就直接指定字段的值
var student1=Student{"小明",19}
var student2=Student{"小红",20}

//在创建结构体变量时，把字段名和字段值写在一起可以不依赖字段的定义顺序
var student3=Student{
	Name: "jack",
	Age: 20,
}
var student4=Student{
	Age: 25,
	Name: "tom",
}


//方式2:返回结构体的指针类型
var student5 *Student=&Student{"小王",29}
var student6=&Student{"小张",39}

//在创建结构体指针变量时，把字段名和字段值写在一起可以不依赖字段的定义顺序
var student7=&Student{
	Name: "susan",
	Age: 26,
}
var student8=&Student{
	Age: 25,
	Name: "candy",
}


func main() {
	fmt.Println(student1,student2,student3,student4,*student5,*student6,*student7,*student8)
}
