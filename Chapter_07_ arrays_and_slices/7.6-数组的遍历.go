package main

import (
	"fmt"
)

//1.常规遍历

//2.for-range结构遍历
func func2() {
	heroes := [...]string{"宋江", "吴用", "卢俊义"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
		fmt.Println()
	}

	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n", v)
	}
}

func main() {
	func2()
}
