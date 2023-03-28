package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	(1)在调用一个函数时,会给该函数分配一个新的空间,编译器会通过自身的处理让这个新的空间和其他栈的空间区分开来
	(2)在每个函数对应的栈中,数据空间是独立的,不会混淆
	(3)当一个函数调用完毕(执行完毕)后,程序会销毁这个函数对应的栈空间
	
	(4)如果一个函数返回多个值,在接受时希望忽略某个返回值,则使用_符号表示占位符忽略,比如：
		_,res3=getsumAndSub(3,9)
	(5)如果返回值只有一个,(返回值类型列表)可以不写
	`)
}
