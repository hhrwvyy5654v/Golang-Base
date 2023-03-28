package main
import(
	"fmt"
)

func main(){
	//string的基本使用
	var address string ="北京欢迎你!"
	fmt.Println(address) 
	//字符串一旦赋值就不能修改了,GO中字符串是不可变的
	//反引号可以以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	str2:="abc\nabc"
	fmt.Println(str2)
	str3:=`package main
	import(
		"fmt"
	)
	
	func main(){
		var b =false
		fmt.Println("b=",b)
		/*
		注意事项：
		1.bool类型占用存储空间是1个字节
		2.bool类型只能取true或者false
		*/
		
	}`
	fmt.Println(str3)
	//字符串的拼接方式
	var str="hello"+" world"
	str+=" Go!"
	fmt.Println(str)
	//当一行字符串太长时，需要使用到多行字符串时，可以按照如下方式处理:分行时需要将+保留在上一行
	str4:="hello"+" world"+"hello"+" world"+"hello"+" world"+"hello"+" world"+
	"hello"+" world"+"hello"+" world"+"hello"+" world"+"hello"+" world"
	fmt.Println(str4)
}