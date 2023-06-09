package main

func main() {

}

/*
1.标识符概念：
(1)Golang对各种变量、方法、函数等命名时使用的字符序列称为标识符；
(2)凡是自己可以起名字的地方都叫标识符

2.标识符的命名规则
(1)由26个英文字母大小写，0-9，_组成
(2)数字不可以开头
(3)Golang中严格区分大小写
(4)标识符不能包含空格
(5)下划线"_"本身在Go中是一个特殊的标识符，称为空标识符，可以代表任何其他的标识符，但它对应的值会被忽略(比如：忽略某个返回值)。所以仅能被作为占位符使用，不能作为标识符使用。
(6)不能以系统保留字作为标识符(一个有25个)

3.标识符命名注意事项
(1)包名：保持package的名字和目录一致，尽量采取有意义的包名，简短有意义，不要和标注库冲突
(2)变量名、函数名、常量名：采用驼峰法
(3)如果变量名、函数名、常量名首字母大写，则可以被其他的访问;如果首字母小写，则只能在本包中使用。在Golang中没有public、private等关键字。

*/
