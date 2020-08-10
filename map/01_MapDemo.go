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
		fmt.Printf("key:%d-----value:%q \n", k)
	}
}
