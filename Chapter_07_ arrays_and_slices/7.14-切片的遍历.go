package main

import (
	"fmt"
)

//for 循环常规方式遍历
func func1() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}
}

//for-range 结构遍历切片
func func2() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}
}

func main() {
	func2()
}
