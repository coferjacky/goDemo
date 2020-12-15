package main

import (
	"fmt"
	"reflect"
)

//基础数据类型的案例
func relectTest01(b interface{}) {
	//1、先将interface{}转为type类型 kind value
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	rVal := reflect.ValueOf(b)
	//不是真正的100
	fmt.Println(rVal)
	n2 := 10 + rVal.Int()
	//真正的100
	fmt.Println(n2)

	//2 将rVal转回为interface{}
	iV := rVal.Interface()
	fmt.Printf("iv =%v,iv=%T", iV, iV)
	//类型断言 转回基础类型

	num2 := iV.(int)
	fmt.Println(num2)

}

/*type Student struct {
	Name string
	Age int
}


//结构体的案例
func reflectTest02(){

}*/

func main() {
	//定义个int
	var num int = 100

	relectTest01(num)
	//定义个结构体

	/*var stu1 Student=Student{
		Name:"tom",
		Age:20,
	}*/

}
