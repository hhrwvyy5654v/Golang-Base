package main

import (
	"fmt"
)

func main() {
	//创建要给Student实例
	var student=utils.model.Student{
		Name:"tom",
		Score:89.8,
	}
	fmt.Println(student)
}
