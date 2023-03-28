package main

import (
	"fmt"
)

//方式1:定义一个切片,然后让切片去引用一个已经创建好的数组
func func1() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	var slice = arr[1:3]
	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)
	fmt.Println("slice len=", len(slice))
	fmt.Println("slice cap=", cap(slice))
}

/*
方式2:通过make来创建切片
基本语法: var 切片名 []type=make([]type,len,[cap])
参数说明:
	type 数据类型
	len 大小
	cap 切片容量(可选参数,分配了cap则要保证cap>=len)
*/
func func2() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice的size:", len(slice))
	fmt.Println("slice的cap:", cap(slice))
}

/*
代码小结:
	(1)通过make方式创建切片可以指定切片的大小和容量
	(2)如果没有给切片的各个元素赋值，那么就会使用默认值[int,float=>0	string=>""	bool=>false]
	(3)通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素

方式1和方式2的区别:
	方式1是直接引用数组，这个数组是事先存在的 ，程序员是可见的
	方式2是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的

*/

//方式3:定义一个切片，直接就指定具体数组，使用原理类似于make的方式
func func3() {
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice))
	fmt.Println("strSlice cap=", cap(strSlice))
}




func main() {
	func3()
}
