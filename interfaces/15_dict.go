package main

import "fmt"

//使用空接口的实现可以保存任意值的字典
type Dictionary struct {
	data map[interface{}]interface{} //键值都是interface{}类型

}

//根据键获取值
func (d *Dictionary) Get(Key interface{}) interface{} {
	return d.data[Key]
}

//设置键值
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

//遍历所有键值，如果回调返回值为false 停止遍历
//定义回调，类型为func(k,v interface{}) bool  ,
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

//清空所有数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

//创建一个字典
func New_Dictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func main() {
	//创建字典实例
	dict := New_Dictionary()
	//设置三个元素到字典里面
	dict.Set("my factory", 60)
	dict.Set("terra craft", 36)
	dict.Set("dont hungry", 24)
	//获取值并打印
	favorite := dict.Get("dont hungry")
	fmt.Println("donthungryp:", favorite)
	//遍历所有字典的元素
	dict.Visit(func(k, v interface{}) bool {
		//将值转为int类型，并判断是否大于40
		if v.(int) > 40 {
			fmt.Println(k, "is pensive")
			return true
		}
		fmt.Println(k, "is cheap")
		return true
	})
}
