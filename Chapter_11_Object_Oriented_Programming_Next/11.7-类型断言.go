package main

import (
	"fmt"
)

// 11.7.1 由一个具体的需要引出类型断言
type Point struct {
	x int
	y int
}

func test01() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	//如何将a赋值给一个Point变量
	var b Point
	/*
		类型断言,表示判断a是否指向Point类型的变量,
		如果是就转成Point类型并赋给b变量,否则就报错
	*/
	b = a.(Point)
	fmt.Println(b)
}

/*
11.7.2 基本介绍

	类型断言,由于接口是一般类型,不知道具体类型。如果要转成具体类型就需要使用类型断言,具体如下：
*/
func test02() {
	//类型断言的其它案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口可以接受任何类型
	y := x.(float32)
	fmt.Printf("y的类型是%T,值是%v", y, y)
}

/*
对上面代码的说明:
	在进行类型断言时，如果类型不匹配就会报panic,因此进行类型断言时要确保原来的空接口指向的就是断言的类型
	如果在进行断言时,要带上检测机制，如果成功就ok,否则也不要报panic
*/

// 类型断言(带检测的)
func test03() {
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口可以接受任意类型

	y, ok := x.(float64)
	if ok == true {
		fmt.Println("convert success")
		fmt.Printf("y的类型时%T,值是%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
}

/*
11.7.3 类型断言的最佳实践1
给Phone结构体增加一个特有的方法call(),当Usb接口接受的是Phone变量时,还需要调用call方法。
*/
//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}
func (p Phone) Call() {
	fmt.Println("手机在打电话...")
}

type Camera struct {
	name string
}

// 让Camera实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果usb是指向Phone结构体变量则还需要调用Call方法
	//类型断言
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
	usb.Stop()
}

func test04() {
	//定义一个Usb接口数组可以存放Phone和Camera的结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Phone{"尼康"}

	/*
		遍历usbArr,Phone还有一个特有的方法call(),请遍历Usb数组，如果是Phone变量。
		除了调用Usb接口声明的方法外,还需要调用Phone特有方法
	*/
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
	fmt.Println(usbArr)
}

/*
11.7.4 类型断言的最佳实践2
写一函数，循环判读传入参数的类型
*/
type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型,值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型,值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型,值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型,值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型,值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是类型不确定,值是%v\n", index, x)
		}
	}
}

func test05() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}

func main() {
	test05()
}
