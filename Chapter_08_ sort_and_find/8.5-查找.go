package main

import (
	"fmt"
)

/*
在Golang中常用的查找：
	(1)顺序查找
	(2)二分查找
*/

//顺序查找
func Sequential_lookup() {
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Printf("请输入要查找的人名:")
	fmt.Scanln(&heroName)

	//方式一
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到%v,下标%v\n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v \n", heroName)
		}
	}

	//方式二
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v,下标%v \n", heroName, index)
	} else {
		fmt.Println("没有找到", heroName)
	}
}

//二分查找
func Binary_lookup(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		Binary_lookup(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		Binary_lookup(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了,下标为%v\n", middle)
	}
}

func main() {
	//Sequential_lookup()
	arr := [6]int{1, 2, 7, 13, 33,89}
	Binary_lookup(&arr, 0, len(arr)-1, 2)
}
