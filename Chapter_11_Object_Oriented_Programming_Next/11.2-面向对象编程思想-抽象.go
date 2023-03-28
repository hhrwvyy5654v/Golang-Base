package main

import (
	"fmt"
)

//把一类事务的共有的属性(字段)和行为(方法)提取出来形成一个物理模型(结构体),这种研究问题的方法称为抽象
//定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//1.存款
func (account *Account) Deposit(money float64, pwd string) {
	//1.判断输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的密码错误")
		return
	}

	//2.判断存款金额是否正确
	if money <= 0 {
		fmt.Println("输入的金额不正确")
		return
	}
	//3.可以存款
	account.Balance += money
	fmt.Println("存款成功")
}

//2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	//1.判断输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的密码错误")
		return
	}
	//2.判断取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("输入的金额不正确")
		return
	}
	//3.可以取款
	account.Balance -= money
	fmt.Println("取款成功")
}

//3.查询余额
func (account *Account) Query(pwd string) {
	//1.判断输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的密码错误")
		return
	}
	fmt.Printf("您的账号:%v\n余额:%v", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "gs123456",
		Pwd:       "123456",
		Balance:   1000.0,
	}
	account.Query("123456")
	account.Deposit(200, "123456")
	account.WithDraw(159.0, "123456")
}
