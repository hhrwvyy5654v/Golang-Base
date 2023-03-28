package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
1.基本介绍
(1)反射可以在运行时动态获取变量的各种信息，比如变量的类型（type），类别（kind）
(2)如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
(3)通过反射，可以修改变量的值，可以调用关联的方法。
(4)使用反射，需要 import("reflect")

2.反射的应用场景
(1)不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口,这种需要对函数或方法反射。
func bridge(funcPtr interface{},args ...interface{})
这种桥接模式的第一个参数funcPtr以接口的形式传入函数指针，
函数参数args以可变参数的形式传入，
bridge函数中可以用反射来动态执行funcPtr函数
(2)对结构体序列化时，如果结构体有指定Tag,也会使用到反射生成对应的字符串

3.反射的重要函数和概念
(1)reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型
(2)reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型reflect.value是一个结构体类型。通过reflect.Value可以获取关于该变量更多信息。
(3)变量、interface{}和reflect.Value是可以相互转换的，这点在开发中会经常使用到。
	`)
}
