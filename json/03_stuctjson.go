package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:mon_age`
	Birthday string
	Skill    string
	Sal      float32
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔king",
		Age:      24,
		Birthday: "2009-09-10",
		Skill:    "七伤拳",
		Sal:      9090.0,
	}
	data11, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Printf("monster序列化=%v", string(data11))
}
func main() {
	testStruct()
}
