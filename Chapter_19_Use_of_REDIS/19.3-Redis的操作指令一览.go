package main

import (
	"fmt"
)

/*
说明:Redis安装好后,默认有16个数据库,初始默认使用0号库,编号是0~15
1.添加key-val[set]
2.查看当前的redis的所有key	[keys *]
3.获取key对应的值	[get key]
4.切换redis数据库	[select index]
5.如何查看当前数据库的key-val数量	[dbsize]
6.清空当前数据库的key-val和清空当前所有数据库的key-val[flushdb flushall]
*/

func main() {
	fmt.Println(`
	Redis的命令一览
	1.http://redisdoc.com
	`)
}
