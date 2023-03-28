package main

import (
	"fmt"
)

/*
continue语句用于结束本次循环，继续执行下一次循环
continue语句出现在多层嵌套的循环语句中时，可以通过标签指明要跳过的是哪一层循环，这个和前面的break标签的使用规则一样
*/

func func1() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}
}

//打印1--100之间的奇数[要求使用for循环+continue]
func func2() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("奇数是", i)
	}
}

func func3() {
here:
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}

func func4() {
	var positive_number int = 0
	var negative_number int = 0
	var num int
	for {
		fmt.Println("清输入一个整数：")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positive_number++
		}
		negative_number++
	}
	fmt.Printf("正数和负数的个数是 %v %v", positive_number, negative_number)
}

func main() {
	func3()
}
