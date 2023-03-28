package main
import(
	"fmt"
)

func main(){
	var days int=97
	var week int =days/7
	var day int =days%7
	fmt.Printf("%d个星期零%d天\n",week,day)
}