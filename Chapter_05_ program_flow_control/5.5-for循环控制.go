package main

import (
	"fmt"
)

//for循环的三种使用方式
func func1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("你好,尚硅谷", i)
	}
}

func func2() {
	j := 1        //循环变量初始化
	for j <= 10 { //循环条件
		fmt.Println("你好,尚硅谷", j)
		j++ //循环变量迭代
	}
}

func func3() {
	k := 1
	for {
		if k <= 10 {
			fmt.Println("你好,尚硅谷", k)
		} else {
			break //跳出循环
		}
	}
}

//*Golang提供for-range方式，可以方便遍历字符串和数组
//传统方式
func func4() {
	var str string = "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i]) //使用到下标
	}
}

//for-range
func func5() {
	str := "abc-ok"
	for index, val := range str {
		fmt.Printf("index=%d val=%c \n", index, val)
	}
}

//[]rune切片
func func6() {
	var str string = "hello world! 北京"
	str2 := []rune(str) //把str转换成[]rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}
}

//for-range遍历方式
func func7() {
	str := "abc-ok 上海"
	for index, val := range str {
		fmt.Printf("index=%d val=%c \n", index, val)
	}
}

func main() {
	func7()
}
