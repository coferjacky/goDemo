package main

//使用sort.Slice进行切片元素排序 可以对各种切片类型进行自定义排序
import (
	"fmt"
	"sort"
)

type Hero_Kind1 int

const (
	None = iota
	Tank1
	Assassin1
	Mage1
)

type Hero1 struct {
	Name string
	Kind Hero_Kind1
}

func main() {
	heros1 := []*Hero1{
		{"吕布", Tank1},
		{"李白", Assassin1},
		{"李白1", Tank1},
	}

	//新添加的soft.Slice()及回调函数部分，去掉了接口实现
	sort.Slice(heros1, func(i, j int) bool {
		if heros1[i].Kind != heros1[j].Kind {
			return heros1[i].Kind < heros1[j].Kind
		}
		return heros1[i].Name < heros1[j].Name
	})
	for _, v := range heros1 {
		fmt.Printf("%+v\n", v)

	}
}
