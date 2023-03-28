package main

import (
	"fmt"
	"os"
)

func func1() {
	file, err := os.Open("./template.go")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//输出文件
	fmt.Printf("file=%v", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}

func main() {
	func1()
}
