package main

import (
	"fmt"
	"reflect"
)

//定义一个monster的结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//显示具体s值-方法0
func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end-----")
}

//放回两个数的和-方法1
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//设置初始化值-方法3
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex

}
func main() {
	//创建一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   100,
		Score: 30.8,
		//此处将sex字段省略
	}
	//将Monster实例传给TestStruct函数
	TestStruct(a)
}

func TestStruct(a interface{}) {
	//获取reflect.Type的类型
	typ := reflect.TypeOf(a)
	//获取reflect.ValueOf类型
	val := reflect.ValueOf(a)
	fmt.Println(typ)
	fmt.Printf("reflect.ValueOf类型%v\n", val)
	//获取到a对应的类别
	kd := val.Kind()
	fmt.Println(kd)
	//判断一下是否为结构体,不是结构体就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("该结构体共计有%d变量\n", num)

	//遍历结构体的所有变量
	for i := 0; i < num; i++ {
		fmt.Printf("filed %d:值为=%v\n", i, val.Field(i))
		//需要通过reflect.Type来获取tag标签值
		tagVal := typ.Field(i).Tag.Get("json")
		//
		if tagVal != "" {
			fmt.Printf("Field %d:tag为=%v\n", i, tagVal)
		}
	}

	//获取该结构体有多少个方法
	numofMethod := val.NumField()
	fmt.Printf("struct has %d methods\n", numofMethod)
	//获取到第二个方法，参数传入nil,注意：方法排序是按照函数名的排序（ascii码）
	val.Method(1).Call(nil)
	//声明一个value切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	//传入参数是[]reflect.Value,返回[]reflect.Value
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}
