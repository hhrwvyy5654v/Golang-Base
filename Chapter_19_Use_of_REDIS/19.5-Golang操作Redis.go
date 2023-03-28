package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" //引入redis包
)

/*
19.5.1 安装第三方开源Redis库
(1)使用第三方开源redis库:github.com/garyburd/redigo/redis
(2)在使用Redis前先安装第三方Redis库,在GOPATH路径下执行安装指令:
go get github.com/garyburd/redigo/redis
(3)安装成功之后可以看到如下包
*/

/*
19.5.2 Set/Get接口
通过Golang添加和获取key-value
*/
func func1() {
	//通过go向redis写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	//2.通过go向redis写入数据string[key-val]
	_, err = conn.Do("Set", "name", "tomjerry猫猫")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3.通过go向redis读取数据string[key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	/*
		因为返回r是interface{}
		因为name对应的值是string,因此我们需要转换
		nameString:=r.(string)
	*/
	fmt.Println("操作 ok", r)
}

/*
19.5.3 操作Hash
说明:通过Golang对Redis操作数据类型
对hash数据结构，filed-val是一个一个放入和读取
*/
func func2() {
	//通过go向redis写入数据和读取数据
	//1.链接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	//2.通过go向redis写入数据string[key-val]
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset	err=", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	//3.通过go向redis读取数据
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hget	err=", err)
		return
	}

	/*
		因为返回r是interface{}
		因为name对应的值是string,因此我们需要转换
		nameString:=r.(string)
	*/
	fmt.Printf("操作ok r1=%v r2=%v\n", r1, r2)
}

func func3() {
	//通过go向redis写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	//2.通过go向redis写入数据string[key-val]
	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet	err=", err)
		return
	}
	//3.通过go向redis读取数据
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}

/*
19.5.4 批量Set/Get数据
说明:通过Golang对Redis操作,一次操作可以Set/Get多个key-val数据
核心代码：
_,err=c.Do("MSet","name","尚硅谷","address","北京昌平~")
r,err:=redis.Strings(c.Do("MGet","name","address"))

for _,v range r{
	fmt.Println(v)
}
*/

/*
19.5.5给数据设置有效时间
说明：通过Golang对Redis操作,给key-value设置有效时间
核心代码：
// 给name数据设置有效时间为10s
_,err=c.Do("expire","name",10)
*/

/*
19.5.6
说明:通过Golang对Redis操作List数据类型
核心代码:
_,err=c.Do("lpush","heroList","no1:宋江",30,"no2:卢俊义",28)
r,err:=redis.String(c.Do("rpop","heroList"))
*/

/*
19.5.7 Redis链接池
说明：通过Golang对Redis操作，还可以通过Redis链接池，流程如下：
(1)事先初始化一定数量的链接，放入到链接池
(2)当Go需要操作Redis时,直接从Redis链接池取出链接即可
(3)这样可以节省临时获取Redis链接的时间,从而提高效率
*/
//定义一个全局的pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数,0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			//初始化链接的代码,链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func func4() {
	conn := pool.Get() //从链接池中取出一个链接
	defer conn.Close() //关闭链接池,一旦关闭则不能再取出

	_, err := conn.Do("Set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

	//如果我们要从pool取出链接,一定要保证
}

func main() {
	func4()
}
