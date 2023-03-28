package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	customerView.go【界面】【含customerService字段】
	1.可以显示界面
	2.可以接受用户的输入
	3.根据用户的输入，调用customerService的方法完成客户的管理【修改,删除,显示等等】

	customerService【处理业务逻辑】
	1.完成对用户的各种窜哦做操作
	2.对客户的增删改查
	3.会声明一个customer的切片
	
	customer【表示数据】model层
	1.表示一个客户
	2.客户各种字段
	`)
}
