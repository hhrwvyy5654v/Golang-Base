package main

import (
	"encoding/json"
	"fmt"
)

//json的饭徐磊话反序列化是指将json字符串反序列成对应的数据类型(比如结构体、map、切片)的操作

//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

//反序列化成struct
func func1() {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	//定义一个Monster实例
	var monster Monster

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v\n", monster, monster.Name)
}

//反序列化成map
func func2() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	//定义一个map
	var a map[string]interface{}
	//注意:反序列化map不需要make,因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v", a)
}

func func3() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	//定义一个slice
	var slice []map[string]interface{}
	//注意:反序列化不需要make,因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v", slice)
}

func main() {
	func3()
}
