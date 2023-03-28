package main

import (
	"fmt"
)

/*
数组的定义
var 数组名 [数组大小]数据类型
var a[5]int
赋初值 a[0]=1 a[1]=30 ……

数组在内存中的布局:
(1)数组地址可以通过数组名来获取 &intArr
(2)数组的第一个地址就是数组的首地址
(3)数组的各个元素的地址间隔是依据数组的类型决定，比如int64->8 int32->4……
*/

func main() {
	var intArr [3]int
	fmt.Println(intArr)
	intArr[0]=10
	intArr[1]=20
	intArr[2]=30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址:%p\nintArr[0]的地址:%p\nintArr[1]的地址:%p\nintArr[2]的地址:%p",&intArr,&intArr[0],&intArr[1],&intArr[2])
}
