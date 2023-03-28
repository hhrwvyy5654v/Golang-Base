package main

import (
	"fmt"
)

/*
说明：编写一个函数fbn( n int) ,要求完成
1) 可以接收一个n int
2) 能够将斐波那契的数列放到切片中
3) 提示，斐波那契的数列形式：
arr[O] = l;arr[l] = l;arr[2]=2; arr[3] = 3; arr[4]=5; arr[5]=8
*/
func func1(n int) []uint64 {
	//声明一个切片，大小为n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	fbnSlice := func1(20)
	fmt.Println("fbnSlice=", fbnSlice)
}
