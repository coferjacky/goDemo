package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//连接到redis
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	defer c.Close()
	//向redis写入数据key:string value:int
	_, err = c.Do("Set", "key11", 998)
	if err != nil {
		fmt.Println(err)
		return
	}
	//从redis取出数据key：string value :string,用redis内置的方法int来转换interface类型的数据
	r, err := redis.Int(c.Do("Get", "key11"))
	if err != nil {
		fmt.Println("get key1 failed", err)
		return
	}
	fmt.Println(r)
}
