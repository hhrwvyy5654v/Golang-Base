package main

import (
	"fmt"
	"reflect"
)

/*
1.reflect.Value.Kind获取变量的类别,返回的是一个常量
2.Type和kind的区别
Type是类型,kind是类别,Type和Kind可能是相同的,也可能不同的
比如:var num int =10	num的Type是int,Kind也是int
比如:var stu Student	stu的Type是"包名".Student,Kind是struct
3.通过反射可以让变量在interface{}和Reflect.Value之间就、互相转换
4.使用反射的方式来获取变量的值(并返回对应的类型),要求数据类型匹配,
比如x是int,那么就应该使用reflect.Value(x).Int(),而不能使用其它的,否则报panic
5.通过反射来修改变量,注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样才能改变传入的变量的值，同时需要使用到reflect.Value.Elem()的方法
6.如何理解reflect.Value.Elem()
	fn.Elem()用于获取指针指向变量,类似于
	var num=10
	var b *int=&num
	*b=3


*/

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type=%T\n", val)
	fmt.Printf("val kind:%v\n", val.Kind())
	//Elem返回V持有的接口保管的值的value封装，或者v持有的指针指向值的value封装
	val.Elem().SetInt(110)
	fmt.Printf("val=%v\n", val)
}

func func1() {
	var num int = 20
	testInt(&num)
	fmt.Println("num=", num)
}

func main() {
	func1()
}
