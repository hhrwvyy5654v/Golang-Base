package main
import(
	"fmt"
)


//Go语言不支持三元运算符
func main(){
	var n int
	var i int =10
	var j int =12
	//传统的三元运算符:n=i>j?i:j
	if i>j{
		n=i
	}else{
		n=j
	}
	fmt.Println("n= ",n)
}