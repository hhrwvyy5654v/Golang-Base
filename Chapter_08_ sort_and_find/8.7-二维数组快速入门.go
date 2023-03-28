package main

import (
	"fmt"
)

func func1() {
	var arr [6][4]int
	var num int = 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			arr[i][j] = num
			num++
		}
	}
	fmt.Println("num:", arr)
}

func main() {
	func1()
}
