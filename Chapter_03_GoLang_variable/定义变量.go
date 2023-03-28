package main
import "fmt"

func main(){
	// Golang的变量使用方式：
	// (1)指定变量类型，声明后若不赋值，则使用默认值
	// int的默认值是0
	var i int
	i=10
	fmt.Println("i=",i)
	// (2)根据自行判定变量类型(类型推导)
	var num = 11.11
	fmt.Println("num=",num)
	// (3)省略var,注意：左侧的变量不应该是已经声明过的，否则会导致编译错误
	//以下的方式相当于 
	//var name string 
	//name = "Tom"
	name:="Tom"
	fmt.Println("name=",name)
}