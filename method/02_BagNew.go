package main

type Bag struct {
	items []int
}

func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}
func main() {
	b := new(Bag)
	b.Insert(1001)
}
