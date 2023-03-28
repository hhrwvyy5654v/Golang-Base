package main

import (
	"flag"
	"fmt"
	"os"
)

//1.os.Args是一个string切片，用来存储所有的命令行参数
func func1() {
	fmt.Println("命令行的参数有:", len(os.Args))
	//遍历os.Args切片就可以得到所有命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v", i, v)
	}
}

//2.使用flag包可以方便地解析命令行参数，，而且顺序可以随意
func func2() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")

	flag.Parse()

	//输出结果
	fmt.Printf("use=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}

func main() {
	func2()
}
