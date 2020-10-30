package main

type Hero_Kind int

const (
	None_Hero_Kind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind Hero_Kind
}

//hero指针切片定义为Hero类型
type Heros []*Hero
