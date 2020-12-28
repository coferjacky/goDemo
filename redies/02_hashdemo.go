package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	dial, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dial.Close()
	_, err = dial.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("Hset err=", err)
		return
	}

	r1, err := dial.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("Hget err=", err)
		return
	}

	r2, err := redis.String(dial.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("Hget err=", err)
	}
	fmt.Printf("操作完成，r1=%v,r2=%v\n", r1, r2)
}
