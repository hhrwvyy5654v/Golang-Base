package main

import (
	"fmt"
)

//1.string底层是一个byte数组，因此string可以镜像进行切片操作
func func1(){
	str:="hello@atguigu"
	slice:=str[6:]
	fmt.Println("slice:",slice)
}

//2.string是不可变的，也就说不能通过类似str[0]='z'的方式来修改字符串
//3.如果要修改字符串，可以先将string->[]type/或者[]rune->修改->重写转成string
func func3(){
	str:="hello@atguigu"
	/*
	转成byte后可以处理英文和数字，但是不能处理中文
	原因是[]byte是字节来处理，而一个汉字是3个字节，就会出现乱码
	解决方法：将string转成[]rune,因为[]rune是按照字符处理，兼容汉字
	*/
	arr1:=[]byte(str)
	arr1[0]='z'
	str=string(arr1)
	fmt.Println("str=",str)
	
	arr2:=[]rune(str)
	arr2[0]='北'
	str=string(arr2)
	fmt.Println("str=",str)

}



func main() {
	func3()
}
