package main

import (
	"fmt"
)

/*
	(1)函数的形参列表可以是多个,返回值列表也可以是多个
	(2)形参列表和返回值列表的数据类型可以是值类型和引用类型
	(3)函数的命名遵循标识符命名规范,首字母不能是数字,
	(4)首字母大写该函数可以被本包文件和其它包文件使用,类似与public;首字母小写,只能被本包文件使用,其它包文件不能使用,类似于private
	(5)函数中的变量是局部的,函数外不生效
	(6)基本数据类型和数组默认都是值传递,即进行值拷贝。在函数内修改,不会影响到原来的值
*/

//(7)如果希望函数内的变量能够修改函数外的变量,可以传入变量的地址&,函数内以指针的方式操作变量
func func7(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("func7() num=", *n1)
}

//(8)Go函数不支持传统的函数重载,会报函数重复定义
//(9)在Go中,函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量,通过变量可以对函数调用
func func9(n1 int, n2 int) int {
	return n1 + n2
}

//(10)函数既然是一种数据类型,因此在Go中,函数可以作为形参,并且调用
func fun10(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

/*
(11)为了简化数据类型定义,GO支持自定义数据类型
		基本语法：type 自定义数据类型名 数据类型	//相当于别名
		案例1：type myInt int 	//这时myInt就等价于int来使用
		案例2：type mySum func(int,int)int		//这时mySum就等价于一个函数类型func(int,int) int
*/
func func11() {
	type myInt int
	var num1 myInt
	var num2 myInt
	num1 = 40
	num2 = int(num1) //这里显然需要显示转换,Go认为myInt和int两个类型
	fmt.Println("num1=", num1, "num2=", num2)
}

//(12)支持对函数返回值命名
func func12(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

//(13)使用_标识符，忽略返回值
/*
(14)Go支持可变参数
	//支持0到多个参数
	func sum(args…int)sum int{

	}
	//支持1到多个参数
	func sum(n1 int,args…int)sum int{

	}
	说明：
		a.args是slice切片，通过args[index]可以访问到各个值
		b.如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后
*/
func func13(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[i]表示取出args切片的第一个元素值，其它依次类推
	}
	return sum
}

func main() {
	/*
		num := 20
		func7(&num)
		fmt.Println("main() num=", num)

		a := func9
		fmt.Printf("\na的类型%T,func8()类型是%T\n", a, func9)
		res := a(10, 40) //类似于res:=func8(10,40)
		fmt.Println("res=", res)

		res2 := fun10(func9, 50, 60)
		fmt.Println("\nres=", res2)

		func11()


		a1,b1:=func12(1,2)
		fmt.Printf("a=%v b=%v\n",a1,b1)
	*/
	res4 := func13(10, 0, -1, 90, 100)
	fmt.Println("res=", res4)
}
