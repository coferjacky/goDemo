package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err", err)
		return false
	}
	filepath := "c:/aaa.ser"
	error := ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println("writefile err", error)
		return false
	}
	return true
}
func (this *Monster) Restore() bool {
	filepath := "c:/aaa.ser"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read err", err)
		return false
	}
	error := json.Unmarshal(data, this)
	if error != nil {
		fmt.Println("UnMarshel err=", err)
		return false
	}
	return true

}
