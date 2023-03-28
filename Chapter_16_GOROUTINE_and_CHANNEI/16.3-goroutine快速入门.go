package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
请编写一个程序，完成如下功能：
(1) 在主线程(可以理解成进程)中，开启一个goroutine,该协程每隔1秒输出"hello,world"
(2) 在主线程中也每隔一秒输出"hello,golang"，输出10次后，退出程序
(3) 要求主线程和goroutine同时执行.
*/
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func func1() {
	go test()	//开启了一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello,golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	func1()
}

/*
小结:
(1)主线程是一个物理线程，直接作用在CPU上，是重量级的，非常耗费CPU资源
(2)协程是从主线程开始的，是轻量级的线程，是逻辑态,对资源消耗相对小
(3)Golang的协程机制的重要特点是可以轻松地开启上万个协程。其它编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这样就凸显了Golang在并发上的优势。
*/