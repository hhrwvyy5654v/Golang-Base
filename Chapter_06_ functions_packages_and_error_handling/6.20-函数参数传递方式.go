package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	两种传递方式:
		(1)值传递
		(2)引用传递
	不管是值传递还是引用传递,传递给函数的都是变量的副本。不同的是,值传递的是值的拷贝,引用传递的是地址的拷贝。
	一般来说,地址拷贝效率高,因为数据量小,而值拷贝决定拷贝的数据大小,数据越大,效率越低。

	值类型和引用类型
	(1)值类型:基本数据类型int系列,float系列,bool,string,数组和结构体struct
	(2)引用类型:指针、slice切片,map,管道chan,interface等都是引用类型

	值传递和引用传递使用特点
		(1)值类型默认是值传递：变量直接存储值，内存通常在栈中分配
		(2)引用类型默认是引用传递:变量存储的是一个地址，这个地址对应的空间才真正存储数据(值),内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收
		(3)如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&,函数内以指针的方式操作变量。
	`)
}
