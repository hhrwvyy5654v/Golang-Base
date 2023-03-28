package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
1.读取文件的内容并显示在终端(带缓冲区的方式),
使用os.Open,file.Close,bufio.NewReader(),reader.ReadString函数和方法
*/
func func1() {
	file, err := os.Open("./template.go")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//当函数退出时，要及时地关闭file
	defer file.Close() //及时关闭file句柄，否则会有内存泄漏

	/*创建*Reader,是带缓冲的
	const (
		defaultBufSize = 4096 //默认缓冲区为4096
	)
	*/
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Println(str)
		fmt.Println("文件读取结束...")
	}
}

/*
2.读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中),
这种方式适用于文件不大的情况。相关方法和函数(ioutil.ReadFile)
*/
func func2() {
	file := "./template.go"
	context, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(context))

}

func main() {
	func2()
}
