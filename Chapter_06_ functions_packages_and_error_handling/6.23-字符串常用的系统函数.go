package main

import (
	"fmt"
	"strconv"
	"strings"
)

//1.统计字符串长度:Golang的编码统一为utf-8(ASCII的字符(字母和数字)占一个字节，汉字占三个字节)
func func1() {
	str := "你好,world!"
	fmt.Println("str len=", len(str))
}

//2.字符串遍历,同时处理有中文的问题 r:[]rune(str)
func func2() {
	str := "hello 北京"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符%v=%c\n", i, r[i])
	}
}

//3.字符串转整数:n,err:=strconv.Atoi("12")
func func3() {
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}
}

//4.整数转字符串 str=strconv.Itoa(12345)
func func4() {
	str := strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T", str, str)
}

//5.字符串转[] byte: var bytes=[]byte("hello go")
func func5() {
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
}

//6.[]byte转字符串:str =string([]byte{97,98,99})
func func6() {
	str := string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)
}

//7.10进制转2,8,16进制:str=strconv.FormatInt(123,2)	//2->8,16
func func7() {
	str1 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是:%v\n", str1)

	str2 := strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是:%v\n", str2)
}

//8.查询字串是否在指定字符串中:strings.Contains("seafood","foo")	//true
func func8() {
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)
}

//9.统计一个字符串有几个指定的子串:strings.Count("ceheese","e")	//4
func func9() {
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)
}

//10.不区分大小写的字符串比较(==是区分字母大小写的):fmt.Println(strings.EqualFold("abc","Abc"))	//true
func func10() {
	b := strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)
	fmt.Println("abc==Abc:", "abc" == "Abc")
}

//返回字串在字符串中第一次出现的index值,如果没有返回-1:strings.Index("NLT_abc","abc")	//4
func func11() {
	index := strings.Index("NLT_abcabcabc", "abc") //4
	fmt.Printf("index=%v\n", index)
}

//12.返回子串在字符串最后一次出现的index值,如果没有返回-1:strings.Index("NLT_abc","abc")	//
func func12() {
	index := strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v", index)
}

//13.将指定的字串替换成另一个字串:strings.Replace("go go hello","go","go语言",n)	//n可以指定你希望替换几个,n=-1表示全部替换
func func13() {
	str1 := "go go hello"
	str2 := strings.Replace(str1, "go", "北京", -1)
	fmt.Printf("str1=%v \nstr2=%v\n", str1, str2)
}

//14.按照指定的某个字符，为了分割标示，将一个字符拆分成字符串数组:strings.Split("hello,world,ok",",")
func func14() {
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)
}

//15.将字符串的字母进行大小写的转化:strings.ToLower("Go")	//go strings.ToUpper("Go")	//GO
func func15() {
	str := "Golang Hello"
	str1 := strings.ToLower(str)
	str2 := strings.ToUpper(str)
	fmt.Printf("str1=%v\nstr2=%v", str1, str2)
}

//16.将字符串左右两边的空格去掉:strings.TrimSpace("  tn a lone gopher ntrn    ")
func func16() {
	str := strings.TrimSpace("  tn a lone gopher ntrn    ")
	fmt.Printf("str=%q\n", str)
}

//17.将字符串左右两边指定的字符去掉:strings.Trim("!hello!","!")
func func17() {
	str := strings.Trim("!hello!", "!")
	fmt.Printf("str=%q\n", str)
}

//18.将字符串左边指定的字符去掉:strings.TrimLeft("!hello!","!")
func func18() {
	str := strings.TrimLeft("!hello!", "!")
	fmt.Printf("str=%q\n", str)
}

//19.将字符串右边指定的字符去掉:strings.TrimRight("!hello!","!")
func func19() {
	str := strings.TrimRight("!hello!", "!")
	fmt.Printf("str=%q\n", str)
}

//20.判断字符串是否以指定的字符串开头:strings.HasPrefix("ftp://192.168.10.1","ftp")
func func20() {
	b := strings.HasSuffix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v", b)
}

func main() {
	func20()
}
