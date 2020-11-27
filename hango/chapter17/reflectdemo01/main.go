package main

import (
	"fmt"
	"reflect"
)

//演示反射
//演示对（基本数据类型、interface{}\reflect.Value）进行反射的基本操作

func reflectTest01(b interface{}) {
	//通过反射获取传入的变量的type,kind,值
	//先获取reflect.type，TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。
	rType := reflect.TypeOf(b)
	fmt.Println("rtype=", rType)

}

func reflectTest02(b interface{}) {
	//通过反射获取传入的变量的type,kind,值
	//先获取reflect.type，TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。
	rType := reflect.TypeOf(b)
	fmt.Println("rtype=", rType)

	//2获取到refect.value
	rVal := reflect.ValueOf(b)
	//num1:=rVal+2  无法当成一般基础类型来使用
	fmt.Printf("rValtype=%T", rVal)
	//如果与2相加可以通过reflect.Type中的方法进行转换后
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	//将该rVal转换回基本类型int
	//1 rVal转成interface{}
	iv := rVal.Interface()
	//2 将interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=", num2)
}

//演示对（结构体、interface{}\reflect.Value）进行反射的基本操作
type student struct {
	Name string
	Age  int
}

func main() {
	//演示对（基本数据类型、interface{}\reflect.Value）进行反射的基本操作
	var num int = 100
	reflectTest01(num)
	stu := student{
		Name: "tom",
		Age:  16,
	}
	reflectTest02(stu)

}
