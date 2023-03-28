package main

import (
	"fmt"
)

/*
16.8.1 channel的关闭
	使用内置函数close可以关闭channel,当channel关闭后,就不能向channel写数据,但是依然从该channel读取数据
*/

func test01() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //关闭管道
	fmt.Println("ok~ok~")
	n1 := <-intChan
	fmt.Println("n1", n1)
}

/*
16.8.2 channel的遍历

	channel支持for--range的方式进行遍历,请注意两个细节
	(1)在遍历时,如果channel没有关闭,则会出现deadlock的错误
	(2)在遍历时,如果channel已经关闭,则会正常遍历数据,遍历完后,就会退出遍历

16.8.3 channel遍历和关闭的案例演示
*/
func test02() {
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	/*
		遍历管道不能使用普通的for循环
		for i:=0;i<len(intChan2);i++{

		}
	*/
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}

func main() {
	test02()
}
