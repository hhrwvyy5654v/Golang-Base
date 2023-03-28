package main

import (
	"fmt"
)

/*
map增加和更新
map["key"]=value
如果key还没有就是增加,key存在就是修改
*/
func func1() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	cities["no3"] = "上海~"
	fmt.Println(cities)
}

/*
map删除
delete(map,"key"),delete是一个内置函数。
如果key存在，就删除key-value;如果key不存在,不操作但是也不会报错
如果我们要删除map所有的key,没有一个专门的方法可以一次删除，可以遍历一次删除，逐个删除或者map=make(...),make一个新的，让原来的成为垃圾，被gc回收
*/
func func2() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	delete(cities, "no2")
	fmt.Println(cities)
}

//map查找
func func3() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值:%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}
}



func main() {
	func3()
}
