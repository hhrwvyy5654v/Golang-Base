package main
import(
	"fmt"
)

/*Go变量使用的三种方式*/

//第一种：指定变量类型，声明后若不赋值则使用默认值
//int的默认值是0
func func1(){
	var i int
	fmt.Println("i=",i)
}

//第二种：根据值自行判定变量类型(类型推导)
func func2(){
	var num =10.11
	fmt.Println("num=",num)
}

//*第三种：省略var,注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误
func func3(){
	name:="Tom"
	fmt.Println("name=",name)
}

//第四种：多变量声明
func func4(){
	//一次性声明多个变量方式1
	var n1,n2,n3 int
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	//一次性声明多个变量方式2
	var n4,name,n5=100,"Tom",888
	fmt.Println("n4=",n4,"name=",name,"n5=",n5)

	//一次性声明多个变量的方式3，同样可以使用类型推导
	n6,language,n7:=100,"chinese",888
	fmt.Println("n6=",n6,"language=",language,"n7=",n7)
}

func main(){
	
}