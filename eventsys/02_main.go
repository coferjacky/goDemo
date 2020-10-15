package main

import "fmt"

//声明角色的结构体
type Actor struct {
}

func main() {

	a := new(Actor)
	//注册名为ON_SKILL回调，实现代码由a的on_event进行处理
	Register_Event("On_Skill", a.On_Event)
	//注册一个on_skill事件，实现代码由Globalevent进行处理，虽然名字是同一个名字，但是前面注册时间不会被覆盖，而是添加到事件系统中去了，关联Onskill 事件的函数列表
	Register_Event("On_SKill", Global_Event)
	Call_Event("On_Skill", 100)
}

//为角色结构添加一个on_event方法，这个方法有para参数
func (a *Actor) On_Event(param interface{}) {
	fmt.Println("actor event:", param)
}

//全局事件。普通函数实现全局事件的处理
func Global_Event(Param interface{}) {
	fmt.Println("global event:", Param)
}
