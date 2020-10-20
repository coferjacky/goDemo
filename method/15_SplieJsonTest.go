package main

import (
	"encoding/json"
	"fmt"
)

//手机例子一枚来演示如何使用匿名结构体分离JSON数据
//手机屏幕
type Screen struct {
	Size         float32 //屏幕尺寸
	Res_x, Res_y int     //屏幕水平和垂直分辨率

}

//手机电池
type Battery struct {
	Capacity int
}

//手工准备一个JSON数据
func gen_Json_Data() []byte {
	//定义一个匿名函数
	raw := &struct {
		Screen
		Battery
		Has_Touch_Id bool
	}{
		Screen: Screen{
			Size:  5.5,
			Res_x: 1920,
			Res_y: 1080,
		},
		Battery: Battery{
			2910,
		},
		Has_Touch_Id: true,
	}
	//使用json.Marchal进行JSON序列化，将raw变量序列化为[]byte格式的JSON数据
	json_Data, _ := json.Marshal(raw)
	return json_Data

}

//开始分离JSON数据
func main() {
	//生产一段JSON数据
	json_Data := gen_Json_Data()
	//将json Data的[]byte类型的JSON数据转换为字符串格式并打印输出。
	fmt.Println(string(json_Data))
	//只需要屏幕和指纹识别信息的结构和实例
	screen_And_Touch := struct {
		Screen
		Has_Touch_ID bool
	}{}
	//反序列化到screen_And_Touch
	json.Unmarshal(json_Data, &screen_And_Touch)

	fmt.Printf("%+v\n", screen_And_Touch)
	//只要电池和指纹信息 的结构和实例
	battery_And_Touch := struct {
		Battery
		Has_Touch_ID bool
	}{}
	//反序列化到screen_And_Touch
	json.Unmarshal(json_Data, &battery_And_Touch)
	fmt.Printf("%+v\n", battery_And_Touch)

}
