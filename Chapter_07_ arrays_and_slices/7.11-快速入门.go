package main

import (
	"fmt"
)

func func1() {
	var intArr [5]int = [...]int{1, 22, 33, 44, 6}
	slice := intArr[1:3] //表示slice引用到intArr这个数组下标为[1,3)的元素
	fmt.Println("intArr=", intArr)
	fmt.Println("slice的元素是:", slice)
	fmt.Println("slice的元素个数:", len(slice))
	fmt.Println("slice的容量:", cap(slice))	//切片的容量可以动态变化

	fmt.Println("intArr[1]的地址:",&intArr[1])
	fmt.Println("slice[0]的地址:",&slice[0])
}

func main() {
	func1()
}
