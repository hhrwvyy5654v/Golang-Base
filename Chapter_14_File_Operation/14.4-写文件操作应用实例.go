package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
1.基本介绍-os.OpenFile函数
func OpenFile (name string, flag int, perm FileMode) (file *File, err error)
说明：os.OpenFile是一个更一般性的文件打开函数,它会使用指定的选项(如O_RDONLY等)、指定的模式(如0666等)打开指定名称的文件。
如果操作成功，返回的文件对象可用于I/O。如果错误，错误底层类型是*PathError。
*/

//1.创建新文件，写入内容5句"hello world !"
func func1() {
	//打开文件
	filepath := "./temp.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入
	str := "hello world !\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*
		因为writer是带缓存，所以在调用writerString方法时,
		内容是先写到缓存的，所以需要调用Flush方法，
		将缓冲的数据真正地写入到文件中，否则文件中会没有数据
	*/
	writer.Flush()
}

//2.打开一个已经存在的文件，将原来的内容覆盖成新的内容为10句"你好,尚硅谷!"
func func2() {
	//打开已经存在的文件
	filepath := "./temp.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入
	str := "你好,尚硅谷!\r\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	/*
		因为writer是带缓存，所以在调用writerString方法时,
		内容是先写到缓存的，所以需要调用Flush方法，
		将缓冲的数据真正地写入到文件中，否则文件中会没有数据
	*/
	writer.Flush()
}

//3.打开一个存在的文件，在原来的内容后追加内容'ABC! ENGLISH!'
func func3() {
	//打开已经存在的文件
	filepath := "./temp.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入
	str := "ABC!ENGLISH!\r\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	/*
		因为writer是带缓存，所以在调用writerString方法时,
		内容是先写到缓存的，所以需要调用Flush方法，
		将缓冲的数据真正地写入到文件中，否则文件中会没有数据
	*/
	writer.Flush()
}

//4.打开一个存在的文件,将原来的内容读出显示在终端，并且追加5句"hello,北京!"
func func4() {
	//打开已经存在的文件
	filepath := "./temp.txt"
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()

	//先读取原来文件内容，并显示在终端上
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //如果读取到文件末尾
			break
		}
		//显示到终端
		fmt.Print(str)
	}

	//准备写入
	str := "hello,北京!\r\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*
		因为writer是带缓存，所以在调用writerString方法时,
		内容是先写到缓存的，所以需要调用Flush方法，
		将缓冲的数据真正地写入到文件中，否则文件中会没有数据
	*/
	writer.Flush()
}

/*
5.编写程序，将一个文件的内容写入到另外一个文件。注意：这两个文件已经存在
说明:使用ioutil.ReadFile/ioutil.WriteFile
*/
func func5() {
	filepath1 := "./temp.txt"
	filepath2 := "./template.go"
	data, err := ioutil.ReadFile(filepath2)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n", err)
		return
	}
	err = ioutil.WriteFile(filepath1, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}

/*
6.判断文件是否存在
golang判断文件或文件夹是否存在的方法为使用os.STat()函数返回的错误值进行判断:
(1)如果返回的错误为nil,说明文件或文件夹存在
(2)如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
(3)如果返回的错误为其它类型，则不确定是否存在
*/
func func6(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	fmt.Print(func6("./temp.txt"))
}
