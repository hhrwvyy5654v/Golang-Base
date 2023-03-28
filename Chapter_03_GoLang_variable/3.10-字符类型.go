package main
import(
	"fmt"
)

func main(){
	var c1 byte='a'
	var c2 byte='0'
	//当我们直接输出byte值时，就是直接输出了对应的字符的码值
	fmt.Println("c1=",c1)
	fmt.Println("c2=",c2)
	//如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2)
	var c3 int ='北'	//overflow溢出
	fmt.Printf("c3=%c c3对应码值=%d\n",c3,c3) 
	//可以直接给某一个变量赋予一个数字，然后按照格式化输出%c会输出该数字的unicode字符
	var c4 int =22269
	fmt.Printf("c4=%c \n",c4)
}