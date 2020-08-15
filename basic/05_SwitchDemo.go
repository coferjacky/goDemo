package main

import (
	"fmt"
)

/*
判断interface变量中实际存储的变量类型
switch x.(type){
	case type:
       statement(s);
    case type:
       statement(s);
	default:
 		statement(s);
}
*/

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x类型:%T", i)
	case int:
		fmt.Printf("x是int型")
	case float64:
		fmt.Printf("x是float64")
	default:
		fmt.Printf("未知数")
	}
}
