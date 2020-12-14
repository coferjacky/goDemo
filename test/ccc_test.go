package main

import (
	"fmt"
	"testing"
)

//编写测试用例,注意TEST后面的名字是需要测试的函数名，TestXxx 其中Xxx第一个字母必须大写
func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10)执行错误,期望值=%v,实际值=%v")
		t.Fatalf("AddUpper(10)执行错误,期望值=%v,实际值=%v", 55, res) //相当于输出一个日志 并退出
	}
	//如果正确就输出日志
	t.Logf("adddupper执行正确。。。")
}
func TestHello(t *testing.T) {
	fmt.Println("Hello被调用")
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 8)
	if res != 2 {
		t.Fatalf("AddUpper(10)执行错误,期望值=%v,实际值=%v", 2, res)
	}
	t.Logf("getSubr执行正确。。。")
}
