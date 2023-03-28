package main

import (
	"fmt"
)

/*
1.如果return是在普通的函数,则表示跳出该函数,即不再执行函数中return后面代码,也可以理解为终止函数
2.如果return是在main函数,表示终止main函数,也就是终止程序
*/

func main() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			return
		}
		fmt.Println("哇哇", i)
	}
	fmt.Println("Hello World!")
}
