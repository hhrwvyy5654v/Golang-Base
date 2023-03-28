package main

import (
	"fmt"
)

/*1.基本语法
	var map 变量名 map[key type]value type
2.key可以是什么类型
	golang中的map的key可以是多种类型,
	比如bool、数字、string、指针、channel,
	还可以是前面几个类型的接口、结构体、数组
通常key为int、string
注意:slice,map还有function不可以,因为这几个没法用==来判断
3.value type可以是什么类型
	value type的类型和key基本一样,通常为：数字(整数、浮点数),string,map,struct
4.声明map是不会分配内存的,初始化需要make,分配后才能赋值和使用
5.map在使用前一定要make,key是不能重复的,value是可以相同的,key-value是无序的,make内置函数数目
*/

func main() {
	var a map[string]string
	a=make(map[string]string,10)
	a["001"]="宋江"
	a["002"]="吴用"
	a["003"]="武松"
	fmt.Println(a)
}
