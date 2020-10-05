package main

import "fmt"

//将slice类型作为key值放进map中
var m = make(map[string]int)

func main() {

	slice1 := []string{"aa", "bb", "dd", "eee"}
	add(slice1)
	count(slice1)
}
func k(list []string) string {
	return fmt.Sprint("%q", list)
}
func add(list []string) {
	m[k(list)]++
}
func count(list []string) int {
	return m[k(list)]
}
