package main

import (
	"fmt"
)

//1.创建一个byte类型的26个元素的数组，分别放置'A'-'Z'。使用for循环访问所有元素并打印出来。
func func1() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i) //注意将i->byte
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}
}

//2.求出一个数组的最大值并得到对应的下标
func func2() {
	var intArr [6]int = [...]int{1, -2, 9, 84, 1000, -4}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v", maxVal, maxValIndex)
}

//3.求出一个数组和的平均值
func func3(){
	var intArr [5]int =[...]int{1,-2,3,4,76}
	sum:=0
	for _,val:=range intArr{
		sum+=val
	}
	//如何让平均值保留到小数
	fmt.Printf("sum=%v 平均值=%v",sum,float64(sum)/float64(len(intArr)))
}


func main() {
	func3()
}
