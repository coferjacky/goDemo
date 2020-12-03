package main

import "fmt"

type integer int

func (i *integer) change() {
	*i = *i + 1
}

func main() {
	var i integer
	i = 10
	(&i).change()
	fmt.Println(i)
}
