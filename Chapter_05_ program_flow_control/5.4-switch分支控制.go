package main

import (
	"fmt"
)

/*
switch的使用的注意事项和细节
1.case/switch后是一个表达式(即：常量值、变量、一个有返回值的函数等都可以)
比如switch test(key)+1{

}
2.case后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
3.case后面可以带多个表达式，使用逗号间隔。比如case表达式1，表达式2……
比如case n2,10,5:
4.case后面的表达式如果是常量(字面量),则要求不能重复
5.case后面不需要带break,程序匹配到一个case后就会执行对应的代码块,然后退出switch,如果一个都匹配不到,则执行default
6.default语句不是必须的
7.switch后也可以不带表达式，类似if--else分支使用。见func2()。
8.switch后也可以直接声明/定义一个变量，分号结束，但不推荐使用。见func3()。
9.switch穿透-fallthrough,如果在case语句块后增加fallthrough,则会继续执行下一个case,也叫switch穿透。见func4()。
10.Type Switch:switch语句还可以被用于type-switch来判断某个interface来判断interface变量中实际指向的变量类型。见func5()。
*/

func func1() {
	var key byte
	fmt.Println("请输入一个字符串 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一,猴子穿新衣")
	case 'b':
		fmt.Println("周二,猴子当小二")
	case 'c':
		fmt.Println("周三,猴子爬雪山")
	default:
		fmt.Println("输入有误")
	}
}

func func2() {
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age==10")
	case age == 20:
		fmt.Println("age==20")
	default:
		fmt.Println("没有匹配到")
	}
}

func func3() {
	switch grade := 90; {
	case grade > 90:
		fmt.Println("成绩优秀……")
	case grade >= 70 && grade <= 90:
		fmt.Println("成绩优良……")
	case grade >= 60 && grade <= 70:
		fmt.Println("成绩及格……")
	default:
		fmt.Println("不及格……")
	}
}

func func4() {
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
		fallthrough //默认只能穿透一层
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}
}

func func5() {
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型~ :%T", i)
	case int:
		fmt.Printf("x是int型")
	case float64:
		fmt.Printf("x是float64型")
	case func(int) float64:
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Printf("未知型号")
	}
}

func main() {
	func4()
}
