package main

import (
	"encoding/json"
	"fmt"
)

/*
1.Golang语言面向对象编程说明
	(1)Golang也支持面向对象编程(OOP),但是和传统的面向对象编程有区别，并不是存粹的面向对象语言，所以我们说Golang支持面向对象编程特性是比较准确的。
	(2)Golang没有类(class),Go语言的结构体(struct)和其它编程语言的类(class)有同等的地位,可以理解Golang是基于struct来实现OOP特性的。
	(3)Golang面向对象编程非常简介,去掉了传统OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等
	(4)Golang仍然有面向对象编程的继承、封装和多态的特性，只是实现的方式和其它OOP语言不一样，比如继承：Golang没有extends关键字，继承是通过匿名字段来实现
	(5)Golang面向对象(OOP)很优雅，OOP本身就是语言类型系统的一部分，通过接口(interface)关联，耦合性低，也非常优雅。也就是说在Golang中面向编程是非常重要的特性。
2.如何声明结构体
基本语法
	type 结构体名称 struct{
		field1 type	//字段
		field2 type	//字段
	}
3.字段/属性
基本介绍
	(1)从概念上讲:结构体字段=属性=field
	(2)字段是结构体的一个组成部分，一般是基本数据类型、数组，也可以是引用类型
4.注意事项和细节说明
	(1)在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值),
	布尔类型是false,数值是0,字符是"";数组类型的默认值和它的元素类型相关,
	比如score[3]int则为[0,0,0];指针,slice,map的零值都是nil,即还没有分配空间
	(2)不同结构体变量的字段是独立,互不影响,一个结构体变量字段的更改不影响另外一个，结构体是值类型

*/

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

//创建一个cat变量
func func1() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃"

	fmt.Println("cat1:", cat1)
	fmt.Println("name:", cat1.Name)
	fmt.Println("age:", cat1.Age)
	fmt.Println("color:", cat1.Color)
	fmt.Println("hobby:", cat1.Hobby)
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

func func2() {
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用slice一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	//使用map一定要make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"

	fmt.Println(p1)
}

//创建结构体变量和访问结构体字段
func func3() {
	//方式1:直接声明
	var person1 Person
	fmt.Println(person1)

	//方式2:{}
	var person2 Person = Person{}
	fmt.Println(person2)

	/*
		(1)方式3和方式4返回的是结构体指针
		(2)结构体指针访问字段的标准方式应该是：(*结构体指针).字段名,比如(*person).Name="tom"
		(3)go做了一个简化，使得也支持结构体指针.字段名,比如person.Name="tom",是因为go编译器底层对person.Name做了转化(*person).Name
	*/
	//方式3:&
	var person3 *Person = new(Person)
	(*person3).Name = "smith"
	person3.Age = 24
	fmt.Println(person3)

	//方式4:{}
	var person4 *Person = &Person{}
	(*person4).Name = "scott"
	person4.Age = 45
	fmt.Println(person4)
}

//struct类型的内存分配机制
func func4() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"

	var p2 *Person = &p1
	p2.Name = "tom"
	fmt.Printf("p1.Name=%v p2.Name=%v \n", p1.Name, p2.Name)
	fmt.Printf("p1.Name=%v p2.Name=%v \n", p1.Name, (*p2).Name)

	fmt.Println("\np1的地址:", &p1)
	fmt.Println("p2的地址:", &p2)
	fmt.Println("\np1的值:", p1)
	fmt.Println("p2的值:", p2)
}

/*
结构体注意事项
(1)结构体的所有字段在内存中是连续的
*/
type Point struct {
	x int
	y int
}

type Rect1 struct {
	left, right, down, up Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func func5() {
	r1 := Rect1{Point{1, 2}, Point{3, 4}, Point{5, 6}, Point{7, 8}}
	fmt.Printf("r1.left.x的地址:%p \nr1.left.y的地址:%p \nr1.right.x的地址:%p \nr1.right.y的地址:%p \n", &r1.left.x, &r1.left.y, &r1.right.x, &r1.right.y)
	fmt.Printf("r1.down.x的地址:%p \nr1.down.y的地址:%p \nr1.up.x的地址:%p \nr1.up.y的地址:%p \n", &r1.down.x, &r1.down.y, &r1.up.x, &r1.up.y)
	//两个*Point类型的本身地址是连续的,但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("\nr2.leftUp本身地址:%p\nr2.rightDown本身地址:%p\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("\nr2.leftUp指向的地址:%p\nr2.rightDown指向的地址:%p\n", r2.leftUp, r2.rightDown)
}

//(2)结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类型)
type A struct {
	Num int
}

type B struct {
	Num int
}

func func6() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)
}

//(3)结构体进行type重新定义(相当于取别名),Golang认为是新的数据类型，但是相互间可以强转
type Student struct {
	Name string
	Age  int
}

type Stu Student

func func7() {
	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1) //强制类型转换
	fmt.Println(stu1, stu2)
}

//(4)struct的每一个字段可以写上一个tag,该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func func8() {
	//1.创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	//2.将monster变量序列化为json格式字符串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr:", string(jsonStr))
}

func main() {
	func8()
}
