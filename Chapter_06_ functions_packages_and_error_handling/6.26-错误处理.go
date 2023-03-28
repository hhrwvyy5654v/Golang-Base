package main

import (
	"fmt"
	"errors"
)

/*
(1)在默认情况下，当发生错误(panic),程序就会退出
(2)如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。
还可以在捕获到错误后，给管理员一个提示(邮件、短信)
(3)Go语言追求简洁优雅，所以Go语言不支持传统的try…catch…finally这种语法
(4)Go中引入的处理方式：defer、panic、recover
(5)几个异常的使用场景可以这么简单描述:Go中可以抛出一个panic异常,然后在defer中通过recover捕获这个异常，然后正常处理
*/

//错误代码
func func1() {
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//使用defer+recover来捕获和处理异常
func func2() {
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {  //说明捕获到异常
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

/*
错误处理的好处:程序不轻易挂掉，如果加入预警代码，就可以让程序更加健壮
自定义错误类型:
	(1)Go中使用errors.New和panic内置函数自定义错误类型
	(2)errors.New("错误说明"),会返回一个error类型的值，表示一个错误
	(3)panic内置函数,接受一个interface{}类型的值(也就是任何值了)作为参数,可以接受error类型的变量,输出错误信息，并退出程序
*/
func func3(name string)(err error){
	if name=="config.ini"{
		//读取…
		return nil
	}else{
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func func4(){
	err:=func3("config2.ini")
	if err!=nil{
		//如果读取文件发送错误，就输出这个错误，并终止错误
		panic(err)
	}
	fmt.Println("fun3()继续执行")
}

func main() {
	//测试
	func4()
	fmt.Println("main()下面的代码")
}
