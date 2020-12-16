package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jacky")
	fmt.Println(str)
}
