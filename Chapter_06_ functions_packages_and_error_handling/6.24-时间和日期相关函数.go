package main

import (
	"fmt"
	"time"
)

/*
基本介绍:
	(1)时间和日期相关函数,需要导出time包
	import(
		"time"
	)
*/

//(2)time.Time 类型，用于表时间
func func1() {
	now := time.Now()
	fmt.Printf("now=%v now type=%T", now, now)
}

//(3)获取其它的日期时间
func func2() {
	now := time.Now()
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
}

//(4)使用Printf或者SPrintf格式化日期时间
func func3() {
	now := time.Now()
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)
}

//(5)使用time.Format()方法完成
func func4() {
	now := time.Now()
	fmt.Printf(now.Format("2006-01-02 15:09:34"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:09:34"))
	fmt.Println()
}

/*
(6)时间的常量
const(
	Nanosecond Duration=1	//纳秒
	Microsecond = 1000*Nanosecond	//微妙
	Millisecond	= 1000*Microsecond	//毫秒
	Second = 1000*Millisecond	//秒
	Minute = 60*Second	//分钟
	Hour = 60*Minute	//小时
)
*/

//(7)结合Sleep来使用一下时间常量
func func5() {
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠:time.sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i == 100 {
			break
		}
	}
}

//time的Unix和UnixNano的方法
func func6() {
	now := time.Now()
	//Unix将t表示为Unix时间,即为时间点January 1,1970 UTC到当前时间t所经过的时间(单位秒)
	//UnixNano将t表示为Unix时间,即从时间January 1,1970 UTC到时间点t所经过的时间(单位纳秒)。
	//如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。
	//注意这意味着Time零值调用UnixNano方法的话，结果是未定义的
	fmt.Printf("unix时间戳=%v\nunixNano时间戳=%v\n", now.Unix(), now.UnixNano())
}

func main() {
	func6()
}
