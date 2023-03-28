package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
1.拷贝文件
说明：将一张图片/电影/mp3拷贝到另一个文件夹
func Copy(dst Writer,src Reader)(written int64,err error)
注意:Copy函数是io包提供的
*/
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcFile,获取到Reader
	reader := bufio.NewReader(srcFile)
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//通过dstFile获取Writer
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}

func func1() {
	srcFile := ""
	dstFile := ""
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}

//2.统计一个文件中含有的英文、数字、空格及其它字符数量
/*
	思路:打开一个文件，创建一个Reader,每读取一行就去统计该行有多少个英文、
	数字、空格和其它字符,然后将结果保存到一个结构体

定义一个结构体，用于保存统计结果
*/
type CharCount struct {
	ChCount    int //英文个数
	NumCount   int //数字个数
	SpaceCount int //空格个数
	OtherCount int //其它字符个数
}

func func2() {

	fileName := ""
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file=%v", err)
		return
	}
	defer file.Close()

	//创建一个CharCount实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环读取fileName的内容
	for {
		str0, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//为了兼容中文字符，将str转成[]rune
		str := []rune(str0)
		//遍历str，进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	//输出统计结果
	fmt.Printf("字符个数:%v\n数字个数:%v\n空格个数:%v,其它字符个数:%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}

func main() {
	fmt.Println(``)
}
