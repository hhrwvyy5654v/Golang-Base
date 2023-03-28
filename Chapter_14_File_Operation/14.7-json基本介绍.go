package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	JSON(JavaScript Object Notation)是一种轻量级的数据交换格式,易于阅读和编写,同时也易于机器解析和生成

	目前JSON已经是主流的数据格式

	JSON易于机器解析和生成,并有效地提升网络传输效率,通常程序在网络传输时会先将数据(结构体、map等)序列化成
	JSON字符串,到接收方得到json字符串时再反序列化恢复成原来的数据类型(结构体、map等)。
	`)
}
