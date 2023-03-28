package main

import (
	"fmt"
)

//统计3个班成绩情况，每个班有5名同学，求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]
func func1() {
	var classNum int = 2
	var stuNum int = 5
	var totalSum float64 = 0.0

	for j := 1; j < classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩\n", j, i)
			fmt.Scanln(&score)
			//累计总分
			sum += score
		}

		fmt.Printf("第%d个班级的平均分是%v\n", j, sum/float64(stuNum))
		//将各个班的总成绩累计到totalSum
		totalSum += sum
	}
	fmt.Printf("各个班级的总成绩%v 所有班级平均分是%v\n", totalSum, totalSum/float64(stuNum))
}

//使用for循环打印出金字塔
func func2() {
	var totalLevel int = 20
	//i表示层数
	for i := 1; i <= totalLevel; i++ {
		//在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		//j表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//打印出九九乘法表
func func3() {
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("%v * %v = %v  ", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	func3()
}
