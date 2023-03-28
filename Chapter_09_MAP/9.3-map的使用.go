package main

import (
	"fmt"
)

//方式1
func func1() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["001"] = "宋江"
	a["002"] = "吴用"
	a["003"] = "武松"
	fmt.Println(a)
}

//方式2
func func2() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
}

//方式3
func func3() {
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	heroes["hero4"] = "林冲"
	fmt.Println("heroes:", heroes)
}

//课堂案例:存放3个学生信息，每个学生有name和sex信息
//思路:map[string]map[string]string
func func4() {
	students := make(map[string]map[string]string)
	students["stu01"]=make(map[string]string,3)
	students["stu01"]["name"]="tom"
	students["stu01"]["sex"]="男"
	students["stu01"]["address"]="北京长安街"

	students["stu02"]=make(map[string]string,3)
	students["stu02"]["name"]="mary"
	students["stu02"]["sex"]="女"
	students["stu02"]["address"]="上海黄浦江"

	fmt.Println(students["stu01"])
	fmt.Println(students["stu02"])
}

func main() {
	func4()
}
