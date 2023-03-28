package main

import (
	"fmt"
	"time"
)

// (1)channel可以声明为只读，或者只写性质
func test01() {
	//1.在默认情况下,管道是双向
	//var chan1 chan int

	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	fmt.Println("chan2=", chan2)

	//3.声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	fmt.Println("num2", num2)
}

// (2)只读和只写的最佳实践案例
func test02() {
	var ch chan int
	ch = make(chan int, 10)
	exitchan := make(chan struct{}, 2)
	go send(ch, exitchan)
	go recv9(ch, exitchan)

	var total = 0
	for _ = range exitchan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束...")
}

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

func recv(ch <-chan int, exitchan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitchan <- a
}

// (3)使用select可以解决从管道取数据的阻塞问题
func test03() {
	//使用select可以解决从管道取数据的阻塞问题
	//1.定义一个管道10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	/*
		传统的方法在遍历管道时,如果不关闭会阻塞而导致deadlock
		问题：在实际开发中,可能我们不好确定什么关闭该管道,可以使用select方式解决
	*/
	//label
	for {
		select {
		//注意:这里如果intChan一直没有关闭,不会一直阻塞而deadlock,会自动到下一个case匹配
		case v := <-intChan:
			fmt.Println("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Println("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都取不到了,不玩了,程序员可以加入逻辑")
			time.Sleep(time.Second)
			return
			//break label
		}
	}
}

/*
(4)goroutine中使用recover，解决协程中出现panic,导致程序崩溃问题
说明:如果我们起了一个协程,但是这个协程出现了panic,
如果我们没有捕获这个panic,就会造成整个程序崩溃,这时我们可以在goroutine中使用
recover来捕获进行处理,这样即使这个协程发生的问题,但是主线程不受影响,可以继续执行
*/
//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

// 函数
func test() {
	//这里我们可以使用defer+recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func test04() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}

func main() {
	test01()
}
