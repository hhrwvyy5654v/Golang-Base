package main
import(
	"fmt"
)

func main(){
	//Golang和java/c不同，Go在不同类型的变量之间赋值时需要显示转换，也就是说Golang数据类型不能自动转化
	var i int32 =100
	//将i从int转换为float32
	var n1 float32=float32(i)
	var n2 int8 =int8(i)
	var n3 int64 =int64(i)	//低精度->高精度
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n",i,n1,n2,n3)
	//在转换中，比如将int64转换成int8,编译时不会报错，只是转换的结果时按照溢出处理。
	var num1 int64=999999
	var num2 int8=int8(num1)
	fmt.Print("num2=",num2,"\n")
}