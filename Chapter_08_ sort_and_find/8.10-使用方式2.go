package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	使用方式2:直接初始化
	声明:var 数组名 [大小][大小]类型=[大小][大小]类型{{初值...},{初值...}}
	赋值(有默认值,比如int类型就是0)

	二维数组在声明/定义时也对应着四种写法[和一维数类似]
	var 数组名 [大小][大小]类型=[大小][大小]类型{{初值...},{初值...}}
	var 数组名 [大小][大小]类型=[...][大小]类型{{初值...},{初值...}}
	var 数组名 = [大小][大小]类型{{初值...},{初值...}}
	var 数组名 =[...][大小]类型{{初值...},{初值...}}

	`)
}
