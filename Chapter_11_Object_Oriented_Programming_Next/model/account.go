package model

import (
	"fmt"
)

// 定义一个结构体
type account struct {
	accountNo string
	balance   float64
	password  string
}

// 工厂模式的函数：相当于构造函数
func NewAccount(accountNo string, balance float64, password string)*account{
	if (len(accountNo)<6||len(accountNo)>10){
		fmt.Println("账号长度错误")
		return nil
	}
	if (len(password)!=6){
		fmt.Println("密码长度错误")
		return nil
	}
	if balance<=20{
		fmt.Println("账户余额错误")
		return nil
	}
	return &account{
		accountNo: accountNo,
		balance: balance,
		password: password,
	}
}


//存款
func (account*account)Deposite(money float64,pwd string){
	//判断密码是否正确
	if pwd!=account.password{
		fmt.Println("密码输入错误")
		return
	}
	//判读存款金额是否正确
	if money<=0.0{
		fmt.Println("金额输入错误")
		return
	}
	account.balance+=money
	fmt.Println("存款成功~~")
}

//取款
func (account*account)WithDraw(money float64,pwd string){
	//判断密码是否正确
	if pwd!=account.password{
		fmt.Println("密码输入错误")
		return
	}
	//判读存款金额是否正确
	if money<=0.0||money>account.balance{
		fmt.Println("金额输入错误")
		return
	}
	account.balance-=money
	fmt.Println("取款成功~~")
}

//查询余额
func (account*account)Query(pwd string){
	//判断密码是否正确
	if pwd!=account.password{
		fmt.Println("密码输入错误")
		return
	}
	fmt.Printf("您的账号:%v 余额:%v",account.accountNo,account.balance)
}