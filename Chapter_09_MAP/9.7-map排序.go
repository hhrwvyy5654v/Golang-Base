package main

import (
	"fmt"
	"sort"
)

/*
基本介绍:
	(1)golang中没有一个专门的方法针对map的key进行排序
	(2)golang中的map默认是无序的，注意也不是按照添加的顺序存放的，每次遍历得到的输出可能不一样
	(3)golang中的map的排序是先将key进行排序，然后根据key值遍历输出
*/

func func1() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 67
	map1[8] = 90

	fmt.Println(map1)

	/*
		1.先按map的key放入到切片中
		2.对切片排序
		3.切片遍历，然后按照key来输出map的值
	*/
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}
}

func main() {
	func1()
}
