package main

import (
	"fmt"
	"reflect"
)

//结构体数据类型的案例
func relectTest02(b interface{}) {
	//1、先将interface{}转为type类型 kind value
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	rVal := reflect.ValueOf(b)

	fmt.Println(rVal)

	//2 将rVal转回为interface{}
	iV := rVal.Interface()
	kind2 := rVal.Kind()
	fmt.Printf("iv =%v,iv type=%T,iv kind=%v\n", iV, iV, kind2)

	//类型断言 转回基础类型
	student := iV.(Student)
	fmt.Println(student.Name)
}

type Student struct {
	Name string
	Age  int
}

func main() {

	//定义个结构体

	var stu1 Student = Student{
		Name: "tom",
		Age:  20,
	}
	relectTest02(stu1)

}
