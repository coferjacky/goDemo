package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1 定义一个结构体
type Heroboss struct {
	Name string
	Age  int
}

//2 声明一个结构体切片
type HeroSlice []Heroboss

//3、实现系统sort排序的三个方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//定义一个heroboss结构体的切片
	var heros HeroSlice

	for i := 0; i < 10; i++ {
		/*randi:=rand.Intn(100)*/
		hb := Heroboss{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}

		heros = append(heros, hb)
	}

	fmt.Println("-------------排序前---------------")

	for index, v := range heros {
		fmt.Printf("序号=%d,结构体是=%v\n", index, v)
	}

	sort.Sort(heros)

	fmt.Println("-------------排序后---------------")

	for index, v := range heros {
		fmt.Printf("序号=%d,结构体是=%v\n", index, v)
	}

}
