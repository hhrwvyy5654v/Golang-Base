package main

import (
	"fmt"
)

/*
1.切片初始化时 var slice=arr[startIndex:endIndex]
说明：从arr数组下标为startIndex,取到下标为endIndex的元素(不含arr[endIndex])
2.切片初始化时，依然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
var slice=arr[0:end]	可以简写为	var slice=arr[:end]
var slice=arr[start:len(arr)]	可以简写为	var slice=arr[start:]
var slice=arr[0:len9(arr)]	可以简写为	var slice=arr[:]
3.cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
4.切片定义完后还不能使用，因为本身是一个空的，需要让其引用到一个数组或者make一个空间切片来使用
5.切片可以继续切片
*/
func func5() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice1 := arr[1:4]
	slice2 := slice1[1:3]
	slice2[0] = 100 //arr、slice1、slice2指向的数据空间是同一个

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr", arr)
}

//6.用append内置函数可以对切片进行动态追加
func func6() {
	var slice []int = []int{100, 200, 300}
	fmt.Println("slice:", slice)
	slice = append(slice, 400, 500, 600)
	fmt.Println("slice:", slice)

	//通过append将切片slice追加给slice
	slice = append(slice, slice...)
	fmt.Println("slice:", slice)
}

/*
切片append操作的底层原理分析：
	切片append操作的本质就是对数组扩容,
	go底层会创建一下新的数组newArr(安装扩容后大小),
	将slice原来包含的元素拷贝到新的数组newArr,
	slice重新引用到newArr



7.切片的拷贝操作
*/
func func7() {
	var slice1 []int = []int{100, 200, 300}
	var slice2 = make([]int, 10)
	copy(slice2, slice1)
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)
}

/*
copy(para1,para2)参数的数据类型是切片
para1和para2的数据空间是独立的，互不影响

8.关于拷贝的注意事项
*/
func func8() {
	var a []int = []int{1, 2, 3, 4, 5}
	var slice = make([]int, 1)
	fmt.Println("slice:", slice)
	copy(slice, a)
	fmt.Println("slice:", slice)
}

//9.切片是引用类型，所以在传递时遵循引用传递机制
func func9() {
	var slice1 []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice1 = arr[:]
	var slice2 = slice1
	slice2[0] = 10
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr:", arr)
}

func func10(slice []int) {
	slice[0] = 100
}

func func11() {
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)
	func10(slice)
	fmt.Println("slice:", slice)
}

func main() {
	func11()
}
