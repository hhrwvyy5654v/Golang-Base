package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	(1)slice的确是一个引用类型
	(2)slice从底层来说是一个数据结构(struct结构体)
	type slice struct{
		ptr *[2]int	//引用数组的地址
		len int	//长度
		cap int	//容量
	}
	
	`)
}
