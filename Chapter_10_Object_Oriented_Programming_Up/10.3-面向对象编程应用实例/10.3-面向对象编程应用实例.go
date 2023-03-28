package main

import (
	"fmt"
)

/*
1.步骤
(1)声明(定义结构体)，确定结构体名
(2)编写结构体字段
(3)编写结构体的方法
2.学生案例
(1)编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型
(2)结构体中声明一个say方法，创建Studneg结构体实例(变量)并访问say方法，并将调用结构打印输出。
(3)在main方法中创建Student结构体实例(变量)，并访问say方法后将调用结果打印输出。
*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息: name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]", student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

func test1() {
	//创建一个Student实例变量
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     10000,
		score:  99.98,
	}
	fmt.Println(stu.say())
}

/*
3.盒子案例
(1)编写一个Box结构体,在其中声明三个字段表示一个立方体的长宽高，要求长宽高要从终端获取
(2)声明一个方法获取立方体的体积
(3)创建一个Box结构体变量，打印给定尺寸的立方体的体积
*/
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len*box.width*box.height
}

/*
4.景区门票案例
(1)一个景区根据游客的年龄收取不同价格的门票，比如年龄大于18收费20元，其它情况门票免费
(2)请编写visitor.结构体，根据年龄段能够购买的门票价格并输出
*/
type Visitor struct{
	Name string
	Age int
}

func (visitor *Visitor)showPrice(){
	if visitor.Age>=90||visitor.Age<=8{
		fmt.Println("考虑到安全就不要完了")
		return
	}
	if visitor.Age>18{
		fmt.Printf("游客姓名为%v 年龄为%v 收费20元\n",visitor.Name,visitor.Age)
	}else{
		fmt.Printf("游客姓名为%v 年龄为%v 免费\n",visitor.Name,visitor.Age)
	}
}

func test2(){
	var visitor Visitor
	fmt.Println("请输入你的姓名")
	fmt.Scanln(&visitor.Name)
	if visitor.Name=="n"{
		fmt.Println("退出程序")
	}
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&visitor.Age)
	visitor.showPrice()
}

func main() {
	test2()
}
