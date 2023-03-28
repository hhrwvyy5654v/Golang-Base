package main

import (
	"fmt"
	"sync"
	"time"
)

/*
16.6 channel(管道)--看个需求

需求：现在要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来，要求使用goroutine完成。
分析思路：
(1)使用goroutine来完成，效率高，但是会出现并发/并行安全问题
(2)这里提出了不同goroutine如何通信的问题

思路:
(1)编写一个函数，来计算各个数的阶乘，并放入map中
(2)启动多个协程，将统计的结果放入到map中
(3)map应该做出一个全局

16.6.1 不同goroutine之间如何通讯
(1)全局变量的互斥锁
(2)使用管道channel来解决

16.6.2 使用全局变量加锁同步改进程序
因为没有对全局变量m加锁,因此会出现资源争夺问题，代码会出现错误，提示concurrent map writes
解决方案：加入互斥锁
计算的数阶乘很大，结果会越界，可以将阶乘改成sum+=unit64(i)
*/
var (
	myMap = make(map[int]int, 10)
	/*
		声明一个全局的互斥锁
		lock是一个全局的互斥锁
		sync是包：synchornized同步
		Mutex：是互斥
	*/
	lock sync.Mutex
)

// test函数就是计算n!,让将这个结果放入到myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将res放入到myMap
	lock.Lock() //加锁
	myMap[n] = res
	lock.Unlock() //解锁
}

func test01() {
	//这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	//休眠10秒钟
	time.Sleep(time.Second * 10)

	//输出结果
	lock.Lock() //加锁
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock() //解锁
}

/*
16.6.3 为什么需要channel
(1)使用全局变量加锁来同步解决goroutine的通讯并不完美
(2)主线程在等待所有goroutine全部完成的时间很难确定
(3)如果主线程休眠时间长了,会加长等待时间;如果等待时间短了，可能还有goroutine处于工作状态，这时候=也会随着主线程的退出而销毁
(4)通过全局变量来加锁实现通讯并不利于多个协程对全局变量的读写操作

16.6.4 channel的基本介绍
(1)channel本质就是一个数据结构--队列
(2)数据是先进先出【FIFO：first in first out】
(3)线程安全，多goroutine访问时不需要加锁，就是说channel本身就是线程安全
(4)channel有类型的,一个string的channel只能存放string类型数据

16.6.5 定义/声明channel
举例:
var inChan chan int(intChan用于存放int数据)
var mapChan chan map[int]string (mapChan用于存放map[string]string类型)
var perChan chan Person
var perChan2 chan *Person
说明:
channel是引用类型
channel必须初始化才能写入数据，即make后才能使用
管道是有类型的,intChan只能写入整数int

16.6.6 管道的初始化,写入数据到管道，从管道读取数据的基本的注意事项
*/
func test02() {
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 10)
	//2.看看intChan是什么
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//intChan<-98	//注意：当我们给管道写入数据时，不能超过其容量
	//4.看看管道的长度和容量
	fmt.Printf("channel len=%v capacity=%v\n", len(intChan), cap(intChan))
	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v capacity=%v\n", len(intChan), cap(intChan))
	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取出就会报告deadlock
	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
}

/*
16.6.7 channel使用的注意事项
(1)channel中只能存放指定的数据类型
(2)channel的数据放满后就不能再放入了
(3)如果从channel取出数据后可以继续放入
(4)在没有使用协程的情况下如果channel数据取完了再取就会报dead lock
*/

// 创建一个intChan,最多可以存放3个int,演示存3个数据到intChan，然后取出三个int
func test03() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 10
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Printf("num1=%v num2=%v num3=%v", num1, num2, num3)
}

// 创建一个mapChan,最多可以存放10个map[string]string的key-val，演示写入和读取
func test04() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 50)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"

	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"

	mapChan <- m1
	mapChan <- m2
}

type Cat struct {
	Name string
	Age  int
}

// 创建一个catChan最多可以存放10个Cat结构体变量,演示写入和读取的用法
func test04() {
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "tom~", Age: 18}
	catChan <- &cat1
	catChan <- &cat2

	cat11 := <-catChan
	cat22 := <-catChan
	fmt.Println(cat11, cat22)
}

// 创建一个allChan，最多可以存放10个任意数据类型变量,演示写入和读取的用法
func test05() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "tom~", Age: 180}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"
	//取出
	cat11 := <-allChan
	cat22 := <-allChan
	v1 := <-allChan
	v2 := <-allChan
	fmt.Println(cat11, cat22, v1, v2)
}

func main() {
	test05()
}
