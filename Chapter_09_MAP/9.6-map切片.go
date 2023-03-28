package main

import (
	"fmt"
)

/*
基本介绍:
	切片的数据类型如果是map,我们称之为slice of map,map切片。这样使用则map个数可以动态变化。

要求：
	使用一个map来记录monster的信息name和age,也就是monster对应一个map，并且妖怪的个数可以动态地增加=>map切片
*/
func func1() {
	//1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	//2.增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	/*
		以下的写法越界：
		if monsters[2]==nil{
			monsters[2]=make(map[string]string,2)
			monsters[2]["name"]="狐狸精"
			monsters[2]["age"]="300"
		}
	*/

	//使用切片的append函数，可以动态地增加monster
	//1.定义monster信息
	newMonster := map[string]string{
		"name": "火源邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}

func main() {
	func1()
}
