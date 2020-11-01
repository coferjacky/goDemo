package main

import (
	"fmt"
	"sort"
)

//对结构体数据进行排序，先按照英雄的分类进行排序(枚举值从小到大)，然后相同分类按照名字进行排序（名字字符升序排序）

//申明英雄的分类
type Hero_Kind int

const (
	None_Hero_Kind = iota
	Tank
	Assassin
	Mage
)

//定义英雄名单的结构
type Hero struct {
	Name string
	Kind Hero_Kind
}

//hero指针切片定义为Hero类型
type Heros []*Hero

//实现sort.Interface接口取元素数量的方法
func (s Heros) Len() int {
	return len(s)
}

//实现sort.Interface接口比较元素的方法
func (s Heros) Less(i, j int) bool {
	if s[i].Kind != s[j].Kind {
		return s[i].Kind <
			s[j].Kind
	}
	return s[i].Name < s[j].Name
}

//实现sort.Interface接口交互元素方法
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"李白1", Tank},
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}

}
