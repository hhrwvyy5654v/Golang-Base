package main

import (
	"fmt"
)

//map遍历:使用for-range结构遍历
func func1() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	for key, value := range cities {
		fmt.Printf("key=%v value=%v\n", key, value)
	}
}

//map的长度
func func2() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	fmt.Println("map的长度:", len(cities))
}

func main() {
	func2()
}
