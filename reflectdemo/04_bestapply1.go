package main

import "fmt"

//定义一个monster的结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//显示具体s值
func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end-----")
}

//放回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
