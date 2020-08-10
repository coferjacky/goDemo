package main

//遍历map
import "fmt"

func main() {
	var m5 map[int]string = map[int]string{1: "luffy", 122: "sanyi"}
	for k, v := range m5 {
		fmt.Printf("key:%d-----value:%q \n", k, v)

	}
	//v可以省略
	for k := range m5 {
		fmt.Printf("key:%d \n", k)

	}

	//判断map中的key是否存在
	var m6 map[int]string = map[int]string{1: "fsfs", 78: "faddddd"}
	if v, has := m6[3]; has { //m6[下标]返回两个值，第一个是value，第二个bool 代表key是否存在,这种写法是在表达式前添加一个赋值语句，在判断has是否为true
		fmt.Print("value=", v, " has=", has)
	} else {
		fmt.Println("value=", v, " has=", has)
	}
}
