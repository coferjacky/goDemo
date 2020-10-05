package main

import (
	"fmt"
	"sort"
)

func main() {
	//var ages map[string]int  声明map后直接赋值nil会报错 ,必须创建map 下一行这种
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["zlic"] = 32
	ages["xli"] = 38 //同键值会覆盖
	fmt.Println(len(ages))
	for _, v := range ages {
		fmt.Println(v)
	}

	//将map键值取出并排序放进slice
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names) //字典排序
	for _, name := range names {
		fmt.Println(name)
	}

}
