package main

import (
	"fmt"
)

//for循环遍历
func func1() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
	}
}

//for-range遍历循环
func func2() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i, v1 := range arr {
		for j, v2 := range v1 {
			fmt.Printf("arr[%v][%v]=%v \t", i, j, v2)
		}
	}
}

func main() {
	func2()
}
