package main

//面向过程实现方法
/*
面向过程的代码描述对象方法概念会越来越麻烦和难以理解
*/
type Bag struct {
	items []int
}

func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

func main() {
	bag := new(Bag)
	Insert(bag, 1001)
}
