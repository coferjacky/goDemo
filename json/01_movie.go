package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Movie struct {
		Title  string
		Year   int  `json:"released"` //这里面的``将替换掉具体的json成员名
		Color  bool `json:"color"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "casaBlance", Year: 942, Color: false,
			Actors: []string{"aa", "bb", "cc"}},

		{Title: "Cool hand luke", Year: 942, Color: false,
			Actors: []string{"aa", "bb", "cc"}},

		{Title: "Bullitt", Year: 942, Color: false,
			Actors: []string{"aa", "bb", "cc"}},
	}

	data, err := json.Marshal(movies) //这个看上去很难阅读
	/*
		编码结构体到json格式
	*/

	data1, err := json.MarshalIndent(movies, "", "   ") //这个看上去就方便阅读了
	if err != nil {
		log.Fatalf("jSON fail：s%", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%s\n", data1)

	/*
		将json格式反编码到一个struct切片中去，只保留title成员
	*/
	var titles []struct{ Title string } //声明一个struct切片

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json fialed:%s", err)
	}
	fmt.Println(titles)

}
