package main

import (
	"fmt"
)

/*
(1)Golang中的方法是作用在指定的数据类型上的(即：和指定的数据类型绑定),因此自定义类型都可以有方法，而不仅仅是struct
(2)方法的声明和调用
type A struct{
	Num int
}
func(a A)test(){
	fmt.Println(a.Num)
}
说明:
	a.func(a A)test(){}表示A结构体有一方法,方法名为test
	b.(a A)体现test方法是和A类型绑定的
*/
type Person struct {
	Name string
}

//与Person类型绑定的方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func func1() {
	var p Person
	p.Name = "tom"
	p.test()
}

/*
对上面的总结:
(1)test方法和Person类型绑定
(2)test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调整
(3)func (p Person) test() ... p表示哪个Person变量调用,p是它的副本,这一点和函数传参非常相似
(4)p这个名字不是固定的，是由程序员指定的
*/

//a.给Person结构体添加speak方法，输出xxx是好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

//b.给Person结构体添加计算方法1，可以计算1+..+1000的结果，说明方法体内和函数一样可以进行各种运算
func (p Person) calculate1() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是", res)
}

//c.给Person结构体添加计算方法2，该方法可以接受一个数n,计算1+...+n的结果
func (p Person) calculate2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是", res)
}

//d.给Person结构体添加getSum方法，可以计算两个数的和并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

//方法a-d的调用
func func2() {
	var p Person
	p.Name = "tom"
	p.speak()
	p.calculate1()
	p.calculate2(57)
	fmt.Printf("res=%v", p.getSum(10, 308))
}

/*
方法的调用和传参机制原理
说明：
(1)方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时会将调用方法的变量当做实参也传递给方法；
如果变量是值类型则进行值拷贝,如果变量是引用类型则进行地址拷贝
(2)案例:请编写一个程序,要求如下。
	a.声明一个结构体Circle,字段为radius
	b.声明一个方法area和Circle绑定,可以返回面积
*/
type Circle struct {
	radius float64
}

func (c Circle) area1() float64 {
	return 3.14 * c.radius * c.radius
}

func func3() {
	var c Circle
	c.radius = 4.0
	res := c.area1()
	fmt.Println("面积是=", res)
}

/*
方法的声明(定义)
func (recevier type)methodName(参数列表)(返回值列表){
	方法体
	return 返回值
}
1.参数列表:表示方法的输入
2.recevier type:表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
3.recevier type:type可以是结构体，也可以其它的自定义类型
4.receiver:就是type类型的一个变量(实例),比如：Person结构体的一个变量(实例)
5.返回值列表：表示返回的值，可以多个
6.方法主体:表示为了实现某一功能代码块
7.return语句不是必须的

方法的注意事项和细节
(1)结构体类型是值类型，在方法调用中遵循值类型和传递机制，是值拷贝传递方式
(2)如果程序员希望在方法中修改结构体变量的值，可以通过结构体指针的方式来处理
*/
func (c *Circle) area2() float64 {
	//因为c是指针，因此标准访问字段的方式是(*c).radius
	return 3.14 * (*c).radius * (*c).radius
	//也等价于 return 3.14 * c.radius * c.radius
}

func func4() {
	var c Circle
	c.radius = 5.0
	res := (&c).area2()
	//编译器底层做了优化,(&c).area2()等价于c.area2(),编译器会自动给c加上&
	fmt.Println("面积是=", res)
}

//(3)Golang中的方法作用在指定的数据类型上(即和指定的数据类型绑定)，因此自定义类型都可以有方法，不仅仅是struct,比如int、发咯爱她等都可以有方法
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写方法改变i的值
func (i *integer) change() {
	*i = *i + 1
}

func func5() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)
}

/*
(4)方法的访问范围控制的规则和函数一样，方法名首字母小写只能在本包访问，方法名首字母大写可以在本包和其它包访问
(5)如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出
*/
type Student struct {
	Name string
	Age  int
}

//给*Student实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func func6() {
	//定义一个Student变量
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	//如果你实现了*Student类型的String方法就会自动调用
	fmt.Println(&stu)
}

//编写结构体MethodUtils,编写一个方法，该方法不需要参数，在方法中打印一个10*8的矩形
type MethodUtils struct {
}

func (mu MethodUtils) Print1() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func func7() {
	var mn MethodUtils
	mn.Print1()
}

//编写方法,提供m和n两个参数，方法中打印一个m*n的矩阵
func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func func8() {
	var mn MethodUtils
	mn.Print2(3, 8)
}

//编写方法计算矩形的面积(可以接受长和宽),将其作为方法返回值,在main方法中调用该方法，接受返回的面积值并打印。
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func func9() {
	var mn MethodUtils
	fmt.Println("矩形的面积:", mn.area(8, 2))
}

//编写方法判断一个数是奇数还是偶数
func (mu MethodUtils) JudgeNum(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func func10() {
	var mn MethodUtils
	fmt.Println(mn.JudgeNum(8))
}

//按照行、列、字符打印对应行数和列数的字符，比如(行：3，列：2，字符：8)
func (mu MethodUtils) Print3(n int, m int, key string) {
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fmt.Println(key)
		}
		fmt.Println()
	}
}

func func11() {
	var mn MethodUtils
	mn.Print3(2, 5, "*")
}

//定义计算机器结构体，实现加减乘除四个功能
type Calculate3 struct {
	Num1 float64
	Num2 float64
}

//实现形式1：分四个方法实现
func (calculate Calculate3) getSum() float64 {
	return calculate.Num1 + calculate.Num2
}
func (calculate Calculate3) getSub() float64 {
	return calculate.Num1 - calculate.Num2
}
func (calculate Calculate3) getPro() float64 {
	return calculate.Num1 * calculate.Num2
}
func (calculate Calculate3) getQuo() float64 {
	return calculate.Num1 / calculate.Num2
}

//实现形式2：用一个方法实现
func (calculate *Calculate3) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calculate.Num1 + calculate.Num2
	case '-':
		res = calculate.Num1 - calculate.Num2
	case '*':
		res = calculate.Num1 * calculate.Num2
	case '/':
		res = calculate.Num1 / calculate.Num2
	default:
		fmt.Println("运算符输入错误")
	}
	return res
}

/*
方法和函数的区别
(1)调用方式不一样
	函数的调用方式：函数名(实参列表)
	方法的调用方式：变量.方法名(实参列表)
(2)对于普通函数，接收者为值类型，不能将指针类型的数据直接传递，反之亦然
(3)对于方法(如struct方法),接收者为值类型时，可以直接用指针类型的变量调用方法，反过来也可以
*/
func (p Person) test01() {
	p.Name = "jack"
	fmt.Println("test01()=", p.Name)
}

func (p *Person) test02() {
	p.Name = "mary"
	fmt.Println("test02()=", p.Name)
}

func func12() {
	var p Person
	p.test01()
	//从形式上是传入地址，但本质依然是值拷贝
	(&p).test01()

	(&p).test02()
	//等价于(&p).test02(),从形式上是传入值类型，但本质依然是地址拷贝
	p.test02()
}

//总结：如果是和值类型,比如(p Person),则是值拷贝;如果是和指针类型,比如是(p *Person)则是地址拷贝

func main() {
	func12()
}
