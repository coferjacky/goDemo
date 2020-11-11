package main

import "fmt"

//判断接口类型

//电子支付方式
type Alipay struct {
}

//为Alipayt添加了Can_User_Face_ID方法，表示电子支付方式支持刷脸
func (a *Alipay) Can_User_Face_Id() {

}

//现金支付方式
type Cash struct {
}

//为Cash添加了Stolen() 方法，表示现金支付方式会被 偷
func (c *Cash) Stolen() {

}

//刷脸特性的接口
type Cantain_Can_Use_Face_Id interface {
	Can_User_Face_Id()
}

//具备被偷特性的接口
type Cantain_Stolen interface {
	Stolen()
}

func print(pay_Method interface{}) {
	switch pay_Method.(type) {
	case Cantain_Can_Use_Face_Id:
		fmt.Printf("%T can use faceid\n", pay_Method)
	case Cantain_Stolen:
		fmt.Printf("%T may be stolen\n", pay_Method)
	}
}
func main() {
	print(new(Alipay))
	print(new(Cash))
}
