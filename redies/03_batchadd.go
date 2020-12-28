package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	dial, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial err=", err)
		return
	}
	defer dial.Close()

	_, err = dial.Do("HMSet", "user02", "name", "zhangsan", "age", 19)
	if err != nil {
		fmt.Println("redis dial err=", err)
		return
	}
	r, err := redis.Strings(dial.Do("HMGet", "user02", "name", "age"))
	fmt.Printf("r=%v\n", r)
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
