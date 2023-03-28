package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn succeed=", conn)
	//功能一:客户端可以发送单行数据，然后就退出
	reader := bufio.NewReadWriter(os.Stdin) //os.Stdin代表标准输入[终端]
	//从终端读取一行用户输入,并准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err=", err)
	}
	if line == "exit" {
		fmt.Println("客户端退出...")
	}
	//再将line发送给服务器
	n, err := conn.Write([]byte(line + "\n"))
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}
	fmt.Printf("客户端发送了%d字节的数据并退出", n)
}
