package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		//如果vs值长度不为0则输出该切片的第一个元素
		return vs[0]
	}
	return ""
}
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
func main() {
	m := Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item")) //返回第一个元素
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))

	m.Add("item", "3")

}
