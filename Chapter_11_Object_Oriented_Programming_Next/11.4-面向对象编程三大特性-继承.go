package main

import (
	"fmt"
)

/*
// 小学生
type Pupil struct {
	Name  string
	Age   int
	Score int
}

// 显示他的成绩
func (p *Pupil) ShowInfo() {
	fmt.Printf("学生名=%v	年龄=%v	成绩=%v\n", p.Name, p.Age, p.Score)
}

func (p *Pupil) SetScore(score int) {
	//业务判断
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

// 大学生
type Graduate struct {
	Name  string
	Age   int
	Score int
}

// 显示他的成绩
func (g *Graduate) ShowInfo() {
	fmt.Printf("学生名=%v	年龄=%v	成绩=%v\n", p.Name, p.Age, p.Score)
}

func (g *Graduate) SetScore(score int) {
	//业务判断
	g.Score = score
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

// 代码冗余...高中生...
func test01() {
	//测试
	var pupil = &Pupil{
		Name: "tom",
		Age:  10,
	}
	pupil.testing()
	pupil.SetScore(90)
	pupil.ShowInfo()

	//测试
	var graduate = &Graduate{
		Name: "marry",
		Age:  20,
	}
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}
*/

/*
小结：
(1)Pupil和Graduate两个结构体的字段和方法几乎完全一样，但是我们却写了相同的代码，代码复用性不强
(2)出现代码冗余，而且不利于维护，同时也不利于功能的拓展
(3)解决方法：通过继承的方式来解决


11.4.2 继承基本介绍和示意图
继承可以解决代码复用，让我们的编程更加靠近人类思维
当多个结构体存在相同的属性(字段)和方法时,可以从这些结构体中抽象出结构体,在该结构体中定义这些相同的属性和方法


11.4.3 嵌套匿名结构体的基本语法
type Goods struct{
	Name string
	Price int
}
type Book struct{
	Goods	//嵌套匿名结构体Goods
	Write string
}
*/

/*
// 编写一个学生考试系统
type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Pupil和Graduate共有的方法也绑定到*Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v	年龄=%v	  成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}

// 小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

// 这时Pupil结构体特有的方法保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

// 大学生
type Graduate struct {
	Student //嵌入了Student匿名结构体
}

// 这时Graduate结构体特有的方法
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func func1() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "tom~"
	graduate.Student.Age = 23
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
}

func main() {
	func1()
}
*/

/*
11.4.5 继承给编程带来的便利
(1)代码的复用性提高了
(2)代码的拓展性和维护性提高了

11.4.6 继承的深入讨论
1)结构体可以使用匿名嵌套结构体所有的字段和方法，即:首字母大写或者小写的字段、方法都可以使用
*/

/*
type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

func func2() {
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOk()
	b.A.hello()
}
*/

/*
2)匿名结构体字段访问可以简化

func func3() {
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOk()
	b.A.hello()

	//如上写法可以简化
	b.Name="smith"
	b.age=20
	b.SayOk()
	b.hello()
}
*/
/*
对上面的代码小结
(1)当我们直接通过b访问字段或方法时，其执行流程如下：比如b.Name
(2)编译器会先看b中嵌入的类型有没有Name,如果有则直接调用B类型的Name字段
(3)如果没有就去看B中嵌入的匿名结构体A有没有声明Name字段,如果有就调用,如果没有继续查找,找不到就报错

3)当结构体和匿名结构有相同的字段或者方法时，编译器采用就近原则访问,如希望访问匿名结构体的字段和方法,可以通过匿名结构体来区分

func func3() {
	var b B
	b.Name = "tom"	//就近原则会访问结构体的name字段
	//b.A.Name	就明确指定访问匿名结构体的字段name
	b.A.Name="jack"

	b.Age=78
	b.say()	//就近原则访问B结构体的say函数
	b.Hello()
	//b.A.Hello()	//就明确指定访问A匿名结构体的方法Hello()
	b.A.hello()
}
*/

/*
4)结构体嵌入两个(或多个)匿名结构体,如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法),
在访问时就必须明确指定匿名结构体名字，否则编译报错。
*/
/*
type A struct {
	Name string
	age  int
}
type B struct {
	Name  string
	Score float64
}
type C struct {
	A
	B
}

func func4() {
	var c C
	//如果c没有Name字段,而A和B有Name就必须通过指定匿名结构体名字来区分
	c.A.Name = "tom" //error
	fmt.Println(c)
}
*/

/*
6)嵌套匿名结构体后也可以在创建结构体变量时,直接指定各个匿名结构体字段的值

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func func5() {
	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时,直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 50000.9}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Price: 50000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

	tv3 := TV2{&Goods{"电视机003", 7000.99}, &Brand{"创维", "河南"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视机004",
			Price: 9000.09,
		},
		&Brand{
			Name:    "长虹",
			Address: "四川",
		},
	}
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
}
*/

/*
11.4.7 课堂练习
结构体的匿名字段是基本数据类型,如何访问

type Monster struct{
	Name string
	Age int
}

type E struct{
	Monster
	int
	n int
}

func func6(){
	var e E
	e.Name="狐狸精"
	e.Age=300
	e.int=20
	e.n=40
	fmt.Println("e=",e)
}
*/

/*
说明:
(1)如果一个结构体有int类型的匿名字段就不能有第二个
(2)如果需要有多个int的字段,则必须给int字段指定名字

11.4.8 面向对象编程-多重继承
多重继承说明

	如果一个struct嵌套了多个匿名结构体,那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承

细节说明

	(1)如嵌入的结构体有相同的字段名或者方法名字，则在访问时需要通过匿名结构体类型来区分
	(2)未来保证代码的简洁性，尽量不使用多重继承

func func7() {
	type Goods struct {
		Name string
		Age  float64
	}
	type Brand struct {
		Name string
		Goods
	}
	type TV struct {
		Goods
		Brand
	}
}
*/
func main() {
	func7()
}
