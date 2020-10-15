package main

//创建一个map实例，这个map通过事件名(string)关联回调列表([]func(interface{})),同一个事件名称可能存在多个事件回调，因此使用回调列表保存
var event_By_Name = make(map[string][]func(interface{}))

//提供给外部的通过事件名注册响应函数入口
func Register_Event(name string, callback func(interface{})) {
	//事件名来查找事件列表
	list := event_By_Name[name]

	//为同一个事件名称在已经注册的实践回调列表中再添加回调函数
	list = append(list, callback)
	//将修改过的函数列表设置到对应事件名总
	event_By_Name[name] = list
}

//调用事件
func Call_Event(name string, param interface{}) {
	//通过名字找到事件列表
	list := event_By_Name[name]
	//遍历这个事件的所有回调
	for _, callback := range list {
		//传入参数调用回调:每个函数回调传入事件参数并调用，就会触发事件实现方的逻辑处理
		callback(param)
	}

}
