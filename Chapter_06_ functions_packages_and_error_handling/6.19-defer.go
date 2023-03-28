package main

import (
	"fmt"
)

//为什么需要defer:在函数中，程序员经常需要创建资源(比如:数据库连接、文件句柄、锁等)，为了在函数执行完毕后及时地释放资源，Go的设计者提供defer(延时机制)
func sum1(n1 int, n2 int) int {
	/*
		当执行到defer时，暂不执行，会将defer后面的语句压入到独立的栈(defer栈)
		当函数执行完毕时，再从defer栈按照先入后出的方式出栈执行
	*/
	defer fmt.Println("ok1 n1", n1)
	defer fmt.Println("ok2 n2", n2)

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

//在defer将语句放入栈时，会将相关的值拷贝同时入栈
func sum2(n1 int, n2 int) int {
	/*
		当执行到defer时，暂不执行，会将defer后面的语句压入到独立的栈(defer栈)
		当函数执行完毕时，再从defer栈按照先入后出的方式出栈执行
	*/
	defer fmt.Println("ok1 n1", n1) //n1=10
	defer fmt.Println("ok2 n2", n2) //n2=20

	n1++ //n1=11
	n2++ //n2=21

	res := n1 + n2 //res=32
	fmt.Println("ok3 res=", res)
	return res
}

/*
defer最主要的价值是当函数执行完毕后可以及时地释放函数创建的资源
func test(){
	//关闭文件资源
	file=openfile(文件名)
	defer file.close()
	//其它代码
}

func test(){
	//释放数据库资源
	connect=openDatabase()
	defer connect.close()
}

(1)在Golang编程中通常的做法是创建资源后，比如(打开了文件，获取了数据库的连接，或者是锁资源),可以立即执行defer file.Close() defer connect.Close()
(2)在defe后可以继续使用创建的资源
(3)当函数执行完毕后,系统会依次congdefer栈中取出语句，关闭资源
(4)这种机制非常简洁，程序员不需要在什么时机关闭资源而烦心

*/

func main() {
	res1 := sum1(10, 20)
	fmt.Println("res1=", res1)
	fmt.Println("")

	res2 := sum2(10, 20)
	fmt.Println("res2=", res2)
}
