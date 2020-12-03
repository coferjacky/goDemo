package model

/*
要求：结构体具有字段 账号（6-10位之间），余额（必须大于20），密码（必须6位）
通过setXXX方法给account字符赋值
*/
import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数字不对")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}

}

func (acc *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("必须6位")
		return
	}
	acc.pwd = pwd
}
func (acc *account) GetPwd() string {
	return acc.pwd

}
