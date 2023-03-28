package main

import (
	"fmt"
)

/*
11.6.1 基本介绍
	变量(实例)具有多种形态,在Go语言中,多态是通过接口来实现的。可以按照统一的接口来调用不同的实现。

11.6.3 接口体现多态的两种形式
多态参数
	在前面的Usb案例中,Usb usb,既可以接受手机变量，又可以接收相机变量,就体现了Usb接口多态。
多态数组
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

type Camera struct {
	name string
}

// 让Camera实现Usb接口方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

func main() {
	/*
		定义一个Usb接口数组,可以存放Phone和Camera的结构体变量
		这里就体现多态数组
	*/
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)
}
