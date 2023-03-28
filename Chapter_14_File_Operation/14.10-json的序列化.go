package main

import (
	"encoding/json"
	"fmt"
)

//JSON的序列化是指将有key-value结构体的数据类型(比如结构体、map、切片)序列化成json字符串的操作

//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

/*
对于结构体的序列化，如果我们希望序列化后key的名字可以自定义，可以给struct指定一个tag标签
*/
type Monster_ struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Salary   float64 `json:"monster_salary"`
	Skill    string  `json:"monster_skill"`
}

//将结构体序列化
func func1() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2022-02-08",
		Salary:   5789,
		Skill:    "牛魔拳",
	}

	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

//将map进行序列化
func func2() {
	//定义一个map
	var a map[string]interface{}
	//使用map需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map序列化后=%v\n", string(data))
}

//对切片进行序列化
func func3() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	//使用map前需要make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前需要make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	//将map进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice序列化后=%v\n", string(data))
}

//对基本数据类型序列化，但是意义不大
func func4() {
	var num1 float64 = 2345.6
	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1序列化后=%v\n", string(data))
}




func main() {
	func3()
}

