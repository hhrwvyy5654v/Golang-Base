package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 给Monster绑定方法Store可以将一个Monster变量(对象),序列化后保存到文件中
func (this *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err=", err)
		return false
	}
	//保存到文件
	filepath := "D:/Go Workspace/src/Basic/Chapter_15_Unit_Testing/monster.ser"
	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println("write file err=", err)
		return false
	}
	return true
}

// 给Monster绑定方法ReStore,可以将一个序列化的Monster从文件中读取,并反序列化为Monster对象,检查反序列化,名字正确
func (this *Monster) ReStore() bool {
	//1.先从文件中，读取序列化中的字符串
	filepath := "D:/Go Workspace/src/Basic/Chapter_15_Unit_Testing/monster.ser"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile err=", err)
		return false
	}

	//2.使用读取到data[]byte,对反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
		return false
	}
	return true
}
