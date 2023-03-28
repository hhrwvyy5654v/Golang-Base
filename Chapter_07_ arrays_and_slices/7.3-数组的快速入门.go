package main

import (
	"fmt"
)

//使用数组解决问题，程序的可维护性增加,方法代码更加清晰，也容易扩展

func main() {
	//1.定义一个数组
	var hens [7]float64
	//2.给数组的每一个元素赋值
	hens[0] = 3.0
	hens[1] = 4.0
	hens[2] = 5.0
	hens[3] = 6.0
	hens[4] = 7.0
	hens[5] = 8.0
	hens[6] = 9.0
	//3.遍历数组求出总体重
	totalweight := 0.0
	for i := 0; i < len(hens); i++ {
		totalweight += hens[i]
	}
	//4.求出平均体重
	avgweight := fmt.Sprintf("%.2f", totalweight/float64(len(hens)))
	fmt.Println("avgweight=", avgweight)
}
