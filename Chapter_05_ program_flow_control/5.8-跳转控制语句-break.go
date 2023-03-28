package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
break语句用于终止某个语句块的执行，用于中断for循环或者跳出switch语句
break默认会跳出最近的for循环
break后面可以指定标签，跳出标签对应的for循环
*/

//随机生成1-100的一个数，当生成99时，看看你一共用了多少次
func func1() {
	var count int = 0
	for {
		//time.Now().Unix():返回一个从1970:01:01的0时0分0秒到现在的秒数
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}
}

//break语句出现在多层嵌套的语句块中，可以通过标签指明要终止的是哪一层语句块
func func2() {
	//lable1:
	for i := 0; i < 4; i++ {
	lable2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable2 //break默认会跳出最近的for循环
			}
			fmt.Println("j=", j)
		}
	}
}

func main() {
	func2()
}
