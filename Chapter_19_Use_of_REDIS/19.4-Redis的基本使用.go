package main

import (
	"fmt"
)


/*
19.4.1 Redis的五大数据类型
Redis的五大数据类型是:String(字符串)、Hash(哈希)、List(列表)、Set(集合)、zset(sorted set:有序集合)

19.4.2 String(字符串)-介绍
string是redis最基本的类型,一个key对应一个value
string类型是二进制安全的。除了普通的字符串外也可以存放图片等数据。
redis中字符串value最大时512M

举例，存放一个地址信息：
address 北京天安门
说明：
key:address
value:北京天安梦

string(字符串)-CRUD
举例说明Redis的String字符串的CRUD操作
set[如果存在就是相当于修改,不存在就是添加]/get/del

19.4.3 String(字符串)-使用细节和注意事项
setex(set with expire)键秒值
mset[同时设置一个或多个key-val对]
mget[同时获取多个key-val]

19.4.4 Hash(哈希,类似golang里的Map)
基本介绍
Redis hash是一个键值对集合。 var user1 map[string]string
Redis hash是一个string类型的filed和value的映射表,hash特别适合用于存储对象。

举例，存放一个User信息:(user1)
user1 name "smith" age 30 job"golang coder"
说明：
key:user1
name 张三

19.4.5 Hash(哈希,类似golang里的Map)-CRUD
举例说明Redis的Hash的CRUD的基本操作
hset/hget/hgetall/hdel
演示添加user信息的案例(name,age)

19.4.6 Hash-使用细节和主义事项
在给user设置name和age时,前面我们是一步一步设置的，使用hmset和hmget可以一次性来设置多个filed的值和返回多个filed的值。
hlen统计一个hash有几个元素
hexists key filed	查看哈希表key中,给定域filed是否存在


19.4.7 课堂练习
举例,存放一个Student信息
stu1 name 张三 age 30 score 80 address 北京


19.4.8 List(列表)-介绍
列表是简单的字符串列表，按照插入顺序排序，可以添加一个元素到列表的头部(左边)或者尾部(右边)
List本质是个链表,List的元素是有序的,元素的值是可以重复的const
举例，存放多个地址信息
city 北京 天津 上海
说明：
key:city
北京 天津 上海就是三个元素

19.4.9 List(列表)-CRUD
举例说明Redis的List的CRUD操作
lpush/rpush/lrange/lpop/rpop/del

19.4.10 List-使用细节和注意事项
(1)index,按照索引下标获得元素(从左到右,编号从0开始)
(2)LLEN key	返回列表key的长度，如果key不存在，则key被解释为一个空列表，返回0
(3)List的其它说明
List数据可以从左或者右插入添加
如果值全部一处，对应的键也就消失了

19.4.11 Set(集合)-介绍
Redis的Set是string类型的无序集合
底层是HashTable数据结构,Set也是存放很多个字符串元素,字符串元素是无序的,而且元素的值不能重复
举例，存放多个邮件列表信息:
email sgg@sohu.com tom@sohu.comconst
说明:
key:email
tn@sohu.com tom@sohu.com就是第二个元素

19.4.12 Set(集合)-CRUD
举例说明Redis的Set的CRUD操作
sadd
smembers[取出所有值]
sismember[判断值是否为成员]
srem[删除指定值]

演示添加多个电子邮件信息的案例

*/

func main() {
	fmt.Println(``)
}
