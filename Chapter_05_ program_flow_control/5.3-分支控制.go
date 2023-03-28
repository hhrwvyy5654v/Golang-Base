package main
import(
	"fmt"
)

//单分支
func func1(){
	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	if age>18{
		fmt.Println("你的年龄大于18,要对自己的行为负责!")
	}
}

//双分支
func func2(){
	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	if age>18{
		fmt.Println("你的年龄大于18,要对自己的行为负责!")
	}else{
		fmt.Println("你的年龄不大，这次放过你了")
	}
}

//多分支
func func3(){
	var score int
	fmt.Println("请输入你的成绩:")
	fmt.Scanln(&score)

	if score==100{
		fmt.Println("奖励一台BMW")
	}else if score>=60&&score<=80{
		fmt.Println("奖励一台Iphone")
	}else{
		fmt.Println("什么都不奖励")
	}
}

func main(){
	
}