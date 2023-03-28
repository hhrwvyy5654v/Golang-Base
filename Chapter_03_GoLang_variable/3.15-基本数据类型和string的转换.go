package main

import (
	"fmt"
	"strconv"
)

//*基本类型转string类型
//方法一：fmt.Sprintf("%参数",表达式)
//Sprintf根据format参数生成格式化的字符串并返回改字符串
func func1() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)
}

//方法二：使用strconv包的函数
func func2() {
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	var str string

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	/*
		strconv.FormatFloat(num4,'f',10,64)
		'f'格式10:表示小数位保留10位
		64:表示这个小数是float64
	*/
	strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv包中有一个函数Itoa
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)
}

//string类型转基本数据类型
func func3() {
	var str1 string = "true"
	var b bool
	/*
	b,_=strconv.ParseBool(str1)
	说明:strconv.ParseBool(str)函数会返回两个值(value bool,err error)
	*/
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T b=%v\n", b, b)

	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str1, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str2 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}

//string转基本数据类型的注意事项:要确保String类型能够转换成有效的数据
func func4(){
	var str4 string="hello"
	var n3 int64=11
	n3,_=strconv.ParseInt(str4,10,64)
	fmt.Printf("n3 type %T n3=%v\n", n3, n3)
}



func main() {
	func3()
}