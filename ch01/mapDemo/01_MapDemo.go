package main

import "fmt"

/*
几种创建map的方法

字典、映射  	key —— value	key： 唯一、无序。 不能是引用类型数据。

			map 不能使用 cap（）

	创建方式：
		1.  var m1 map[int]string		--- 不能存储数据

		2. m2 := map[int]string{}		---能存储数据

		3. m3 := make(map[int]string)		---默认len = 0

		4. m4 := make(map[int]string, 10)   ----创建map的同时指定长度

	初始化：

		1. var m map[int]string = map[int]string{ 1: "aaa", 2:"bbb"}	保证key彼此不重复。

		2. m := map[int]string{ 1: "aaa", 2:"bbb"}

	赋值:

		赋值过程中，如果新map元素的key与原map元素key 相同 	——> 覆盖（替换）

		赋值过程中，如果新map元素的key与原map元素key 不同	——> 添加

	map的使用：

		遍历map：

			for  key值， value值 := range map {

			}

			for  key值 := range map {

			}

		判断map中key是否存在。

			 map[下标] 运算：返回两个值， 第一个表 value 的值，如果value不存在。 nil

						第二个表 key是否存在的bool类型。存在 true， 不存在false

	删除map：

		delete()函数： 	参1： 待删除元素的map	参2： key值

		delete（map， key）	删除一个不存在的key ， 不会报错。

		map 做函数参数和返回值，传引用。





*/
func main() {
	//方法1
	var m1 map[int]string //声明一个map,没有空间不能直接存储key
	// m1[100]="Green"   //执行到这里报错 panic: assignment to entry in nil map
	if m1 == nil {
		fmt.Println("map is nil")
	}

	//方法2
	fmt.Println("------方法2---------")
	m2 := map[int]string{}
	fmt.Println(m2)
	fmt.Println(len(m2))
	m2[100] = "caoyuanyuan"
	fmt.Println(m2)
	fmt.Println(len(m2))

	//方法三
	fmt.Println("------方法3---------")
	m3 := make(map[int]string)
	fmt.Println(m3)
	fmt.Println(len(m3))
	m3[100] = "yuanyuan"
	fmt.Println(m3)
	fmt.Println(len(m3))

	//方法四
	fmt.Println("------方法4---------")
	m4 := make(map[int]string, 5)
	fmt.Println(m4)
	fmt.Println(len(m4))
	m4[100] = "yuanyuan"
	fmt.Println(m4)
	fmt.Println(len(m4))

	//初始化方式1
	fmt.Println("------map---------")
	var m5 map[int]string = map[int]string{1: "caofei", 2: "mycofer"}
	fmt.Println(m5)

	//初始化方式2
	fmt.Println("------map2---------")
	m6 := map[int]string{1: "caofei", 2: "mycofer"}
	fmt.Println(m6)

	//赋值方式1
	fmt.Println("------赋值1---------")
	m7 := make(map[int]string, 1)
	m7[1] = "cofer"
	m7[20] = "jack"
	fmt.Println(m7)

	//赋值方式2
	fmt.Println("------赋值2---------")
}
