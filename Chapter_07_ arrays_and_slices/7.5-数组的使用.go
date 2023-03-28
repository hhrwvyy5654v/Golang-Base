package main

import (
	"fmt"
)

//四种初始化数组的方式
func func1() {
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1=", arr1)

	var arr2 = [3]int{4, 5, 6}
	fmt.Println("arr2=", arr2)

	//[...]是规定的写法
	var arr3=[...]int{8,9,10}
	fmt.Println("arr3=", arr3)

	var arr4=[...]int{1:800,0:3738,2:789}
	fmt.Println("arr4=", arr4)
}

func main() {
	func1()
}
