package main

import (
	"fmt"
	"sort"
)

//sort.String能比07更加简化-对整形切片进行排序

func main() {
	namesss := []string{
		"3.fdadsaf",
		"5.fdafdsafsa",
		"zfdafasdf",
		"2fdafdasfa",
		"1fdasfdsaf",
	}
	sort.Strings(namesss)
	for _, v := range namesss {
		fmt.Printf("%s\n", v)
	}
}
