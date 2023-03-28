package main

import (
	"fmt"
)

func main() {
	var i int = 5
	//二进制输出
	fmt.Printf("%b \n", i)
	//八进制：0-7,满8进一，数字以0开头表示
	var j int = 011
	fmt.Println("j=", j)
	//0-9以及A-F，满16进1，以0x或0X开头表示
	var k int = 0x11
	fmt.Println("k=", k)
}
