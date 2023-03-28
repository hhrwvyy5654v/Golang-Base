package main

import (
	"fmt"
	"reflect"
)

//编写一个案例演示对(结构体类型、interface{}、reflect.Value)进行反射的基本操作
func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type、kind值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//下面将rVal转成interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}


//专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type、kind值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//下面将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)

	//将interface{}通过断言转成需要的类型，简单使用了一带检测的类型断言
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//1.先定义一个int
	var num int = 100
	reflectTest01(num)
	fmt.Println("")
	//2.定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
