package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
11.5.1 基本介绍
	Golang中多态特性主要是通过接口来体现的
11.5.3 接口快速入门
*/

// 声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作……")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作……")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作……")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作……")
}

type Computer struct {
}

/*
编写一个方法Working,接收一个Usb接口类型变量。
只要是实现了Usb接口(所谓实现Usb接口就是实现了Usb接口声明所有方法)
*/
func (c Computer) Working(usb Usb) { //Usb变量会根据传入的实参来判断到底是Phone还是Camera
	//通过usb接口变量来调用Start()和Stop()方法
	usb.Start()
	usb.Stop()
}

func test01() {
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}

/*
11.5.4 接口概念的再说明

	interface类型可以定义一组方法,但是这些不需要实现。并且interface不能包含任何变量。

到某个自定义类型(比如结构体Phone)要使用的时候,再根据具体情况把这些方法写出来。

11.5.5 基本语法
小结说明:

	(1)接口里的所有方法都没有方法体,即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想。
	(2)Golang中的接口,不需要显式的实现。只要一个变量,含有接口类型的所有方法,这个变量就实现这个接口。因此Golang中没有implement这样的关键字。

11.5.6 接口使用的应用场景

11.5.7 注意事项和细节
(1)接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量
*/
type AInterface interface {
	Say()
}
type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

func test02() {
	var stu Stu //结构体变量,实现了Say(),实现了AInterface
	var a AInterface = stu
	a.Say()
}

/*
(2)接口中所有的方法都没有方法体,即都是没有实现的方法
(3)在Golang中,一个自定义类型需要将某个接口的所有方法都实现,我们说这个自定义类型实现了该接口
(4)一个自定义类型只有实现了某个接口才能将该自定义的实例(变量)赋给接口类型
(5)只要是自定义数据类型就可以实现接口,不仅仅是结构类型
*/
type integer int

func (i integer) Say() {
	fmt.Println("integer Say() i=", i)
}
func test03() {
	var i integer = 10
	var b AInterface = i
	b.Say() //输出integer Say i = 10
}

/*
(6)一个自定义类型可以实现多个接口
*/
type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

// Monster实现了AInterface和BInterface
func test04() {
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}

/*
(7)Golang接口中不能有任何变量
type AInterface interface{
	Name string
	test01()
	test01()
}
*/

/*
(8)一个接口(比如A接口)可以继承多个别的接口(比如B,C接口),这时如果要实现A接口也必须将B,C接口的方法也全部实现
*/
type CInterface interface {
	test01()
}

type DInterface interface {
	test02()
}

// 接口EInterface继承了CInterface和DInterface
type EInterface interface {
	CInterface
	DInterface
	test03()
}

// 如果需要实现EInterface就需要将CInterface和DInterface都实现
type Stu struct {
}

func (stu Stu) test01() {
}
func (stu Stu) test02() {
}
func (stu Stu) test03() {
}

func test05() {
	var stu Stu
	var e EInterface = stu
	e.test01()
}

/*
(9)interface类型默认是一个指针(引用类型),如果没有对interface初始化就使用会输出nil
(10)空接口interface{}没有任何方法,所以所有类型都实现了空接口,即我们可以把任何一个变量赋给空接口
*/
type T interface {
}

func test06() {
	var t T = stu //ok
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}

/*
11.5.9 接口编程的最佳实践
实现对Hero结构体切片的排序：sort.Sort(data Interface)
*/
//1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3.实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

/*
Less方法就是决定使用你使用什么标准进行排序
1.按Hero的年龄从小打到大排序
*/
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//修改成对Name的排序
	//return hs[i].Name<hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换
	//temp:=hs[i]
	//hs[i]=hs[j]
	//下面一句话等价于三句话
	hs[i], hs[j] = hs[j], hs[i]
}

// 1.声明Student结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

// 将Student的切片,按照Score从大到小排序
func test07() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对intSlice切片进行排序
	/*
		1.冒泡排序
		2.也可以使用系统提供的方法
	*/
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//测试是否可以对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprint("英雄|%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将hero append到heroes切片
		heroes = append(heroes, hero)
	}

	//看看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("--------排序后--------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}

/*
11.5.10 实现接口 vs 继承
*/
//Money结构体
type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

// LittleMoney结构体
type LittleMonkey struct {
	Monkey //继承
}

// 让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习,会飞翔...")
}

// 让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习,会游泳...")
}

func test08() {
	//创建一个LittleMonkey
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}

/*
小结:
(1)对A结构体继承了B结构体,那么A结构就自动地继承了B结构体的字段和方法，并且可以直接使用
(2)当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此可以认为接口时对继承机制的补充

接口和继承解决的问题不同:
继承的价值主要在于:解决代码的复用性和可维护性
接口的价值主要在于:设计,设计好各种规范(方法),让其它自定义类型去实现这些方法

接口比继承更加灵活 Person Student	BirdAble LittleMonkey
接口比继承更加灵活,继承是满足is-a的关系,而接口只是满足like-a的关系
接口在一定程度上实现代码解耦
*/

func main() {
	test08()
}
