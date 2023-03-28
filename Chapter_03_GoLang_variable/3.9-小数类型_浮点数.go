package main

import (
	"fmt"
)

//演示golang中小数类型使用
func main() {
	//浮点数=符号位+指数位+位数位
	var price float32 = 89.12
	fmt.Println("price=", price)
	var num1 float32 = -0.0089
	var num2 float32 = -7809656.09
	fmt.Println("num1=", num1, "\nnum2=", num2)
	//尾数部分可能丢失，造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "\nnum3=", num4)
	//golang的浮点类型默认为float64类型
	var num5=1.1
	fmt.Printf("num5的数据类型是%T\n",num5)
	//浮点数常量有两种表示形式
	num6:=5.12
	num7:=.123
	fmt.Println("num6=", num6, "\nnum7=", num7)
	num8:=5.1234e2	//5.1234*10^2
	num9:=5.1234E2	//5.1234*10^2
	num10:=5.1234E-2	//5.1234/10^2
	fmt.Println("num8=", num8, "\nnum9=", num9,"\nnum10=", num10)

}
